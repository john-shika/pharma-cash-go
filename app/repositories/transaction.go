package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/repositories"
	models2 "pharma-cash-go/app/models"
)

type TransactionRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models2.Transaction]
}

type TransactionRepository struct {
	repositories.BaseRepositoryImpl[models2.Transaction]
}

func NewTransactionRepository(DB *gorm.DB) TransactionRepositoryImpl {
	return &TransactionRepository{
		BaseRepositoryImpl: repositories.NewBaseRepository[models2.Transaction](DB),
	}
}
