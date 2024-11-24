package models

import "nokowebapi/apis/models"

type Packaging struct {
	models.BaseModel
	PackagingType string    `db:"packagingType" gorm:"unique;not null;index;" mapstructure:"packagingType" json:"packagingType" yaml:"packagingType"`
}

func (p *Packaging) TableName() string {
	return "packagings"
}
