package models

import "nokowebapi/apis/models"

type Cart struct {
	models.BaseModel
	UserID        uint `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`
	ProductID     uint `db:"product_id" gorm:"index;not null;" mapstructure:"product_id" json:"productId"`
	TransactionID uint `db:"transaction_id" gorm:"index;null;" mapstructure:"transaction_id" json:"transactionId"`
	PackageTotal  int  `db:"package_total" gorm:"index;not null;" mapstructure:"package_total" json:"packageTotal"`
	UnitExtra     int  `db:"unit_extra" gorm:"index;not null;" mapstructure:"unit_extra" json:"unitExtra"`
	Closed        bool `db:"closed" gorm:"index;not null;" mapstructure:"closed" json:"closed"`

	User        models.User `db:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"user" json:"user"`
	Product     Product     `db:"-" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"product" json:"product"`
	Transaction Transaction `db:"-" gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"transaction" json:"transaction"`
}

func (Cart) TableName() string {
	return "carts"
}
