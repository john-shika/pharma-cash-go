package schemas

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
)

type ProductBody struct {
	Barcode          string  `mapstructure:"barcode" json:"barcode" validate:"ascii"`
	Brand            string  `mapstructure:"brand" json:"brand"`
	ProductName      string  `mapstructure:"product_name" json:"productName"`
	Supplier         string  `mapstructure:"supplier" json:"supplier"`
	Description      string  `mapstructure:"description" json:"description"`
	Category         string  `mapstructure:"category" json:"category"`
	Expires          string  `mapstructure:"expires" json:"expires" validate:"dateOnly"`
	PurchasePrice    string  `mapstructure:"purchase_price" json:"purchasePrice" validate:"decimal,min=0"`
	SupplierDiscount float32 `mapstructure:"supplier_discount" json:"supplierDiscount" validate:"numeric,min=0"`
	VAT              float32 `mapstructure:"vat" json:"vat" validate:"numeric,min=0"`
	ProfitMargin     float32 `mapstructure:"profit_margin" json:"profitMargin" validate:"numeric,min=0"`
	PackageID        string  `mapstructure:"package_id" json:"packageId" validate:"uuid"`
	PackageTotal     float32 `mapstructure:"package_total" json:"packageTotal" validate:"number,min=0"`
	UnitID           string  `mapstructure:"unit_id" json:"unitId" validate:"uuid"`
	UnitAmount       float32 `mapstructure:"unit_amount" json:"unitAmount" validate:"number,min=0"`
	UnitExtra        float32 `mapstructure:"unit_extra" json:"unitExtra" validate:"omitempty"`
}

func ToProductModel(productBody *ProductBody, packageModel *models2.Package, unitModel *models2.Unit) *models2.Product {
	return &models2.Product{
		Barcode:          productBody.Barcode,
		Brand:            productBody.Brand,
		ProductName:      productBody.ProductName,
		Supplier:         productBody.Supplier,
		Description:      productBody.Description,
		Category:         productBody.Category,
		Expires:          sqlx.DateOnly(productBody.Expires),
		PurchasePrice:    decimal.RequireFromString(productBody.PurchasePrice),
		SupplierDiscount: productBody.SupplierDiscount,
		VAT:              productBody.VAT,
		ProfitMargin:     productBody.ProfitMargin,
	}
}

type ProductResult struct {
	Barcode          string          `mapstructure:"barcode" json:"barcode"`
	Brand            string          `mapstructure:"brand" json:"brand"`
	ProductName      string          `mapstructure:"product_name" json:"productName"`
	Supplier         string          `mapstructure:"supplier" json:"supplier"`
	Description      string          `mapstructure:"description" json:"description"`
	Category         string          `mapstructure:"category" json:"category"`
	Expires          string          `mapstructure:"expires" json:"expires"`
	PurchasePrice    decimal.Decimal `mapstructure:"purchase_price" json:"purchasePrice"`
	SupplierDiscount float32         `mapstructure:"supplier_discount" json:"supplierDiscount"`
	VAT              float32         `mapstructure:"vat" json:"vat"`
	ProfitMargin     float32         `mapstructure:"profit_margin" json:"profitMargin"`
	PackageId        uuid.UUID       `mapstructure:"package_id" json:"packageId"`
	PackageTotal     float32         `mapstructure:"package_total" json:"packageTotal"`
	UnitID           uuid.UUID       `mapstructure:"unit_id" json:"unitId"`
	UnitAmount       float32         `mapstructure:"unit_amount" json:"unitAmount"`
	UnitExtra        float32         `mapstructure:"unit_extra" json:"unitExtra"`
}

func ToProductResult(product *models2.Product) ProductResult {
	return ProductResult{
		Barcode:          product.Barcode,
		Brand:            product.Brand,
		ProductName:      product.ProductName,
		Supplier:         product.Supplier,
		Description:      product.Description,
		Category:         product.Category,
		Expires:          product.Expires.Format(nokocore.DateOnlyFormat),
		PurchasePrice:    product.PurchasePrice,
		SupplierDiscount: product.SupplierDiscount,
		VAT:              product.VAT,
		ProfitMargin:     product.ProfitMargin,
		PackageId:        product.Package.UUID,
		PackageTotal:     product.PackageTotal,
		UnitID:           product.Unit.UUID,
		UnitAmount:       product.UnitAmount,
		UnitExtra:        product.UnitExtra,
	}
}
