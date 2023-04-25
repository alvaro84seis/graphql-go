package repository_test

import "github.com/alvaro84seis/gqlgen-todos/app/domain/entity"

type userRepository struct {
	// Aquí podríamos tener alguna conexión a una base de datos o alguna API externa.
}

func NewRepository() *userRepository {
	return &userRepository{}
}

func (ur *userRepository) FindByID(id string) (*entity.User, error) {
	// Aquí implementaríamos la lógica para buscar un usuario por su ID.
	return &entity.User{
		ID:    id,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}, nil
}
