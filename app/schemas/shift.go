package schemas

import (
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
)

type ShiftBody struct {
	Name      string `mapstructure:"name" json:"name" validate:"ascii,min=1"`
	StartDate string `mapstructure:"start_date" json:"startDate" validate:"timeOnly"`
	EndDate   string `mapstructure:"end_date" json:"endDate" validate:"timeOnly"`
}

func ToShiftModel(shift *ShiftBody) *models2.Shift {
	return &models2.Shift{
		Name:      shift.Name,
		StartDate: sqlx.ParseTimeOnly(shift.StartDate),
		EndDate:   sqlx.ParseTimeOnly(shift.EndDate),
	}
}

type ShiftResult struct {
	Name      string            `mapstructure:"name" json:"name"`
	StartDate sqlx.TimeOnlyImpl `mapstructure:"start_date" json:"startDate,omitempty"`
	EndDate   sqlx.TimeOnlyImpl `mapstructure:"end_date" json:"endDate,omitempty"`
}

func ToShiftResult(shift *models2.Shift) ShiftResult {
	if shift != nil {
		return ShiftResult{
			Name:      shift.Name,
			StartDate: shift.StartDate.TimeOnly,
			EndDate:   shift.EndDate.TimeOnly,
		}
	}

	return ShiftResult{}
}
