package factories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/factories"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"pharma-cash-go/app/models"
)

func ShiftFactory(DB *gorm.DB) []any {
	shifts := []any{
		models.Shift{
			Name:      "Day Shift",
			StartDate: sqlx.ParseTimeOnly("07:00:00"),
			EndDate:   sqlx.ParseTimeOnly("14:00:00"),
		},
		models.Shift{
			Name:      "Night Shift",
			StartDate: sqlx.ParseTimeOnly("14:00:00"),
			EndDate:   sqlx.ParseTimeOnly("21:00:00"),
		},
	}

	return factories.BaseFactory(DB, shifts, "name = ?", func(shift any) []any {
		return []any{
			nokocore.GetValueWithSuperKey(shift, "name"),
		}
	})
}
