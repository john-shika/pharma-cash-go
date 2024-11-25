package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/repositories"
	models2 "pharma-cash-go/app/models"
)

type PackageRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models2.Package]
}

type PackageRepository struct {
	repositories.BaseRepository[models2.Package]
}

func NewPackageRepository(DB *gorm.DB) PackageRepositoryImpl {
	return &PackageRepository{
		BaseRepository: repositories.NewBaseRepository[models2.Package](DB),
	}
}
