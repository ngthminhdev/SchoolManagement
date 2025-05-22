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
	return s.FindAll(ctx, options)
}

func (s *BaseService[T, R]) FindById(ctx context.Context, options *dto.GetByIdOptions) (T, error) {
	return s.FindById(ctx, options)
}

func (s *BaseService[T, R]) Create(ctx context.Context, entity T) (T, error) {
	return s.Create(ctx, entity)
}
