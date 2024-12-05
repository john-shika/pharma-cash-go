package models

import (
	"nokowebapi/apis/models"
)

type CartVerificationOpname struct {
	models.BaseModel
	ProductID        uint   `db:"product_id" gorm:"index" mapstructure:"product_id" json:"productId"`
	IsMatch          bool   `db:"is_match" gorm:"index" mapstructure:"is_match" json:"isMatch"`
	NotMatchReason   string `db:"not_match_reason" gorm:"index;not null;" mapstructure:"not_match_reason" json:"notMatchReason"`
	RealPackageTotal int    `db:"real_package_total" gorm:"index;not null;" mapstructure:"real_package_total" json:"realPackageTotal"`
	RealUnitExtra    int    `db:"real_unit_extra" gorm:"index;not null;" mapstructure:"real_unit_extra" json:"realUnitExtra"`
	UserID           uint   `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`

	User    models.User `db:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"user" json:"user"`
	Product Product     `db:"-" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"product" json:"product"`
}

func (CartVerificationOpname) TableName() string {
	return "cart_verification_opnames"
}
