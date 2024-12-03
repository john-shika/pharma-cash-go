package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
)

type CartBody struct {
	ProductID uuid.UUID `mapstructure:"product_id" json:"productId"`
	Quantity  int       `mapstructure:"quantity" json:"quantity"`
}

func ToCartModel(cart *CartBody) *models2.Cart {
	return &models2.Cart{
		Quantity: cart.Quantity,
	}
}

type CartResult struct {
	UUID      uuid.UUID     `mapstructure:"uuid" json:"uuid"`
	ProductID uuid.UUID     `mapstructure:"product_id" json:"productId"`
	Product   ProductResult `mapstructure:"product" json:"product"`
	Quantity  int           `mapstructure:"quantity" json:"quantity"`
	CreatedAt string        `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt string        `mapstructure:"updated_at" json:"updatedAt"`
}

func ToCartResult(cart *models2.Cart) CartResult {
	if cart != nil {
		createdAt := nokocore.ToTimeUtcStringISO8601(cart.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(cart.UpdatedAt)
		return CartResult{
			UUID:      cart.UUID,
			ProductID: cart.Product.UUID,
			Product:   ToProductResult(&cart.Product),
			Quantity:  cart.Quantity,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
	}

	return CartResult{}
}
