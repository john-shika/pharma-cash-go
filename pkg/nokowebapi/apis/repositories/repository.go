package repositories

import "gorm.io/gorm"

type BaseRepositoryImpl interface{}

type BaseRepository struct {
	DB *gorm.DB
}

func NewBaseRepository(DB *gorm.DB) BaseRepositoryImpl {
	return &BaseRepository{
		DB: DB,
	}
}
