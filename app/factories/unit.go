package factories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/factories"
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

	size := len(temp)
	result := make([]any, size)
	for i := 0; i < size; i++ {
		result[i] = temp[i]
	}

	return result
}
