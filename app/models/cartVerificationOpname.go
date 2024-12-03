package models

import (
	"nokowebapi/apis/models"
	"nokowebapi/sqlx"
)

type CartVerificationOpname struct {
	models.BaseModel
	UserID         uint              `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`
	ProductID      uint              `db:"product_id" gorm:"index" mapstructure:"product_id" json:"productId"`
	AmountPackage  string            `db:"amount_package" gorm:"index;not null;" mapstructure:"amount_package" json:"amountPackage"`
	AmountUnit     string            `db:"amount_unit" gorm:"index;not null;" mapstructure:"amount_unit" json:"amountUnit"`
	UnitTotal      string            `db:"unit_total" gorm:"index;not null;" mapstructure:"unit_total" json:"unitTotal"`
	NotMatchReason string            `db:"not_match_reason" gorm:"index;not null;" mapstructure:"not_match_reason" json:"notMatchReason"`
	SubmitedAt     sqlx.NullDateOnly `db:"submited_at" gorm:"index;null;" mapstructure:"submited_at" json:"submitedAt"`

	User    models.User `db:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"user" json:"user"`
	Product Product     `db:"-" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"product" json:"product"`
}

func (CartVerificationOpname) TableName() string {
	return "cart_verification_opnames"
}
