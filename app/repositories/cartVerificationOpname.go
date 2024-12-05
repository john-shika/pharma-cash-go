package repositories

import (
	"nokowebapi/apis/repositories"
	models2 "pharma-cash-go/app/models"

	"gorm.io/gorm"
)

type CartVerificationOpnameRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models2.CartVerificationOpname]
}

type CartVerificationOpnameRepository struct {
	repositories.BaseRepositoryImpl[models2.CartVerificationOpname]
}

func NewCartVerificationOpnameRepository(DB *gorm.DB) CartVerificationOpnameRepositoryImpl {
	return &CartVerificationOpnameRepository{
		repositories.NewBaseRepository[models2.CartVerificationOpname](DB),
	}
}
