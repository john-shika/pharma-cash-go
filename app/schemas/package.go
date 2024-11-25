package schemas

import (
	models2 "pharma-cash-go/app/models"
)

type PackageBody struct {
	PackageType string `mapstructure:"package_type" json:"packageType" validate:"ascii,min=1"`
}

func ToPackageModel(body *PackageBody) *models2.Package {
	return &models2.Package{
		PackageType: body.PackageType,
	}
}

type PackageResult struct {
	PackageType string `mapstructure:"package_type" json:"packageType"`
}

func ToPackageResult(pkg *models2.Package) PackageResult {
	return PackageResult{
		PackageType: pkg.PackageType,
	}
}
