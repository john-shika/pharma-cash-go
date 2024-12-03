package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/repositories"
	models2 "pharma-cash-go/app/models"
)

type UnitRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models2.Unit]
}

type UnitRepository struct {
	repositories.BaseRepositoryImpl[models2.Unit]
}

func NewUnitRepository(DB *gorm.DB) UnitRepositoryImpl {
	return &UnitRepository{
		BaseRepositoryImpl: repositories.NewBaseRepository[models2.Unit](DB),
	}
}
