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
		"SELECT id FROM public.users WHERE email = $1 OR phone = $2",
		params...,
	)
	if err != nil {
		helper.LogError(err, "Find user exists error")
		return nil, fmt.Errorf("Find user exists error")
	}

	if existsUser != nil {
		helper.LogError(err, "User has exists")
		return nil, fmt.Errorf("User has exists")
	}

	hashedPassword, err := helper.HashPassword(body.Password)
	if err != nil {
		helper.LogError(err, "Hash password error")
		return nil, fmt.Errorf("Hash password error")
	}

	prepareUser := &entities.UserEntity{
		Name:     body.Name,
		Email:    body.Email,
		Phone:    body.Phone,
		Gender:   body.Gender,
		Password: hashedPassword,
	}

	return s.repository.Create(ctx, prepareUser)
}

var _ IUserService = (*UserService)(nil)
