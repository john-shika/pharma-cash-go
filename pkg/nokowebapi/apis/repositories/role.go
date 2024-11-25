package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/models"
)

type RoleRepositoryImpl interface {
	BaseRepositoryImpl[models.Role]
}

type RoleRepository struct {
	BaseRepository[models.Role]
}

func NewRoleRepository(DB *gorm.DB) RoleRepositoryImpl {
	return &RoleRepository{
		BaseRepository: NewBaseRepository[models.Role](DB),
	}
}
