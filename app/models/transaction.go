package models

import (
	"github.com/shopspring/decimal"
	"nokowebapi/apis/models"
)

type Transaction struct {
	models.BaseModel
	UserID uint            `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`
	Pay    decimal.Decimal `db:"pay" gorm:"index;not null;" mapstructure:"pay" json:"pay"`
	Signed bool            `db:"signed" gorm:"index;not null;" mapstructure:"signed" json:"signed"`
	Closed bool            `db:"closed" gorm:"index;not null;" mapstructure:"closed" json:"closed"`

	Carts []Cart      `db:"-" gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"carts" json:"carts"`
	User  models.User `db:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"user" json:"user"`
}
