package models

import "nokowebapi/apis/models"

type Unit struct {
	models.BaseModel
	UnitType string  `db:"unitType" gorm:"unique;not null;index;" mapstructure:"unitType" json:"unitType" yaml:"unitType"`
}

func (p *Unit) TableName() string {
	return "units"
}
