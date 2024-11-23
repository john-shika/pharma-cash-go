package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
	"time"
)

type SessionBody struct {
	UserID    uuid.UUID `json:"userId" yaml:"user_id" form:"user_id" validate:"uuid"`
	TokenId   string    `json:"tokenId" yaml:"token_id" form:"token_id" validate:"uuid"`
	IPAddress string    `json:"ipAddr" yaml:"ip_addr" form:"ip_addr" validate:"ipaddr"`
	UserAgent string    `json:"userAgent" yaml:"user_agent" form:"user_agent" validate:"ascii"`
	Expires   string    `json:"expires" yaml:"expires" form:"expires" validate:"datetime"`
}

func ToSessionModel(session *SessionBody, userId int, expires time.Time) *models.Session {
	return &models.Session{
		UserID:    userId,
		TokenId:   session.TokenId,
		IPAddress: session.IPAddress,
		UserAgent: session.UserAgent,
		Expires:   expires,
	}
}

type SessionResult struct {
	UserID         uuid.UUID  `json:"userId" yaml:"user_id" form:"user_id"`
	TokenId        string     `json:"tokenId" yaml:"token_id" form:"token_id"`
	RefreshTokenId string     `json:"refreshTokenId" yaml:"refresh_token_id" form:"refresh_token_id"`
	IPAddress      string     `json:"ipAddr" yaml:"ip_addr" form:"ip_addr"`
	UserAgent      string     `json:"userAgent" yaml:"user_agent" form:"user_agent"`
	Expires        string     `json:"expires" yaml:"expires" form:"expires"`
	CreatedAt      string     `json:"createdAt" yaml:"created_at" form:"created_at"`
	UpdatedAt      string     `json:"updatedAt" yaml:"updated_at" form:"updated_at"`
	User           UserResult `json:"user" yaml:"user"`
}

func ToSessionResult(session *models.Session, user UserResult) SessionResult {
	return SessionResult{
		UserID:         session.User.UUID,
		TokenId:        session.TokenId,
		RefreshTokenId: session.RefreshTokenId.String,
		IPAddress:      session.IPAddress,
		UserAgent:      session.UserAgent,
		Expires:        nokocore.ToTimeUtcStringISO8601(session.Expires),
		CreatedAt:      nokocore.ToTimeUtcStringISO8601(session.CreatedAt),
		UpdatedAt:      nokocore.ToTimeUtcStringISO8601(session.UpdatedAt),
		User:           user,
	}
}
