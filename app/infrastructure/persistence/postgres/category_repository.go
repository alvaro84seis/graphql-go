package postgres

import (
	"strconv"

	"github.com/alvaro84seis/gqlgen-todos/app/domain/entity"
	"github.com/alvaro84seis/gqlgen-todos/app/infrastructure/graph/model"
	"github.com/alvaro84seis/gqlgen-todos/app/infrastructure/persistence/postgres/postgres_connection"
	"github.com/alvaro84seis/gqlgen-todos/app/infrastructure/persistence/postgres/postgres_model"
)

type categoryRepository struct {
	connection postgres_connection.Connection
}

func NewCategoryRepository(connection postgres_connection.Connection) *categoryRepository {
	return &categoryRepository{
		connection: connection,
	}
}

func (s categoryRepository) FindByID(filter model.ICategoria) (*entity.Category, error) {

	db, _ := s.connection.GetConnection()
	dd, _ := db.DB()

	defer dd.Close()
	var query string
	var value string = ""
	if filter.ID == nil && filter.Code != nil {
		query = "code = ?"
		value = *filter.Code
	}
	if filter.ID != nil && filter.Code == nil {
		query = "id = ?"
		value = strconv.Itoa(*filter.ID)
	}

	categoryDB := &postgres_model.CategoryDB{}

	if err := db.Find(categoryDB, query, value).Error; err != nil {
		return nil, err
	}
	var parentCode string
	if categoryDB.ParentCode != nil {
		parentCode = *categoryDB.ParentCode
	}

	categoryResult := &entity.Category{
		ID:           categoryDB.ID,
		Code:         categoryDB.Code,
		DefaultLabel: categoryDB.DefaultLabel,
		Path:         categoryDB.Path,
		Level:        categoryDB.Level,
		ParentCode:   parentCode,
		LabelPath:    categoryDB.LabelPath,
	}

	return categoryResult, nil
}
