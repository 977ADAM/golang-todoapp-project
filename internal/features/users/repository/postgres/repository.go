package userspostgresrepository

import (
	corepostgrespool "github.com/977ADAM/golang-todoapp-project/internal/core/repository/postgres/pool"
)

type UsersRepository struct {
	pool corepostgrespool.Pool
}

func NewUsersRepository(
	pool corepostgrespool.Pool,
) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}