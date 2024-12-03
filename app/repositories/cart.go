package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/repositories"
	models2 "pharma-cash-go/app/models"
)

type CartRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models2.Cart]
}

type CartRepository struct {
	repositories.BaseRepositoryImpl[models2.Cart]
}

func NewCartRepository(DB *gorm.DB) CartRepositoryImpl {
	return &CartRepository{
		repositories.NewBaseRepository[models2.Cart](DB),
	}
}
