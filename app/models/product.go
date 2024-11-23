package models

import "nokowebapi/apis/models"

type Product struct {
	models.BaseModel
	Barcode          string  `db:"barcode" gorm:"unique;not null;index;" mapstructure:"barcode" json:"barcode" yaml:"barcode"`
	Merk             string  `db:"merk" gorm:"unique;not null;index;" mapstructure:"merk" json:"merk" yaml:"merk"`
	ProductName      string  `db:"productName" gorm:"unique;not null;index;" mapstructure:"productName" json:"productName" yaml:"productName"`
	Supplier         string  `db:"supplier" gorm:"unique;not null;index;" mapstructure:"supplier" json:"supplier" yaml:"supplier"`
	Description      string  `db:"description" gorm:"unique;not null;index;" mapstructure:"description" json:"description" yaml:"description"`
	Category         string  `db:"category" gorm:"unique;not null;index;" mapstructure:"category" json:"category" yaml:"category"`
	Expired          string  `db:"expired" gorm:"unique;not null;index;" mapstructure:"expired" json:"expired" yaml:"expired"`
	PurchasePrice    float32 `db:"purchasePrice" gorm:"unique;not null;index;" mapstructure:"purchasePrice" json:"purchasePrice" yaml:"purchasePrice"`
	SupplierDiscount float32 `db:"supplierDiscount" gorm:"unique;not null;index;" mapstructure:"supplierDiscount" json:"supplierDiscount" yaml:"supplierDiscount"`
	Ppn              float32 `db:"ppn" gorm:"unique;not null;index;" mapstructure:"ppn" json:"ppn" yaml:"ppn"`
	ProfitMargin     float32 `db:"profitMargin" gorm:"unique;not null;index;" mapstructure:"profitMargin" json:"profitMargin" yaml:"profitMargin"`
	PackagingType    string  `db:"packagingType" gorm:"unique;not null;index;" mapstructure:"packagingType" json:"packagingType" yaml:"packagingType"`
	TotalPackaging   float32 `db:"totalPackaging" gorm:"unique;not null;index;" mapstructure:"totalPackaging" json:"totalPackaging" yaml:"totalPackaging"`
	UnitType         string  `db:"unitType" gorm:"unique;not null;index;" mapstructure:"unitType" json:"unitType" yaml:"unitType"`
	UnitAmount       float32 `db:"unitAmount" gorm:"unique;not null;index;" mapstructure:"unitAmount" json:"unitAmount" yaml:"unitAmount"`
	UnitResidu       float32 `db:"unitResidu" gorm:"unique;not null;index;" mapstructure:"unitResidu" json:"unitResidu" yaml:"unitResidu"`
}

func (Product) TableName() string {
	return "products"
}
