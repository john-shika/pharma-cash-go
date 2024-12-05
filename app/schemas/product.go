package schemas

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
	utils2 "pharma-cash-go/app/utils"
)

type ProductBody struct {
	Barcode          string   `mapstructure:"barcode" json:"barcode" form:"barcode" validate:"ascii"`
	Brand            string   `mapstructure:"brand" json:"brand" form:"brand"`
	ProductName      string   `mapstructure:"product_name" json:"productName" form:"product_name"`
	Supplier         string   `mapstructure:"supplier" json:"supplier" form:"supplier"`
	Description      string   `mapstructure:"description" json:"description" form:"description" validate:"ascii,omitempty"`
	Expires          string   `mapstructure:"expires" json:"expires" form:"expires" validate:"dateOnly"`
	PurchasePrice    string   `mapstructure:"purchase_price" json:"purchasePrice" form:"purchase_price" validate:"decimal"`
	SupplierDiscount int      `mapstructure:"supplier_discount" json:"supplierDiscount" form:"supplier_discount" validate:"numeric"`
	VAT              int      `mapstructure:"vat" json:"tax" form:"tax" validate:"numeric"` // tax
	ProfitMargin     int      `mapstructure:"profit_margin" json:"profitMargin" form:"profit_margin" validate:"numeric"`
	PackageID        string   `mapstructure:"package_id" json:"packageId" form:"package_id" validate:"uuid,omitempty"`
	PackageType      string   `mapstructure:"package_type" json:"packageType" form:"package_type" validate:"omitempty"`
	PackageTotal     int      `mapstructure:"package_total" json:"packageTotal" form:"package_total" validate:"number"`
	UnitID           string   `mapstructure:"unit_id" json:"unitId" form:"unit_id" validate:"uuid,omitempty"`
	UnitType         string   `mapstructure:"unit_type" json:"unitType" form:"unit_type" validate:"omitempty"`
	UnitScale        int      `mapstructure:"unit_scale" json:"unitScale" form:"unit_scale" validate:"number,min=1"`
	UnitExtra        int      `mapstructure:"unit_extra" json:"unitExtra" form:"unit_extra" validate:"number"`
	Categories       []string `mapstructure:"categories" json:"categories" form:"categories" validate:"ascii,omitempty"`
	Category         string   `mapstructure:"category" json:"category" form:"category" validate:"ascii,omitempty"`
}

func ToProductModel(product *ProductBody) *models2.Product {
	if product != nil {
		var categories []models2.Category
		for i, category := range product.Categories {
			nokocore.KeepVoid(i)
			if category = nokocore.ToPascalCase(category); category != "" {
				categoryModel := models2.Category{
					CategoryName: category,
				}
				categories = append(categories, categoryModel)
			}
		}
		if category := nokocore.ToPascalCase(product.Category); category != "" {
			categoryModel := models2.Category{
				CategoryName: category,
			}
			categories = append(categories, categoryModel)
		}
		discount := float64(product.SupplierDiscount) / 100
		vat := float64(product.VAT) / 100
		margin := float64(product.ProfitMargin) / 100

		extra, div := utils2.Modulo(product.UnitExtra, product.UnitScale)
		product.PackageTotal += div
		product.UnitExtra = extra

		return &models2.Product{
			Barcode:          product.Barcode,
			Brand:            product.Brand,
			ProductName:      product.ProductName,
			Supplier:         product.Supplier,
			Description:      product.Description,
			Expires:          sqlx.ParseDateOnlyNotNull(product.Expires),
			PurchasePrice:    decimal.RequireFromString(product.PurchasePrice),
			SupplierDiscount: discount,
			VAT:              vat,
			ProfitMargin:     margin,
			PackageTotal:     product.PackageTotal,
			UnitScale:        product.UnitScale,
			UnitExtra:        product.UnitExtra,
			Categories:       categories,
		}
	}

	return nil
}

type ProductResult struct {
	UUID             uuid.UUID       `mapstructure:"uuid" json:"uuid"`
	Barcode          string          `mapstructure:"barcode" json:"barcode"`
	Brand            string          `mapstructure:"brand" json:"brand"`
	ProductName      string          `mapstructure:"product_name" json:"productName"`
	Supplier         string          `mapstructure:"supplier" json:"supplier"`
	Description      string          `mapstructure:"description" json:"description"`
	Expires          string          `mapstructure:"expires" json:"expires"`
	PurchasePrice    decimal.Decimal `mapstructure:"purchase_price" json:"purchasePrice"`
	SalePrice        decimal.Decimal `mapstructure:"sale_price" json:"salePrice"`
	SupplierDiscount int             `mapstructure:"supplier_discount" json:"supplierDiscount"`
	VAT              int             `mapstructure:"vat" json:"tax"` // tax
	ProfitMargin     int             `mapstructure:"profit_margin" json:"profitMargin"`
	PackageId        uuid.UUID       `mapstructure:"package_id" json:"packageId"`
	PackageType      string          `mapstructure:"package_type" json:"packageType"`
	PackageTotal     int             `mapstructure:"package_total" json:"packageTotal"`
	UnitID           uuid.UUID       `mapstructure:"unit_id" json:"unitId"`
	UnitType         string          `mapstructure:"unit_type" json:"unitType"`
	UnitScale        int             `mapstructure:"unit_scale" json:"unitScale"`
	UnitExtra        int             `mapstructure:"unit_extra" json:"unitExtra"`
	UnitTotal        int             `mapstructure:"unit_total" json:"unitTotal"`
	CreatedAt        string          `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt        string          `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt        string          `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
	Categories       []string        `mapstructure:"categories" json:"categories"`
	Category         string          `mapstructure:"category" json:"category"`
}

func ToProductResult(product *models2.Product) ProductResult {
	if product != nil {
		var categories []string
		for i, category := range product.Categories {
			nokocore.KeepVoid(i)
			categories = append(categories, category.CategoryName)
		}
		var category string
		if len(categories) > 0 {
			category = categories[0]
		}
		createdAt := nokocore.ToTimeUtcStringISO8601(product.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(product.UpdatedAt)
		var deletedAt string
		if product.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(product.DeletedAt.Time)
		}
		discount := product.SupplierDiscount * 100
		vat := product.VAT * 100
		margin := product.ProfitMargin * 100

		unitTotal := product.UnitScale * product.PackageTotal
		unitTotal += product.UnitExtra

		return ProductResult{
			UUID:             product.UUID,
			Barcode:          product.Barcode,
			Brand:            product.Brand,
			ProductName:      product.ProductName,
			Supplier:         product.Supplier,
			Description:      product.Description,
			Expires:          product.Expires.Format(nokocore.DateOnlyFormat),
			PurchasePrice:    product.PurchasePrice,
			SalePrice:        product.SalePrice,
			SupplierDiscount: int(discount),
			VAT:              int(vat),
			ProfitMargin:     int(margin),
			PackageId:        product.Package.UUID,
			PackageType:      product.Package.PackageType,
			PackageTotal:     product.PackageTotal,
			UnitID:           product.Unit.UUID,
			UnitType:         product.Unit.UnitType,
			UnitScale:        product.UnitScale,
			UnitExtra:        product.UnitExtra,
			UnitTotal:        unitTotal,
			CreatedAt:        createdAt,
			UpdatedAt:        updatedAt,
			DeletedAt:        deletedAt,
			Categories:       categories,
			Category:         category,
		}
	}

	return ProductResult{}
}
