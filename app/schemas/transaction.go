package schemas

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
)

type TransactionBody struct {
	UserID uuid.UUID `mapstructure:"user_id" json:"userId" form:"user_id" validate:"uuid,omitempty"`
	Total  string    `mapstructure:"total" json:"total" form:"total" validate:"decimal,omitempty"`
	Pay    string    `mapstructure:"pay" json:"pay" form:"pay" validate:"decimal"`
}

func ToTransactionModel(transaction *TransactionBody) *models2.Transaction {
	if transaction != nil {
		return &models2.Transaction{
			Total: decimal.RequireFromString(transaction.Total),
			Pay:   decimal.RequireFromString(transaction.Pay),
		}
	}

	return nil
}

type TransactionResult struct {
	UUID      uuid.UUID       `mapstructure:"uuid" json:"uuid"`
	Total     decimal.Decimal `mapstructure:"total" json:"total"`
	Pay       decimal.Decimal `mapstructure:"pay" json:"pay"`
	Exchange  decimal.Decimal `mapstructure:"exchange" json:"exchange"`
	Verified  bool            `mapstructure:"verified" json:"verified"`
	CreatedAt string          `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt string          `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt string          `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}

func ToTransactionResult(transaction *models2.Transaction) TransactionResult {
	if transaction != nil {
		createdAt := nokocore.ToTimeUtcStringISO8601(transaction.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(transaction.UpdatedAt)
		var deletedAt string
		if transaction.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(transaction.DeletedAt.Time)
		}
		return TransactionResult{
			UUID:      transaction.UUID,
			Total:     transaction.Total,
			Pay:       transaction.Pay,
			Exchange:  transaction.Exchange,
			Verified:  transaction.Verified,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		}
	}

	return TransactionResult{}
}
