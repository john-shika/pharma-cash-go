package schemas

import (
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
	"time"

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

type StockOpnameResultCreate struct {
	UUID       uuid.UUID         `mapstructure:"uuid" json:"uuid"`
	SubmitedAt sqlx.NullDateOnly `mapstructure:"submited_at" json:"submitedAt"`
	IsVerified bool              `mapstructure:"is_verified" json:"isVerified"`
	CreatedBy  uuid.UUID         `mapstructure:"created_by" json:"createdBy"`
	CreatedAt  string            `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt  string            `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt  string            `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}

type StockOpnameResultGet struct {
	ProductUUID       uuid.UUID  `json:"productId"`
	Barcode           string     `json:"barcode"`
	ProductName       string     `json:"productName"`
	Brand             string     `json:"brand"`
	PackageTotal      int        `json:"packageTotal"`
	UnitAmount        int        `json:"unitAmount"`
	UnitExtra         int        `json:"unitExtra"`
	UnitTotal         int        `json:"unitTotal"`
	IsMatch           bool       `json:"isMatch"`
	CartStockOpnameId *uuid.UUID `json:"cartStockOpnameId"`
	NotMatchReason    string     `json:"notMatchReason"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
}

type StockOpnameResultGetVerify struct {
	ProductUUID        uuid.UUID  `json:"productId"`
	Barcode            string     `json:"barcode"`
	ProductName        string     `json:"productName"`
	Brand              string     `json:"brand"`
	SystemPackageTotal int        `json:"systemPackageTotal"`
	SystemUnitScale    int        `json:"systemUnitScale"`
	SystemUnitExtra    int        `json:"systemUnitExtra"`
	SystemUnitTotal    int        `json:"systemUnitTotal"`
	IsMatch            bool       `json:"isMatch"`
	NotMatchReason     string     `json:"notMatchReason"`
	RealPackageTotal   int        `json:"realPackageTotal"`
	RealUnitExtra      int        `json:"realUnitExtra"`
	RealUnitTotal      int        `json:"realUnitTotal"`
	CreatedAt          time.Time  `json:"createdAt"`
	UpdatedAt          time.Time  `json:"updatedAt"`
}

type CartVerificationOpnameResult struct {
	CartVerificationOpnameId uuid.UUID     `mapstructure:"cart_verification_opname_id" json:"cartVerificationOpnameId"`
	ProductId                uuid.UUID     `mapstructure:"product_id" json:"productId"`
	NotMatchReason           string        `mapstructure:"not_match_reason" json:"notMatchReason"`
	PackageTotal             int           `mapstructure:"package_total" json:"packageTotal"`
	UnitScale                int           `mapstructure:"unit_scale" json:"unitScale"`
	UnitExtra                int           `mapstructure:"unit_extra" json:"unitExtra"`
	UnitTotal                int           `mapstructure:"unit_total" json:"unitTotal"`
	IsMatch                  bool          `mapstructure:"is_match" json:"isMatch"`
	Warehouse                WarehouseInfo `mapstructure:"warehouse" json:"warehouse"`

	CreatedBy uuid.UUID `mapstructure:"created_by" json:"createdBy"`
	CreatedAt string    `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt string    `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt string    `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}
type WarehouseInfo struct {
	RealPackageTotal int    `mapstructure:"real_package_total" json:"realPackageTotal"`
	RealUnitExtra    int    `mapstructure:"real_unit_extra" json:"realUnitExtra"`
	RealUnitTotal    int    `mapstructure:"real_unit_total" json:"realUnitTotal"`
	NotMatchReason   string `mapstructure:"not_match_reason" json:"notMatchReason"`
}

func ToStockOpnameResultCreate(stockOpname *models2.StockOpname) StockOpnameResultCreate {
	if stockOpname != nil {
		createdAt := nokocore.ToTimeUtcStringISO8601(stockOpname.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(stockOpname.UpdatedAt)
		var deletedAt string
		if stockOpname.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(stockOpname.DeletedAt)
		}
		return StockOpnameResultCreate{
			UUID:       stockOpname.UUID,
			SubmitedAt: stockOpname.SubmitedAt,
			IsVerified: stockOpname.IsVerified,
			CreatedBy:  stockOpname.User.UUID,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
			DeletedAt:  deletedAt,
		}
	}

	return StockOpnameResultCreate{}
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
			CartVerificationOpnameId: cartVerificationOpname.UUID,
			ProductId:                cartVerificationOpname.Product.UUID,
			IsMatch:                  cartVerificationOpname.IsMatch,
			NotMatchReason:           cartVerificationOpname.NotMatchReason,
			PackageTotal:             cartVerificationOpname.Product.PackageTotal,
			UnitScale:                cartVerificationOpname.Product.UnitAmount,
			UnitExtra:                cartVerificationOpname.Product.UnitExtra,
			UnitTotal:                (cartVerificationOpname.Product.PackageTotal * cartVerificationOpname.Product.UnitAmount) + cartVerificationOpname.Product.UnitExtra,
			Warehouse: WarehouseInfo{
				RealPackageTotal: cartVerificationOpname.RealPackageTotal,
				RealUnitExtra:    cartVerificationOpname.RealUnitExtra,
				RealUnitTotal:    (cartVerificationOpname.RealPackageTotal * cartVerificationOpname.Product.UnitAmount) + cartVerificationOpname.RealUnitExtra,
				NotMatchReason:   cartVerificationOpname.NotMatchReason,
			},
			CreatedBy: cartVerificationOpname.User.UUID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		}
	}

	return CartVerificationOpnameResult{}
}
