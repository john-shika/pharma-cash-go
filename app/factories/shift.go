package factories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/factories"
	"nokowebapi/sqlx"
	"pharma-cash-go/app/models"
)

func ShiftFactory(DB *gorm.DB) []any {
	shifts := []models.Shift{
		{
			Name:      "Day Shift",
			StartDate: sqlx.ParseTimeOnly("07:00:00"),
			EndDate:   sqlx.ParseTimeOnly("14:00:00"),
		},
		{
			Name:      "Night Shift",
			StartDate: sqlx.ParseTimeOnly("14:00:00"),
			EndDate:   sqlx.ParseTimeOnly("21:00:00"),
		},
	}

	temp := factories.BaseFactory[models.Shift](DB, shifts, "name = ?", func(shift models.Shift) []any {
		return []any{
			shift.Name,
		}
	})

	size := len(temp)
	result := make([]any, size)
	for i := 0; i < size; i++ {
		result[i] = temp[i]
	}

	return result
}
