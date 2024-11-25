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
	repositories.BaseRepository[models2.Unit]
}

func NewUnitRepository(DB *gorm.DB) UnitRepositoryImpl {
	return &UnitRepository{
		BaseRepository: repositories.NewBaseRepository[models2.Unit](DB),
	}
}
