package postgres_connection

import (
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	// _ "github.com/lib/pq"
)

const LogMode = true

var connection *gorm.DB

type Connection interface {
	GetConnection() (*gorm.DB, error)
	CloseConnection()
}

type DbConnection struct {
	opts *Options
	dsn  string
}

func NewPostgresqlConnection(opts ...*Options) *DbConnection {
	databaseOptions := MergeOptions(opts...)
	dsn := databaseOptions.GetDsnConnection()
	if dsn == "" {
		// log.Fatal(errors.New("Error creating connection, empty dsn").Error())
	}
	return &DbConnection{
		opts: databaseOptions,
		dsn:  dsn,
	}
}

func (r *DbConnection) GetConnection() (*gorm.DB, error) {
	var err error
	if connection == nil || !isAlive() {
		// log.Info("Trying to connect to DB")
		connection, err = gorm.Open(postgres.Open(r.dsn), &gorm.Config{})
		if err != nil {
			// log.WithError(err).Error("error trying to connect to DB")
			return nil, err
		} else {
			// log.Info("Connected to DB")
		}
	}
	// connection.LogMode(LogMode)
	return connection, nil
}

func (r *DbConnection) CloseConnection() {
	sqlDB, err := connection.DB()
	if err != nil {
		// log.WithError(err).Error("No se pudo obtener el objeto sql.DB")
	}
	if err := sqlDB.Close(); err != nil {
		// log.WithError(err).Error("error trying to close connection")
	} else {
		// log.Info("Connection Closed")
	}
}

func isAlive() bool {
	sqlDB, err := connection.DB()
	if err != nil {
		// log.WithError(err).Error("No se pudo obtener el objeto sql.DB")
	}
	if err := sqlDB.Ping(); err != nil {
		// log.WithError(err).Error("error trying to Ping to Db")
		return false
	}
	return true
}
