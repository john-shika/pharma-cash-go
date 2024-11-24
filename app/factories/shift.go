package factories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/factories"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"pharma-cash-go/app/models"
)

func ShiftFactory(DB *gorm.DB) []models.Shift {
	var err error
	nokocore.KeepVoid(err)

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

	return factories.BaseFactory[models.Shift](DB, shifts, "name = ?", func(shift models.Shift) []any {
		return []any{
			shift.Name,
		}
	})
}
