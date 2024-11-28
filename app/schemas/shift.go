package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
)

type ShiftBody struct {
	Name      string `mapstructure:"name" json:"name" form:"name" validate:"alphabet"`
	StartDate string `mapstructure:"start_date" json:"startDate" form:"start_date" validation:"timeOnly"`
	EndDate   string `mapstructure:"end_date" json:"endDate" form:"end_date" validation:"timeOnly"`
}

func ToShiftModel(shift *ShiftBody) *models2.Shift {
	if shift != nil {
		return &models2.Shift{
			Name:      shift.Name,
			StartDate: sqlx.ParseTimeOnly(shift.StartDate),
			EndDate:   sqlx.ParseTimeOnly(shift.EndDate),
		}
	}

	return nil
}

type ShiftResult struct {
	UUID      uuid.UUID     `mapstructure:"uuid" json:"uuid"`
	Name      string        `mapstructure:"name" json:"name"`
	StartDate sqlx.TimeOnly `mapstructure:"start_date" json:"startDate,omitempty"`
	EndDate   sqlx.TimeOnly `mapstructure:"end_date" json:"endDate,omitempty"`
}

func ToShiftResult(shift *models2.Shift) ShiftResult {
	if shift != nil {
		return ShiftResult{
			UUID:      shift.UUID,
			Name:      shift.Name,
			StartDate: shift.StartDate.TimeOnly,
			EndDate:   shift.EndDate.TimeOnly,
		}
	}

	return ShiftResult{}
}
