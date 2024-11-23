package models

import "nokowebapi/apis/models"

type Product struct {
	models.BaseModel
	Name string `db:"name" gorm:"unique;not null;index;" mapstructure:"name" json:"name"`
}

func (p *Product) TableName() string {
	return "products"
}
