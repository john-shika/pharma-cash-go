package schemas

import (
	models2 "pharma-cash-go/app/models"
)

type PackageBody struct {
	PackageType string `mapstructure:"package_type" json:"packageType" validate:"ascii,min=1"`
}

func ToPackageModel(packageBody *PackageBody) *models2.Package {
	if packageBody != nil {
		return &models2.Package{
			PackageType: packageBody.PackageType,
		}
	}

	return nil
}

type PackageResult struct {
	PackageType string `mapstructure:"package_type" json:"packageType"`
}

func ToPackageResult(packageModel *models2.Package) PackageResult {
	if packageModel != nil {
		return PackageResult{
			PackageType: packageModel.PackageType,
		}
	}

	return PackageResult{}
}
