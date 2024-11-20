package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/apis/repositories"
)

type UserRepository struct {
	repositories.BaseRepository[models.User]
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{
		BaseRepository: repositories.NewBaseRepository[models.User](DB),
	}
}

func (u *UserRepository) Find(wheres ...any) (*models.User, error) {
	return u.BaseRepository.Find(wheres...)
}

func (u *UserRepository) Create(model *models.User) error {
	return u.BaseRepository.Create(model)
}
