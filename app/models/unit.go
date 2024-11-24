package models

import "nokowebapi/apis/models"

type Unit struct {
	models.BaseModel
	UnitType string  `db:"unitType" gorm:"unique;not null;index;" mapstructure:"unitType" json:"unitType" yaml:"unitType"`
	Product  Product `db:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"-" json:"-"`
}

func (p *Unit) TableName() string {
	return "unit"
}
