package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
	"time"
)

type SessionBody struct {
	UserID    uuid.UUID `mapstructure:"user_id" json:"userId" form:"user_id" validate:"uuid"`
	TokenId   string    `mapstructure:"token_id" json:"tokenId" form:"token_id" validate:"uuid"`
	IPAddress string    `mapstructure:"ip_addr" json:"ipAddr" form:"ip_addr" validate:"ipaddr"`
	UserAgent string    `mapstructure:"user_agent" json:"userAgent" form:"user_agent" validate:"ascii"`
	Expires   string    `mapstructure:"expires" json:"expires" form:"expires" validate:"datetime"`
}

func ToSessionModel(session *SessionBody, userId int, expires time.Time) *models.Session {
	if session != nil {
		return &models.Session{
			UserID:    userId,
			TokenId:   session.TokenId,
			IPAddress: session.IPAddress,
			UserAgent: session.UserAgent,
			Expires:   expires,
		}
	}

	return nil
}

type SessionResult struct {
	UserID         uuid.UUID  `mapstructure:"user_id" json:"userId"`
	TokenId        string     `mapstructure:"token_id" json:"tokenId"`
	RefreshTokenId string     `mapstructure:"refresh_token_id" json:"refreshTokenId,omitempty"`
	IPAddress      string     `mapstructure:"ip_addr" json:"ipAddr"`
	UserAgent      string     `mapstructure:"user_agent" json:"userAgent"`
	Expires        string     `mapstructure:"expires" json:"expires"`
	CreatedAt      string     `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt      string     `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt      string     `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
	User           UserResult `mapstructure:"user" json:"user,omitempty"`
}

func ToSessionResult(session *models.Session, user UserResult) SessionResult {
	var deletedAt string
	if session != nil {
		if session.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(session.DeletedAt.Time)
		}
		return SessionResult{
			UserID:         session.User.UUID,
			TokenId:        session.TokenId,
			RefreshTokenId: session.RefreshTokenId.String,
			IPAddress:      session.IPAddress,
			UserAgent:      session.UserAgent,
			Expires:        nokocore.ToTimeUtcStringISO8601(session.Expires),
			CreatedAt:      nokocore.ToTimeUtcStringISO8601(session.CreatedAt),
			UpdatedAt:      nokocore.ToTimeUtcStringISO8601(session.UpdatedAt),
			DeletedAt:      deletedAt,
			User:           user,
		}
	}

	return SessionResult{}
}
