package models

import "nokowebapi/apis/models"

type Unit struct {
	models.BaseModel
	UnitType string    `db:"unit_type" gorm:"unique;index;not null;" mapstructure:"unit_type" json:"unitType"`
	Products []Product `db:"-" mapstructure:"products" json:"products"`
}

func (Unit) TableName() string {
	return "units"
}
