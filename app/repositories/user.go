package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/apis/repositories"
)

type UserRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models.User]
}

type UserRepository struct {
	repositories.BaseRepository[models.User]
}

func NewUserRepository(DB *gorm.DB) UserRepositoryImpl {
	return &UserRepository{
		BaseRepository: repositories.NewBaseRepository[models.User](DB),
	}
}
