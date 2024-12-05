package schemas

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
	utils2 "pharma-cash-go/app/utils"
)

type CartBody struct {
	UserID        uuid.UUID `mapstructure:"user_id" json:"userId" form:"user_id" validate:"uuid,omitempty"`
	ProductID     uuid.UUID `mapstructure:"product_id" json:"productId" form:"product_id" validate:"uuid"`
	TransactionID uuid.UUID `mapstructure:"transaction_id" json:"transactionId" form:"transaction_id" validate:"uuid,omitempty"`
	PackageTotal  int       `mapstructure:"package_total" json:"packageTotal" form:"package_total" validate:"number,omitempty"`
	UnitExtra     int       `mapstructure:"unit_extra" json:"unitExtra" form:"unit_extra" validate:"number,omitempty"`
}

func ToCartModel(cart *CartBody) *models2.Cart {
	if cart != nil {
		return &models2.Cart{
			PackageTotal: cart.PackageTotal,
			UnitExtra:    cart.UnitExtra,
		}
	}

	return nil
}

func ToCartModelWithProductModel(cart *CartBody, product *models2.Product) *models2.Cart {
	if cart != nil && product != nil {
		packageTotal := cart.PackageTotal
		unitExtra := cart.UnitExtra

		extra, div := utils2.Modulo(unitExtra, product.UnitScale)
		packageTotal += div
		unitExtra = extra

		return &models2.Cart{
			ProductID:    product.ID,
			PackageTotal: packageTotal,
			UnitExtra:    unitExtra,
			Closed:       false,
		}
	}

	return nil
}

type CartResult struct {
	UUID         uuid.UUID       `mapstructure:"uuid" json:"uuid"`
	ProductID    uuid.UUID       `mapstructure:"product_id" json:"productId"`
	Product      ProductResult   `mapstructure:"product" json:"product"`
	PackageTotal int             `mapstructure:"package_total" json:"packageTotal"`
	UnitExtra    int             `mapstructure:"unit_extra" json:"unitExtra"`
	SubTotal     decimal.Decimal `mapstructure:"sub_total" json:"subTotal"`
	Closed       bool            `mapstructure:"closed" json:"closed"`
	CreatedAt    string          `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt    string          `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt    string          `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}

func ToCartResult(cart *models2.Cart) CartResult {
	if cart != nil {
		createdAt := nokocore.ToTimeUtcStringISO8601(cart.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(cart.UpdatedAt)
		var deletedAt string
		if cart.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(cart.DeletedAt.Time)
		}
		return CartResult{
			UUID:         cart.UUID,
			ProductID:    cart.Product.UUID,
			Product:      ToProductResult(&cart.Product),
			PackageTotal: cart.PackageTotal,
			UnitExtra:    cart.UnitExtra,
			SubTotal:     cart.SubTotal,
			Closed:       cart.Closed,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
			DeletedAt:    deletedAt,
		}
	}

	return CartResult{}
}
