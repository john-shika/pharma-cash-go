package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/models"
)

type RoleRepositoryImpl interface {
	BaseRepositoryImpl[models.Role]
}

type RoleRepository struct {
	BaseRepositoryImpl[models.Role]
}

func NewRoleRepository(DB *gorm.DB) RoleRepositoryImpl {
	return &RoleRepository{
		NewBaseRepository[models.Role](DB),
	}
}
