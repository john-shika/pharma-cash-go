package models

import (
	"nokowebapi/apis/models"
	"nokowebapi/sqlx"
)

type StockOpname struct {
	models.BaseModel
	UserID     uint              `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`
	SubmitedAt sqlx.NullDateOnly `db:"submited_at" gorm:"index;null;" mapstructure:"submited_at" json:"submitedAt"`
	IsVerified bool              `db:"is_verified" gorm:"index;not null;" mapstructure:"is_verified" json:"isVerified"`

	User models.User `db:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"user" json:"user"`
}

func (StockOpname) TableName() string {
	return "stock_opnames"
}
