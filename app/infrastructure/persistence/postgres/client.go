package postgres

import (
	"os"
	"strconv"

	"github.com/alvaro84seis/gqlgen-todos/app/infrastructure/persistence/postgres/postgres_connection"
	"github.com/alvaro84seis/gqlgen-todos/app/infrastructure/persistence/postgres/postgres_model"
)

func AutoMigrateEntities(connection postgres_connection.Connection) {
	// log.Info("AutoMigrateEntities...")
	migrate := postgres_connection.NewMigrate(connection)
	migrate.AutoMigrateAll(
		postgres_model.CategoryDB{},
	)
	// log.Info("AutoMigrateEntities... OK")
}

func CreateDbConnection() *postgres_connection.DbConnection {
	var port int
	var err error

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_database := os.Getenv("DB_DATABASE")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")

	if len(db_host) == 0 || len(db_port) == 0 || len(db_database) == 0 ||
		len(db_username) == 0 || len(db_password) == 0 {
		// log.Fatal("invalid connection data error")
	}

	if port, err = strconv.Atoi(db_port); err != nil {
		// log.Fatal("invalid port error")
	}

	connection := postgres_connection.NewPostgresqlConnection(postgres_connection.Config().
		Host(db_host).
		Port(port).
		DatabaseName(db_database).
		User(db_username).
		Password(db_password),
	)
	return connection
}
