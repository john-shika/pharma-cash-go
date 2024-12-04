package schemas

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
)

type TransactionBody struct {
	UserID uuid.UUID `mapstructure:"user_id" json:"userId" form:"user_id" validate:"uuid"`
	Pay    string    `mapstructure:"pay" json:"pay" form:"pay" validate:"decimal"`
	Signed bool      `mapstructure:"signed" json:"signed" form:"signed" validate:"boolean,omitempty"`
	Closed bool      `mapstructure:"closed" json:"closed" form:"closed" validate:"boolean,omitempty"`
}

func ToTransactionModel(transaction *TransactionBody) *models2.Transaction {
	if transaction != nil {
		return &models2.Transaction{
			Pay:    decimal.RequireFromString(transaction.Pay),
			Signed: transaction.Signed,
			Closed: transaction.Closed,
		}
	}

	return nil
}

type TransactionResult struct {
	UUID      uuid.UUID       `mapstructure:"uuid" json:"uuid"`
	Pay       decimal.Decimal `mapstructure:"pay" json:"pay"`
	Signed    bool            `mapstructure:"signed" json:"signed"`
	Closed    bool            `mapstructure:"closed" json:"closed"`
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
			Pay:       transaction.Pay,
			Signed:    transaction.Signed,
			Closed:    transaction.Closed,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		}
	}

	return TransactionResult{}
}
