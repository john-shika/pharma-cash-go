package models

import "nokowebapi/apis/models"

type Package struct {
	models.BaseModel
	PackageType  string  `db:"package_type" gorm:"unique;index;not null;" mapstructure:"package_type" json:"packageType"`

	Products []Product `db:"-" mapstructure:"products" json:"products"`
}

func (Package) TableName() string {
	return "packages"
}
