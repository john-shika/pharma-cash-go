package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/repositories"
	models2 "pharma-cash-go/app/models"
)

type ShiftRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models2.Shift]
}

type ShiftRepository struct {
	repositories.BaseRepositoryImpl[models2.Shift]
}

func NewShiftRepository(DB *gorm.DB) ShiftRepositoryImpl {
	return &ShiftRepository{
		BaseRepositoryImpl: repositories.NewBaseRepository[models2.Shift](DB),
	}
}
