package repository

import "github.com/alvaro84seis/gqlgen-todos/app/domain/entity"

type Repository interface {
	FindByID(id string) (*entity.User, error)
}
