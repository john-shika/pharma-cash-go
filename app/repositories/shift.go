package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/repositories"
	"pharma-cash-go/app/models"
)

type ShiftRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models.Shift]
}

type ShiftRepository struct {
	repositories.BaseRepository[models.Shift]
}

func NewShiftRepository(DB *gorm.DB) ShiftRepositoryImpl {
	return &ShiftRepository{
		BaseRepository: repositories.NewBaseRepository[models.Shift](DB),
	}
}
