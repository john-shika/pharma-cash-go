package models

import (
	"nokowebapi/apis/models"

	"github.com/google/uuid"
)

type VerificationOpname struct {
	models.BaseModel
	ProductID          uuid.UUID `db:"product_id" gorm:"index" mapstructure:"product_id" json:"productId"`
	StockOpnameID      uint      `db:"stock_opname_id" gorm:"index" mapstructure:"stock_opname_id" json:"stockOpnameId"`
	SystemPackageTotal int       `db:"system_package_total" gorm:"index;not null;" mapstructure:"system_package_total" json:"systemPackageTotal"`
	SystemUnitExtra    int       `db:"system_unit_extra" gorm:"index;not null;" mapstructure:"system_unit_extra" json:"systemUnitExtra"`
	SystemUnitTotal    int       `db:"system_unit_total" gorm:"index;not null;" mapstructure:"system_unit_total" json:"systemUnitTotal"`
	IsMatch            bool      `db:"is_match" gorm:"index" mapstructure:"is_match" json:"isMatch"`
	NotMatchReason     string    `db:"not_match_reason" gorm:"index;not null;" mapstructure:"not_match_reason" json:"notMatchReason"`
	RealPackageTotal   int       `db:"real_package_total" gorm:"index;not null;" mapstructure:"real_package_total" json:"realPackageTotal"`
	RealUnitExtra      int       `db:"real_unit_extra" gorm:"index;not null;" mapstructure:"real_unit_extra" json:"realUnitExtra"`
	RealUnitTotal      int       `db:"real_unit_total" gorm:"index;not null;" mapstructure:"real_unit_total" json:"realUnitTotal"`
	UserID             uint      `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`

	Product     Product     `db:"-" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"product" json:"product"`
	User        models.User `db:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"user" json:"user"`
	StockOpname StockOpname `db:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"stock_opname" json:"stockOpname"`
}

func (VerificationOpname) TableName() string {
	return "verification_opnames"
}
