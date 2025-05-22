package services

import (
	"GolangBackend/internal/dto"
	"GolangBackend/internal/entities"
	"GolangBackend/internal/repositories"
	"context"
)

type IUserService interface {
	IBaseService[*entities.UserEntity]
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

func (s *UserService) Create(ctx context.Context, entity *entities.UserEntity) (*entities.UserEntity, error) {
	return s.repository.Create(ctx, entity)
}

func (s *UserService) FindById(ctx context.Context, options *dto.GetByIdOptions) (*entities.UserEntity, error) {
	return s.repository.FindById(ctx, options)
}

func (s *UserService) FindAll(ctx context.Context, options *dto.ListOptions) ([]*entities.UserEntity, error) {
	return s.repository.FindAll(ctx, options)
}

var _ IUserService = (*UserService)(nil)
