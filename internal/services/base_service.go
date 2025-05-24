package services

import (
	"GolangBackend/internal/dto"
	"GolangBackend/internal/entities"
	"GolangBackend/internal/repositories"
	"context"
)

type IBaseService[T entities.IBaseEntity] interface {
	FindById(ctx context.Context, options *dto.GetByIdOptions) (T, error)
	FindAll(ctx context.Context, options *dto.ListOptions) ([]T, error)

	Create(ctx context.Context, entity T) (T, error)
	Update(ctx context.Context, id string, entity T) (T, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type BaseService[T entities.IBaseEntity, R repositories.IBaseRepository[T]] struct {
	repository R
}

func NewBaseService[T entities.IBaseEntity, R repositories.IBaseRepository[T]](repository R) *BaseService[T, R] {
	return &BaseService[T, R]{
		repository: repository,
	}
}

func (s *BaseService[T, R]) FindAll(ctx context.Context, options *dto.ListOptions) ([]T, error) {
	return s.repository.FindAll(ctx, options)
}

func (s *BaseService[T, R]) FindById(ctx context.Context, options *dto.GetByIdOptions) (T, error) {
	return s.repository.FindById(ctx, options)
}

func (s *BaseService[T, R]) Create(ctx context.Context, entity T) (T, error) {
	return s.repository.Create(ctx, entity)
}

func (s *BaseService[T, R]) Update(ctx context.Context, id string, entity T) (T, error) {
	return s.repository.Update(ctx, id, entity)
}

func (s *BaseService[T, R]) Delete(ctx context.Context, id string) (bool, error) {
	return s.repository.Delete(ctx, id)
}
