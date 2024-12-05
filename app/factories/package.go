package factories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/factories"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
)

func PackageFactory(DB *gorm.DB) []any {
	packages := []models2.Package{
		{
			PackageType: "BOX",
		},
	}

	temp := factories.BaseFactory[models2.Package](DB, packages, "package_type = ?", func(packageModel models2.Package) []any {
		return []any{
			packageModel.PackageType,
		}
	})

	return nokocore.ToSliceAny(temp)
}
