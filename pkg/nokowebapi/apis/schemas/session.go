package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
)

type SessionBody struct {
	UserID    uuid.UUID `mapstructure:"user_id" json:"userId" form:"user_id" validate:"uuid"`
	TokenId   string    `mapstructure:"token_id" json:"tokenId" form:"token_id" validate:"uuid"`
	IPAddress string    `mapstructure:"ip_addr" json:"ipAddr" form:"ip_addr" validate:"ipaddr"`
	UserAgent string    `mapstructure:"user_agent" json:"userAgent" form:"user_agent" validate:"ascii"`
	Expires   string    `mapstructure:"expires" json:"expires" form:"expires" validate:"datetimeISO8601"`
}

func ToSessionModel(session *SessionBody) *models.Session {
	if session != nil {
		expires := nokocore.Unwrap(nokocore.ParseTimeUtcByStringISO8601(session.Expires))
		return &models.Session{
			TokenID:   session.TokenId,
			IPAddress: session.IPAddress,
			UserAgent: session.UserAgent,
			Expires:   expires,
		}
	}

	return nil
}

type SessionResult struct {
	UUID           uuid.UUID `mapstructure:"uuid" json:"uuid"`
	UserID         uuid.UUID `mapstructure:"user_id" json:"userId"`
	TokenId        string    `mapstructure:"token_id" json:"tokenId"`
	RefreshTokenId string    `mapstructure:"refresh_token_id" json:"refreshTokenId,omitempty"`
	IPAddress      string    `mapstructure:"ip_addr" json:"ipAddr"`
	UserAgent      string    `mapstructure:"user_agent" json:"userAgent"`
	Expires        string    `mapstructure:"expires" json:"expires"`
	CreatedAt      string    `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt      string    `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt      string    `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
	Used           bool      `mapstructure:"used" json:"used"`
}

func ToSessionResult(session *models.Session) SessionResult {
	if session != nil {
		expires := nokocore.ToTimeUtcStringISO8601(session.Expires)
		createdAt := nokocore.ToTimeUtcStringISO8601(session.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(session.UpdatedAt)
		var deletedAt string
		if session.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(session.DeletedAt.Time)
		}
		return SessionResult{
			UUID:           session.UUID,
			UserID:         session.User.UUID,
			TokenId:        session.TokenID,
			RefreshTokenId: session.RefreshTokenID.String,
			IPAddress:      session.IPAddress,
			UserAgent:      session.UserAgent,
			Expires:        expires,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			DeletedAt:      deletedAt,
		}
	}

	return SessionResult{}
}
