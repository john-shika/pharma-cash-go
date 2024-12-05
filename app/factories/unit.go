package factories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/factories"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
)

func UnitFactory(DB *gorm.DB) []any {
	units := []models2.Unit{
		{
			UnitType: "Pcs",
		},
	}

	temp := factories.BaseFactory[models2.Unit](DB, units, "unit_type = ?", func(unit models2.Unit) []any {
		return []any{
			unit.UnitType,
		}
	})

	return nokocore.ToSliceAny(temp)
}
