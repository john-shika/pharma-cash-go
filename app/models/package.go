package models

import "nokowebapi/apis/models"

type Package struct {
	models.BaseModel
	PackageType  string  `db:"package_type" gorm:"unique;index;not null;" mapstructure:"package_type" json:"packageType"`
	PackageTotal float32 `db:"package_total" gorm:"index;not null;" mapstructure:"package_total" json:"packageTotal"`

	Products []Product `db:"-" mapstructure:"products" json:"products"`
}

func (p *Package) TableName() string {
	return "packages"
}
