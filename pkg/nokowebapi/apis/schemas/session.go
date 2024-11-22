package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/apis/models"
	"time"
)

type SessionBody struct {
	UserID    uuid.UUID `json:"userId" yaml:"user_id" form:"user_id" validate:"uuid"`
	TokenId   string    `json:"tokenId" yaml:"token_id" form:"token_id" validate:"uuid"`
	IPAddress string    `json:"ipAddr" yaml:"ip_addr" form:"ip_addr" validate:"ipaddr"`
	UserAgent string    `json:"userAgent" yaml:"user_agent" form:"user_agent" validate:"ascii"`
	Expires   string    `json:"expires" yaml:"expires" form:"expires" validate:"datetime"`
	User      UserBody  `json:"-" yaml:"-" validate:"-"`
}

type SessionResp struct {
	UserID         uuid.UUID `json:"userId" yaml:"user_id" form:"user_id"`
	TokenId        string    `json:"tokenId" yaml:"token_id" form:"token_id"`
	RefreshTokenId string    `json:"refreshTokenId" yaml:"refresh_token_id" form:"refresh_token_id"`
	IPAddress      string    `json:"ipAddr" yaml:"ip_addr" form:"ip_addr"`
	UserAgent      string    `json:"userAgent" yaml:"user_agent" form:"user_agent"`
	Expires        string    `json:"expires" yaml:"expires" form:"expires"`
	CreatedAt      time.Time `json:"createdAt" yaml:"created_at" form:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" yaml:"updated_at" form:"updated_at"`
	User           UserResp  `json:"-" yaml:"-"`
}

func ToSessionResp(session *models.Session) SessionResp {
	return SessionResp{
		UserID:         session.User.UUID,
		TokenId:        session.TokenId,
		RefreshTokenId: session.RefreshTokenId.String,
		IPAddress:      session.IPAddress,
		UserAgent:      session.UserAgent,
		Expires:        session.Expires.String(),
		User:           ToUserResp(&session.User),
	}
}
