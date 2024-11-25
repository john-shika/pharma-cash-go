package models

import "nokowebapi/apis/models"

type Package struct {
	models.BaseModel
	PackageType string    `db:"package_type" gorm:"unique;not null;index;" mapstructure:"package_type" json:"packageType"`
	Products    []Product `db:"-" mapstructure:"products" json:"products"`
}

func (p *Package) TableName() string {
	return "packages"
}
