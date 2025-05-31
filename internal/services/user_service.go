package services

import (
	"GolangBackend/helper"
	"GolangBackend/internal/dto"
	"GolangBackend/internal/entities"
	"GolangBackend/internal/repositories"
	"context"
	"fmt"
)

type IUserService interface {
	IBaseService[*entities.UserEntity]

	Register(ctx context.Context, body *dto.RegisterDTO) (*entities.UserEntity, error)
	Login(ctx context.Context, body *dto.LoginDTO) (*entities.UserResponse, error)
}

type UserService struct {
	*BaseService[*entities.UserEntity, repositories.IUserRepository]
	repository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) *UserService {
	return &UserService{
		BaseService: NewBaseService(repository),
		repository:  repository,
	}
}

func (s *UserService) Register(ctx context.Context, body *dto.RegisterDTO) (*entities.UserEntity, error) {
	var params []any = []any{body.Email, body.Phone}

	existsUser, err := s.repository.ExecuteOne(ctx,
		"SELECT id FROM public.users WHERE email = $1 OR phone = $2 LIMIT 1",
		params...,
	)
	if err != nil {
		helper.LogError(err, "Find user exists error")
		return nil, fmt.Errorf("Find user exists error")
	}

	if existsUser != nil {
		helper.LogError(err)
		return nil, fmt.Errorf("User has exists")
	}

	hashedPassword, err := helper.HashPassword(body.Password)
	if err != nil {
		helper.LogError(err)
		return nil, fmt.Errorf("Hash password error")
	}

	prepareUser := entities.UserEntity{
		Name:     body.Name,
		Email:    body.Email,
		Phone:    body.Phone,
		Gender:   body.Gender,
		Password: hashedPassword,
	}

	return s.repository.Create(ctx, &prepareUser)
}

func (s *UserService) Login(ctx context.Context, body *dto.LoginDTO) (*entities.UserResponse, error) {
	var params []any = []any{body.Account}

	user, err := s.repository.ExecuteOne(ctx,
		"SELECT * FROM public.users WHERE email = $1 OR phone = $1 LIMIT 1",
		params...,
	)

	if err != nil || user == nil {
		return nil, fmt.Errorf("User is not registered")
	}

	isValidPassword := helper.ComparePassword(user.Password, body.Password)

	if !isValidPassword {
		helper.LogError(err, "Account or password is not correct")
		return nil, fmt.Errorf("Account or password is not correct")
	}

	signData := &UserJWT{
		Name: user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Roles: "",
	}

	accessToken, err := SignJWT(signData)
	if err != nil {
		helper.LogError(err)
		return nil, err
	}

	userResponse := &entities.UserResponse{
		UserEntity:   *user,
		AccessToken:  accessToken,
		RefreshToken: accessToken,
	}

	return userResponse, nil
}

var _ IUserService = (*UserService)(nil)
