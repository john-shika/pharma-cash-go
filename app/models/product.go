package models

import (
	"nokowebapi/apis/models"
	"time"
)

type Product struct {
	models.BaseModel
	Barcode          string    `db:"barcode" gorm:"unique;not null;index;" mapstructure:"barcode" json:"barcode" yaml:"barcode" validate:"required,alphanum"`
	Merk             string    `db:"merk" gorm:"not null;index;" mapstructure:"merk" json:"merk" yaml:"merk" validate:"required"`
	ProductName      string    `db:"productName" gorm:"not null;index;" mapstructure:"productName" json:"productName" yaml:"productName" validate:"required"`
	Supplier         string    `db:"supplier" gorm:"not null;index;" mapstructure:"supplier" json:"supplier" yaml:"supplier" validate:"required"`
	Description      string    `db:"description" gorm:"not null;index;" mapstructure:"description" json:"description" yaml:"description" validate:"required"`
	Category         string    `db:"category" gorm:"not null;index;" mapstructure:"category" json:"category" yaml:"category" validate:"required"`
	Expired          time.Time `db:"expired" gorm:"not null;index;" mapstructure:"expired" json:"expired" yaml:"expired" validate:"required"`
	PurchasePrice    float32   `db:"purchasePrice" gorm:"not null;index;" mapstructure:"purchasePrice" json:"purchasePrice" yaml:"purchasePrice" validate:"required,gt=0"`
	SupplierDiscount float32   `db:"supplierDiscount" gorm:"not null;index;" mapstructure:"supplierDiscount" json:"supplierDiscount" yaml:"supplierDiscount" validate:"required,gt=0"`
	Ppn              float32   `db:"ppn" gorm:"not null;index;" mapstructure:"ppn" json:"ppn" yaml:"ppn" validate:"required,gt=0"`
	ProfitMargin     float32   `db:"profitMargin" gorm:"not null;index;" mapstructure:"profitMargin" json:"profitMargin" yaml:"profitMargin" validate:"required,gt=0"`
	PackagingID      int       `db:"packagingId" gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" mapstructure:"packagingId" json:"packagingId" validate:"required"`
	TotalPackaging   float32   `db:"totalPackaging" gorm:"not null;index;" mapstructure:"totalPackaging" json:"totalPackaging" yaml:"totalPackaging" validate:"required,gt=0"`
	UnitID           int       `db:"unitId" gorm:"index;not null;" mapstructure:"unitId" json:"unitId" validate:"required"`
	UnitAmount       float32   `db:"unitAmount" gorm:"not null;index;" mapstructure:"unitAmount" json:"unitAmount" yaml:"unitAmount" validate:"required,gt=0"`
	UnitResidu       float32   `db:"unitResidu" gorm:"not null;index;" mapstructure:"unitResidu" json:"unitResidu" yaml:"unitResidu"`

	Packaging Packaging `gorm:"foreignKey:PackagingID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
}

func (p *Product) TableName() string {
	return "products"
}
