package schemas

import (
	models2 "pharma-cash-go/app/models"
)

type UnitBody struct {
	UnitType string `mapstructure:"unit_type" json:"unitType" validate:"ascii,min=1"`
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
	UnitType string `mapstructure:"unit_type" json:"unitType"`
}

func ToUnitResult(unit *models2.Unit) UnitResult {
	if unit != nil {
		return UnitResult{
			UnitType: unit.UnitType,
		}
	}

	return UnitResult{}
}
