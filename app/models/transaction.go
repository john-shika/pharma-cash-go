package models

import (
	"github.com/shopspring/decimal"
	"nokowebapi/apis/models"
)

type Transaction struct {
	models.BaseModel
	UserID   uint            `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`
	Total    decimal.Decimal `db:"total" gorm:"index;not null;" mapstructure:"total" json:"total"`
	Pay      decimal.Decimal `db:"pay" gorm:"index;not null;" mapstructure:"pay" json:"pay"`
	Exchange decimal.Decimal `db:"exchange" gorm:"index;not null;" mapstructure:"exchange" json:"exchange"`
	Verified bool            `db:"verified" gorm:"index;not null;" mapstructure:"verified" json:"verified"`

	Carts []Cart      `db:"-" gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"carts" json:"carts"`
	User  models.User `db:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"user" json:"user"`
}
