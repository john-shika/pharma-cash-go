package factories

import (
	"fmt"
	"gorm.io/gorm"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"pharma-cash-go/app/models"
	"pharma-cash-go/app/repositories"
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

	shiftRepository := repositories.NewShiftRepository(DB)

	var check *models.Shift
	for i, shift := range shifts {
		nokocore.KeepVoid(i)

		if check, err = shiftRepository.First("name = ?", shift.Name); err != nil {
			console.Warn(err.Error())
			continue
		}

		if check != nil {
			console.Warn(fmt.Sprintf("shift '%s' already exists", shift.Name))
			continue
		}

		if err = shiftRepository.Create(&shift); err != nil {
			console.Warn(err.Error())
			continue
		}

		console.Warn(fmt.Sprintf("shift '%s' has been created", shift.Name))
	}

	return shifts
}
