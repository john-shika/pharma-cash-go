package schemas

import "github.com/google/uuid"

type UserBody struct {
	Username string        `json:"username" yaml:"username" form:"username" validate:"ascii"`
	Password string        `json:"password" yaml:"password" form:"password" validate:"password"`
	Email    string        `json:"email" yaml:"email" form:"email" validate:"email,omitempty"`
	Phone    string        `json:"phone" yaml:"phone" form:"phone" validate:"phone,omitempty"`
	Admin    bool          `json:"admin" yaml:"admin" form:"admin" validate:"boolean,omitempty"`
	Role     string        `json:"role" yaml:"role" form:"role" validate:"ascii,omitempty"`
	Level    int           `json:"level" yaml:"level" form:"level" validate:"number,omitempty"` // FUTURE: can handle min=N,max=N
	Sessions []SessionBody `json:"sessions" yaml:"sessions" validate:"-"`
}

type SessionBody struct {
	UserID         uuid.UUID `json:"userId" yaml:"user_id" form:"user_id"`
	TokenId        string    `json:"tokenId" yaml:"token_id" form:"token_id"`
	RefreshTokenId string    `json:"refreshTokenId" yaml:"refresh_token_id" form:"refresh_token_id"`
	IPAddress      string    `json:"ipAddr" yaml:"ip_addr" form:"ip_addr"`
	UserAgent      string    `json:"userAgent" yaml:"user_agent" form:"user_agent"`
	Expires        string    `json:"expires" yaml:"expires" form:"expires"`
}
