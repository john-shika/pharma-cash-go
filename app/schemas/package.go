package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
)

type PackageBody struct {
	PackageType string `mapstructure:"package_type" json:"packageType" form:"package_type" validate:"ascii"`
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
	UUID        uuid.UUID `mapstructure:"uuid" json:"uuid"`
	PackageType string    `mapstructure:"package_type" json:"packageType"`
	CreatedAt   string    `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt   string    `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt   string    `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}

func ToPackageResult(packageModel *models2.Package) PackageResult {
	if packageModel != nil {
		createdAt := nokocore.ToTimeUtcStringISO8601(packageModel.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(packageModel.UpdatedAt)
		var deletedAt string
		if packageModel.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(packageModel.DeletedAt.Time)
		}
		return PackageResult{
			UUID:        packageModel.UUID,
			PackageType: packageModel.PackageType,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			DeletedAt:   deletedAt,
		}
	}

	return PackageResult{}
}
