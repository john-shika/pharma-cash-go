package schemas

import (
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"

	"github.com/google/uuid"
)

type StockOpnameBody struct {
	UnitType string `mapstructure:"unit_type" json:"unitType" form:"unit_type" validate:"ascii"`
}

type CartVerificationOpnameBody struct {
	NotMatchReason   string `mapstructure:"not_match_reason" json:"notMatchReason" form:"not_match_reason" validate:"ascii"`
	RealPackageTotal int    `mapstructure:"real_package_total" json:"realPackageTotal" form:"real_package_total" validate:"number"`
	RealUnitExtra    int    `mapstructure:"real_unit_extra" json:"realUnitExtra" form:"real_unit_extra" validate:"number"`
}

func ToStockOpnameModel(unit *StockOpnameBody) *models2.Unit {
	if unit != nil {
		return &models2.Unit{
			UnitType: unit.UnitType,
		}
	}

	return nil
}

type StockOpnameResult struct {
	UUID       uuid.UUID         `mapstructure:"uuid" json:"uuid"`
	SubmitedAt sqlx.NullDateOnly `mapstructure:"submited_at" json:"submitedAt"`
	IsVerified bool              `mapstructure:"is_verified" json:"isVerified"`
	CreatedBy  uuid.UUID         `mapstructure:"created_by" json:"createdBy"`
	CreatedAt  string            `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt  string            `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt  string            `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}

type CartVerificationOpnameResult struct {
	UUID             uuid.UUID `mapstructure:"uuid" json:"uuid"`
	ProductId        uuid.UUID `mapstructure:"product_id" json:"productId"`
	NotMatchReason   string    `mapstructure:"not_match_reason" json:"notMatchReason"`
	IsMatch          bool      `mapstructure:"is_match" json:"isMatch"`
	RealPackageTotal int       `mapstructure:"real_package_total" json:"realPackageTotal"`
	RealUnitExtra    int       `mapstructure:"real_unit_extra" json:"realUnitExtra"`
	RealUnitTotal    int       `mapstructure:"real_unit_total" json:"realUnitTotal"`
	CreatedBy        uuid.UUID `mapstructure:"created_by" json:"createdBy"`
	CreatedAt        string    `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt        string    `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt        string    `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}

func ToStockOpnameResult(stockOpname *models2.StockOpname) StockOpnameResult {
	if stockOpname != nil {
		createdAt := nokocore.ToTimeUtcStringISO8601(stockOpname.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(stockOpname.UpdatedAt)
		var deletedAt string
		if stockOpname.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(stockOpname.DeletedAt)
		}
		return StockOpnameResult{
			UUID:       stockOpname.UUID,
			SubmitedAt: stockOpname.SubmitedAt,
			IsVerified: stockOpname.IsVerified,
			CreatedBy:  stockOpname.User.UUID,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
			DeletedAt:  deletedAt,
		}
	}

	return StockOpnameResult{}
}

func ToCartVerificationOpnameResult(cartVerificationOpname *models2.CartVerificationOpname) CartVerificationOpnameResult {
	if cartVerificationOpname != nil {
		createdAt := nokocore.ToTimeUtcStringISO8601(cartVerificationOpname.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(cartVerificationOpname.UpdatedAt)
		var deletedAt string
		if cartVerificationOpname.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(cartVerificationOpname.DeletedAt)
		}
		return CartVerificationOpnameResult{
			UUID:             cartVerificationOpname.UUID,
			ProductId:        cartVerificationOpname.Product.UUID,
			NotMatchReason:   cartVerificationOpname.NotMatchReason,
			IsMatch:          cartVerificationOpname.IsMatch,
			RealPackageTotal: cartVerificationOpname.RealPackageTotal,
			RealUnitExtra:    cartVerificationOpname.RealUnitExtra,
			RealUnitTotal:    cartVerificationOpname.RealUnitTotal,
			CreatedBy:        cartVerificationOpname.User.UUID,
			CreatedAt:        createdAt,
			UpdatedAt:        updatedAt,
			DeletedAt:        deletedAt,
		}
	}

	return CartVerificationOpnameResult{}
}
