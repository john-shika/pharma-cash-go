package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/repositories"
	models2 "pharma-cash-go/app/models"
)

type ProductRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models2.Product]
}

type ProductRepository struct {
	repositories.BaseRepository[models2.Product]
}

func NewProductRepository(DB *gorm.DB) ProductRepositoryImpl {
	return &ProductRepository{
		BaseRepository: repositories.NewBaseRepository[models2.Product](DB),
	}
}
