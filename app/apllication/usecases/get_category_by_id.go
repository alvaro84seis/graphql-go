// application/usecase/get_user_by_id.go
package usecase

import (
	"github.com/alvaro84seis/gqlgen-todos/app/domain/entity"
	"github.com/alvaro84seis/gqlgen-todos/app/domain/repository"
	"github.com/alvaro84seis/gqlgen-todos/app/infrastructure/graph/model"
)

type GetCategoryByIDUseCase interface {
	GetCategoryByID(filter model.ICategoria) (*entity.Category, error)
}

type getCategoryByIDUseCase struct {
	repository repository.CategoryRepository
}

func NewGetCategoryByID(repository repository.CategoryRepository) *getCategoryByIDUseCase {
	return &getCategoryByIDUseCase{repository: repository}
}

func (g *getCategoryByIDUseCase) GetCategoryByID(filter model.ICategoria) (*entity.Category, error) {
	user, err := g.repository.FindByID(filter)
	if err != nil {
		return nil, err
	}
	return user, nil
}
