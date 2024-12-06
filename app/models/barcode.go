package models

import "nokowebapi/apis/models"

type Barcode struct {
	models.BaseModel
	Code   string `db:"code" gorm:"not null;" mapstructure:"code" json:"code"`
	Closed bool   `db:"closed" gorm:"not null;" mapstructure:"closed" json:"closed"`
}

func (Barcode) TableName() string {
	return "barcodes"
}
