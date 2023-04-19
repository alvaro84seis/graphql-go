// application/usecase/get_user_by_id.go
package usecase

import (
	"github.com/alvaro84seis/gqlgen-todos/app/domain/entity"
	"github.com/alvaro84seis/gqlgen-todos/app/domain/repository"
)

type GetUserByIDUseCase interface {
	GetUserByID(id string) (*entity.User, error)
}

type getUserByIDUseCase struct {
	repository repository.Repository
}

func NewGetUserByID(repository repository.Repository) *getUserByIDUseCase {
	return &getUserByIDUseCase{repository: repository}
}

func (g *getUserByIDUseCase) GetUserByID(id string) (*entity.User, error) {
	user, err := g.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
