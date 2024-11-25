package models

import (
	"database/sql"
	"time"
)

type Session struct {
	BaseModel
	UserID         int            `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`
	TokenId        string         `db:"token_id" gorm:"unique;index;not null;" mapstructure:"token_id" json:"tokenId"`
	RefreshTokenId sql.NullString `db:"refresh_token_id" gorm:"unique;index;null;" mapstructure:"refresh_token_id" json:"refreshTokenId"`
	IPAddress      string         `db:"ip_addr" gorm:"index;not null;" mapstructure:"ip_addr" json:"ipAddr"`
	UserAgent      string         `db:"user_agent" gorm:"index;not null;" mapstructure:"user_agent" json:"userAgent"`
	Expires        time.Time      `db:"expires" gorm:"not null;" mapstructure:"expires" json:"expires"`
	User           User           `db:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"user" json:"user,omitempty"`
}

func (s *Session) TableName() string {
	return "sessions"
}
