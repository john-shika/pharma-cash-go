package schemas

import "github.com/google/uuid"

type UserBody struct {
	Username string        `json:"username" yaml:"username" form:"username" validate:"required"`
	Password string        `json:"password" yaml:"password" form:"password" validate:"required"`
	Email    string        `json:"email" yaml:"email" form:"email" validate:"omitempty,email"`
	Phone    string        `json:"phone" yaml:"phone" form:"phone" validate:"omitempty,e164"`
	Admin    bool          `json:"admin" yaml:"admin" form:"admin" validate:"omitempty,boolean"`
	Role     string        `json:"role" yaml:"role" form:"role" validate:"omitempty,ascii"`
	Level    int           `json:"level" yaml:"level" form:"level" validate:"omitempty,number"`
	Sessions []SessionBody `json:"sessions" yaml:"sessions"`
}

type SessionBody struct {
	UserID         uuid.UUID `json:"userId" yaml:"user_id" form:"user_id"`
	TokenId        string    `json:"tokenId" yaml:"token_id" form:"token_id"`
	RefreshTokenId string    `json:"refreshTokenId" yaml:"refresh_token_id" form:"refresh_token_id"`
	IPAddress      string    `json:"ipAddr" yaml:"ip_addr" form:"ip_addr"`
	UserAgent      string    `json:"userAgent" yaml:"user_agent" form:"user_agent"`
	Expires        string    `json:"expires" yaml:"expires" form:"expires"`
}
