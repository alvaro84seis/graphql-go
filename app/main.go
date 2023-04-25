package main

import (
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	usecase "github.com/alvaro84seis/gqlgen-todos/app/apllication/usecases"

	"github.com/alvaro84seis/gqlgen-todos/app/infrastructure/graph/generated"
	"github.com/alvaro84seis/gqlgen-todos/app/infrastructure/graph/resolvers"
	"github.com/alvaro84seis/gqlgen-todos/app/infrastructure/persistence/postgres"
	repository_test "github.com/alvaro84seis/gqlgen-todos/app/infrastructure/persistence/postgres_test"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = "3001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// usecases
	connectionPostgres := postgres.CreateDbConnection()
	repositoryPostgres := postgres.NewCategoryRepository(connectionPostgres)

	repositoryTest := repository_test.NewRepository()

	getCategoryByID := usecase.NewGetCategoryByID(repositoryPostgres)
	getUserByID := usecase.NewGetUserByID(repositoryTest)

	resolver := &resolvers.Resolver{
		GetUserByIDUseCase:     getUserByID,
		GetCategoryByIDUseCase: getCategoryByID,
	}

	// Playground
	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))

	// GraphQL
	e.POST("/query", echo.WrapHandler(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))))

	e.Logger.Fatal(e.Start(":8080"))
}
