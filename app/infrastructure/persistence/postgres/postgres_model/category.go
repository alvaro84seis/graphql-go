package postgres_model

type CategoryDB struct {
	ID           int     `gorm:"primaryKey;autoIncrement"`
	Code         string  `gorm:"not null"`
	DefaultLabel string  `gorm:"not null;"`
	Path         string  `gorm:"not null"`
	Level        int     `gorm:"not null"`
	ParentCode   *string `gorm:"foreignKey:id"`
	LabelPath    string
}

func (CategoryDB) TableName() string {
	return "structure_cl_current.category"
}
