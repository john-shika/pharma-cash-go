package models

import (
	"github.com/shopspring/decimal"
	"nokowebapi/apis/models"
)

type Transaction struct {
	models.BaseModel
	UserId   uint            `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`
	Money    decimal.Decimal `db:"money" gorm:"not null;" mapstructure:"money" json:"money"`
	Exchange decimal.Decimal `db:"exchange" gorm:"not null;" mapstructure:"exchange" json:"exchange"`
	Valid    bool            `db:"valid" gorm:"index;not null;" mapstructure:"valid" json:"valid"`

	Carts []Cart      `db:"-" gorm:"foreignKey:TransactionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"carts" json:"carts"`
	User  models.User `db:"-" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"user" json:"user"`
}
