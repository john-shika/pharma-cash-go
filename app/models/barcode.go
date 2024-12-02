package models

import "nokowebapi/apis/models"

type Barcode struct {
	models.BaseModel
	Code   string `db:"code" gorm:"unique;index;not null;" mapstructure:"code" json:"code"`
	Closed bool   `db:"closed" gorm:"index;not null;" mapstructure:"closed" json:"closed"`
}

func (b *Barcode) TableName() string {
	return "barcodes"
}
