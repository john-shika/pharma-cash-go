package schemas

import (
	models2 "pharma-cash-go/app/models"
)

type UnitBody struct {
	UnitType string `mapstructure:"unit_type" json:"unitType" validate:"ascii,min=1"`
}

func ToUnitModel(body *UnitBody) *models2.Unit {
	return &models2.Unit{
		UnitType: body.UnitType,
	}
}

type UnitResult struct {
	UnitType string `mapstructure:"unit_type" json:"unitType"`
}

func ToUnitResult(unit *models2.Unit) UnitResult {
	return UnitResult{
		UnitType: unit.UnitType,
	}
}
