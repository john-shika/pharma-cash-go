package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/repositories"
	models2 "pharma-cash-go/app/models"
)

type StockRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models2.StockOpname]
}

type StockRepository struct {
	repositories.BaseRepositoryImpl[models2.StockOpname]
}

func NewStockRepository(DB *gorm.DB) StockRepositoryImpl {
	return &StockRepository{
		repositories.NewBaseRepository[models2.StockOpname](DB),
	}
}
