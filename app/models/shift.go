package models

import (
	"nokowebapi/apis/models"
	"nokowebapi/sqlx"
)

type Shift struct {
	models.BaseModel
	Name      string            `db:"name" gorm:"unique;index;not null;" mapstructure:"name" json:"name"`
	StartDate sqlx.NullTimeOnly `db:"start_date" gorm:"type:TIME;index;null;" mapstructure:"start_date" json:"startDate"`
	EndDate   sqlx.NullTimeOnly `db:"end_date" gorm:"type:TIME;index;null;" mapstructure:"end_date" json:"endDate"`
	Employees []Employee        `db:"-" mapstructure:"employees" json:"employees"`
}

func (s *Shift) TableName() string {
	return "shifts"
}
