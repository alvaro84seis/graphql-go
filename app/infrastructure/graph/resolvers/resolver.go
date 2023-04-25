package resolvers

import usecase "github.com/alvaro84seis/gqlgen-todos/app/apllication/usecases"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GetUserByIDUseCase     usecase.GetUserByIDUseCase
	GetCategoryByIDUseCase usecase.GetCategoryByIDUseCase
}
