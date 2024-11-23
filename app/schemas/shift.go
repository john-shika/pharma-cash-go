package schemas

import (
	"nokowebapi/nokocore"
	"pharma-cash-go/app/models"
	"time"
)

type ShiftResult struct {
	Name      string    `mapstructure:"name" json:"name"`
	StartDate time.Time `mapstructure:"start_date" json:"startDate"`
	EndDate   time.Time `mapstructure:"end_date" json:"endDate"`
}

func ToShiftResult(shift *models.Shift) ShiftResult {
	if shift != nil {
		timeNow := nokocore.GetTimeUtcNow()
		dateOnly := timeNow.Format("2006-01-02")
		timeSeed := nokocore.Unwrap(time.Parse("2006-01-02", dateOnly))
		return ShiftResult{
			Name:      shift.Name,
			StartDate: timeSeed.Add(shift.StartDate.ToTimeDuration()),
			EndDate:   timeSeed.Add(shift.EndDate.ToTimeDuration()),
		}
	}

	return ShiftResult{}
}
