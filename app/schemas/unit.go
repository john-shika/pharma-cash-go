package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
)

type UnitBody struct {
	UnitType string `mapstructure:"unit_type" json:"unitType" form:"unit_type" validate:"ascii"`
}

func ToUnitModel(unit *UnitBody) *models2.Unit {
	if unit != nil {
		return &models2.Unit{
			UnitType: unit.UnitType,
		}
	}

	return nil
}

type UnitResult struct {
	UUID      uuid.UUID `mapstructure:"uuid" json:"uuid"`
	UnitType  string    `mapstructure:"unit_type" json:"unitType"`
	CreatedAt string    `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt string    `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt string    `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}

func ToUnitResult(unit *models2.Unit) UnitResult {
	if unit != nil {
		createdAt := nokocore.ToTimeUtcStringISO8601(unit.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(unit.UpdatedAt)
		var deletedAt string
		if unit.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(unit.DeletedAt)
		}
		return UnitResult{
			UUID:      unit.UUID,
			UnitType:  unit.UnitType,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		}
	}

	return UnitResult{}
}
