package models

import (
	"nokowebapi/apis/models"
	"nokowebapi/sqlx"
)

type Employee struct {
	models.BaseModel
	UserID    uint              `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`
	ShiftID   uint              `db:"shift_id" gorm:"index;not null;" mapstructure:"shift_id" json:"shiftId"`
	ShiftDate sqlx.NullDateOnly `db:"shift_date" gorm:"index;null;" mapstructure:"shift_date" json:"shiftDate"`
	Shift     Shift             `db:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"shifts" json:"shifts"`
	User      models.User       `db:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"user" json:"user"`
}

func (Employee) TableName() string {
	return "employees"
}
