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
	VAT              float32 `mapstructure:"vat" json:"tax" validate:"numeric,min=0"` // input "tax"
	ProfitMargin     float32 `mapstructure:"profit_margin" json:"profitMargin" validate:"numeric,min=0"`
	PackageID        string  `mapstructure:"package_id" json:"packageId" validate:"uuid"`
	PackageTotal     float32 `mapstructure:"package_total" json:"packageTotal" validate:"number,min=0"`
	UnitID           string  `mapstructure:"unit_id" json:"unitId" validate:"uuid"`
	UnitAmount       float32 `mapstructure:"unit_amount" json:"unitAmount" validate:"number,min=0"`
	UnitExtra        float32 `mapstructure:"unit_extra" json:"unitExtra" validate:"omitempty"`
}

func ToProductModel(product *ProductBody, packageModel *models2.Package, unitModel *models2.Unit) *models2.Product {
	if product != nil {
		return &models2.Product{
			Barcode:          product.Barcode,
			Brand:            product.Brand,
			ProductName:      product.ProductName,
			Supplier:         product.Supplier,
			Description:      product.Description,
			Category:         product.Category,
			Expires:          sqlx.ParseDateOnlyNotNull(product.Expires),
			PurchasePrice:    decimal.RequireFromString(product.PurchasePrice),
			SupplierDiscount: product.SupplierDiscount,
			VAT:              product.VAT,
			ProfitMargin:     product.ProfitMargin,
			PackageID:        packageModel.ID,
			PackageTotal:     product.PackageTotal,
			UnitID:           unitModel.ID,
			UnitAmount:       product.UnitAmount,
			UnitExtra:        product.UnitExtra,
			Package:          *packageModel,
			Unit:             *unitModel,
		}
	}

	return nil
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
	VAT              float32         `mapstructure:"vat" json:"tax"` // output "tax"
	ProfitMargin     float32         `mapstructure:"profit_margin" json:"profitMargin"`
	PackageId        uuid.UUID       `mapstructure:"package_id" json:"packageId"`
	PackageTotal     float32         `mapstructure:"package_total" json:"packageTotal"`
	UnitID           uuid.UUID       `mapstructure:"unit_id" json:"unitId"`
	UnitAmount       float32         `mapstructure:"unit_amount" json:"unitAmount"`
	UnitExtra        float32         `mapstructure:"unit_extra" json:"unitExtra"`
}

func ToProductResult(product *models2.Product) ProductResult {
	if product != nil {
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

	return ProductResult{}
}
