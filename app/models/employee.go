package models

import (
	"nokowebapi/apis/models"
	"time"
)

type Employee struct {
	models.BaseModel
	UserID    int         `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"user_id"`
	ShiftID   int         `db:"shift_id" gorm:"index;not null;" mapstructure:"shift_id" json:"shiftId"`
	ShiftDate time.Time   `db:"shift_date" gorm:"index;not null;" mapstructure:"shift_date" json:"shiftDate"`
	User      models.User `db:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (s *Employee) TableName() string {
	return "employees"
}
