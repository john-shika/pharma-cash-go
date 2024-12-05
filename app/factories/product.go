package factories

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"nokowebapi/apis/factories"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
)

func ProductFactory(DB *gorm.DB) []any {
	products := []models2.Product{
		{
			Barcode:          "00000001",
			Brand:            "Yamaha",
			ProductName:      "Yamaha F310",
			Supplier:         "Yamaha",
			Description:      "Yamaha F310",
			Expires:          sqlx.ParseDateOnlyNotNull("2023-01-01"),
			PurchasePrice:    decimal.RequireFromString("10000.00"),
			SalePrice:        decimal.RequireFromString("15000.00"),
			SupplierDiscount: 0,
			VAT:              0,
			ProfitMargin:     0,
			PackageID:        1,
			PackageTotal:     1,
			UnitID:           1,
			UnitScale:        1,
			UnitExtra:        0,
			Categories: []models2.Category{
				{
					CategoryName: "Yamaha Guitar",
				},
				{
					CategoryName: "Guitar",
				},
			},
		},
	}

	temp := factories.BaseFactory[models2.Product](DB, products, "brand = ? AND product_name = ? AND supplier = ?", func(product models2.Product) []any {
		return []any{
			product.Brand,
			product.ProductName,
			product.Supplier,
		}
	})

	return nokocore.ToSliceAny(temp)
}
