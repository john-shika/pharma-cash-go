package models

import (
	"gorm.io/gorm"
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

func (Shift) TableName() string {
	return "shifts"
}

func (s *Shift) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (s *Shift) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}
