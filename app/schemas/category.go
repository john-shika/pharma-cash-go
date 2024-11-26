package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
)

type CategoryBody struct {
	CategoryName string `mapstructure:"category_name" json:"categoryName" form:"category_name" validate:"ascii"`
}

func ToCategoryModel(category *CategoryBody) *models2.Category {
	if category != nil {
		return &models2.Category{
			CategoryName: category.CategoryName,
		}
	}

	return nil
}

type CategoryResult struct {
	UUID         uuid.UUID `mapstructure:"uuid" json:"uuid"`
	CategoryName string    `mapstructure:"category_name" json:"categoryName"`
	CreatedAt    string    `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt    string    `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt    string    `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}

func ToCategoryResult(category *models2.Category) CategoryResult {
	if category != nil {
		createdAt := nokocore.ToTimeUtcStringISO8601(category.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(category.UpdatedAt)
		var deletedAt string
		if category.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(category.DeletedAt.Time)
		}
		return CategoryResult{
			UUID:         category.UUID,
			CategoryName: category.CategoryName,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
			DeletedAt:    deletedAt,
		}
	}

	return CategoryResult{}
}
