package repositories

import "GolangBackend/internal/entities"

type IUserRepository interface {
	IBaseRepository[*entities.UserEntity]
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
