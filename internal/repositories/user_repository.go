package repositories

import (
	"GolangBackend/internal/dto"
	"GolangBackend/internal/entities"
	"context"
)

type IUserRepository interface {
	IBaseRepository[*entities.UserEntity]

	Register(ctx context.Context, body *dto.RegisterDTO) (*entities.UserEntity, error)
}

type UserRepository struct {
	*BaseRepository[*entities.UserEntity]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository(
			"public.users",
			func() *entities.UserEntity { return &entities.UserEntity{} },
		),
	}
}

func (r *UserRepository) Register(ctx context.Context, body *dto.RegisterDTO) (*entities.UserEntity, error) {
	return nil, nil
}
