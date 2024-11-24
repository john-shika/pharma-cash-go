package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/repositories"
	"pharma-cash-go/app/models"
)

type EmployeeRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models.Employee]
}

type EmployeeRepository struct {
	repositories.BaseRepository[models.Employee]
}

func NewEmployeeRepository(DB *gorm.DB) EmployeeRepositoryImpl {
	return &EmployeeRepository{
		BaseRepository: repositories.NewBaseRepository[models.Employee](DB),
	}
}
