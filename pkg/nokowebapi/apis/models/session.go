package models

import (
	"database/sql"
	"time"
)

type Session struct {
	Model
	UserID         uint64         `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId,required" yaml:"user_id"`
	TokenId        string         `db:"token_id" gorm:"index;" mapstructure:"token_id" json:"tokenId,omitempty" yaml:"token_id"`
	RefreshTokenId sql.NullString `db:"refresh_token_id" gorm:";" mapstructure:"refresh_token_id" json:"refreshTokenId,omitempty" yaml:"refresh_token_id"`
	IPAddress      string         `db:"ip_addr" gorm:"index;" mapstructure:"ip_addr" json:"ipAddr,omitempty" yaml:"ip_addr"`
	UserAgent      string         `db:"user_agent" gorm:"index;" mapstructure:"user_agent" json:"userAgent,omitempty" yaml:"user_agent"`
	Expires        time.Time      `db:"expires" gorm:"not null;" mapstructure:"expires" json:"expires" yaml:"expires"`
	User           User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"-" json:"-" yaml:"-"`
}

func (Session) TableName() string {
	return "sessions"
}
