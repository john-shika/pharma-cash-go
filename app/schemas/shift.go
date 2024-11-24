package schemas

import (
	"nokowebapi/sqlx"
	"pharma-cash-go/app/models"
)

type ShiftResult struct {
	Name      string            `mapstructure:"name" json:"name"`
	StartDate sqlx.TimeOnlyImpl `mapstructure:"start_date" json:"startDate,omitempty"`
	EndDate   sqlx.TimeOnlyImpl `mapstructure:"end_date" json:"endDate,omitempty"`
}

func ToShiftResult(shift *models.Shift) ShiftResult {
	if shift != nil {
		return ShiftResult{
			Name:      shift.Name,
			StartDate: shift.StartDate.TimeOnly,
			EndDate:   shift.EndDate.TimeOnly,
		}
	}

	return ShiftResult{}
}
