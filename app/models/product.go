package models

import (
	"github.com/shopspring/decimal"
	"nokowebapi/apis/models"
	"nokowebapi/sqlx"
)

type Product struct {
	models.BaseModel
	Barcode          string          `db:"barcode" gorm:"unique;index;not null;" mapstructure:"barcode" json:"barcode"`
	Brand            string          `db:"brand" gorm:"index;not null;" mapstructure:"brand" json:"brand"`
	ProductName      string          `db:"product_name" gorm:"index;not null;" mapstructure:"product_name" json:"productName"`
	Supplier         string          `db:"supplier" gorm:"index;not null;" mapstructure:"supplier" json:"supplier"`
	Description      string          `db:"description" gorm:"index;not null;" mapstructure:"description" json:"description"`
	Category         string          `db:"category" gorm:"index;not null;" mapstructure:"category" json:"category"`
	Expires          sqlx.DateOnly   `db:"expires" gorm:"index;not null;" mapstructure:"expires" json:"expires"`
	PurchasePrice    decimal.Decimal `db:"purchase_price" gorm:"index;not null;" mapstructure:"purchase_price" json:"purchasePrice"`
	SupplierDiscount float32         `db:"supplier_discount" gorm:"index;not null;" mapstructure:"supplier_discount" json:"supplierDiscount"`
	VAT              float32         `db:"vat" gorm:"index;not null;" mapstructure:"vat" json:"vat"`
	ProfitMargin     float32         `db:"profit_margin" gorm:"index;not null;" mapstructure:"profit_margin" json:"profitMargin"`
	PackageID        int             `db:"package_id" gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" mapstructure:"package_id" json:"packageId"`
	PackageTotal     float32         `db:"package_total" gorm:"index;not null;" mapstructure:"package_total" json:"packageTotal"`
	UnitID           int             `db:"unit_id" gorm:"index;not null;" mapstructure:"unit_id" json:"unitId"`
	UnitAmount       float32         `db:"unit_amount" gorm:"index;not null;" mapstructure:"unit_amount" json:"unitAmount"`
	UnitExtra        float32         `db:"unit_extra" gorm:"index;not null;" mapstructure:"unit_extra" json:"unitExtra"`

	Package Package `db:"-" gorm:"foreignKey:PackageID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" mapstructure:"package" json:"package"`
	Unit    Unit    `db:"-" gorm:"foreignKey:UnitID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" mapstructure:"unit" json:"unit"`
}

func (p *Product) TableName() string {
	return "products"
}
