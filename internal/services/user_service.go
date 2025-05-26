package services

import (
	"GolangBackend/internal/dto"
	"GolangBackend/internal/entities"
	"GolangBackend/internal/repositories"
	"context"
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
	return s.repository.Register(ctx, body)
}

var _ IUserService = (*UserService)(nil)
