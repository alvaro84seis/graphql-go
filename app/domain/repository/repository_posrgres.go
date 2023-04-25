package repository

import (
	"github.com/alvaro84seis/gqlgen-todos/app/domain/entity"
	"github.com/alvaro84seis/gqlgen-todos/app/infrastructure/graph/model"
)

type CategoryRepository interface {
	FindByID(filter model.ICategoria) (*entity.Category, error)
}
