package schemas

import (
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"

	"github.com/google/uuid"
)

type StockOpnameBody struct {
	UnitType string `mapstructure:"unit_type" json:"unitType" form:"unit_type" validate:"ascii"`
}

func ToStockOpnameModel(unit *StockOpnameBody) *models2.Unit {
	if unit != nil {
		return &models2.Unit{
			UnitType: unit.UnitType,
		}
	}

	return nil
}

type StockOpnameResult struct {
	UUID       uuid.UUID         `mapstructure:"uuid" json:"uuid"`
	SubmitedAt sqlx.NullDateOnly `mapstructure:"submited_at" json:"submitedAt"`
	IsVerified bool              `mapstructure:"is_verified" json:"isVerified"`
	CreatedBy  uuid.UUID         `mapstructure:"created_by" json:"createdBy"`
	CreatedAt  string            `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt  string            `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt  string            `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}

func ToStockOpnameResult(stockOpname *models2.StockOpname) StockOpnameResult {
	if stockOpname != nil {
		createdAt := nokocore.ToTimeUtcStringISO8601(stockOpname.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(stockOpname.UpdatedAt)
		var deletedAt string
		if stockOpname.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(stockOpname.DeletedAt)
		}
		return StockOpnameResult{
			UUID:       stockOpname.UUID,
			SubmitedAt: stockOpname.SubmitedAt,
			IsVerified: stockOpname.IsVerified,
			CreatedBy:  stockOpname.User.UUID,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
			DeletedAt:  deletedAt,
		}
	}

	return StockOpnameResult{}
}
