package factories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/factories"
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

	size := len(temp)
	result := make([]any, size)
	for i := 0; i < size; i++ {
		result[i] = temp[i]
	}

	return result
}
