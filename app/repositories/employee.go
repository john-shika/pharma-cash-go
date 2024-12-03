package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/repositories"
	models2 "pharma-cash-go/app/models"
)

type EmployeeRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models2.Employee]
}

type EmployeeRepository struct {
	repositories.BaseRepositoryImpl[models2.Employee]
}

func NewEmployeeRepository(DB *gorm.DB) EmployeeRepositoryImpl {
	return &EmployeeRepository{
		BaseRepositoryImpl: repositories.NewBaseRepository[models2.Employee](DB),
	}
}
