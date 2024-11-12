package models

import (
	"database/sql"
	"time"
)

type Session struct {
	Model
	UUID           string         `db:"uuid" gorm:"unique;not null;index" mapstructure:"uuid" json:"uuid" yaml:"uuid"`
	UserID         uint           `db:"user_id" gorm:"not null;index" mapstructure:"user_id" json:"userId,required" yaml:"user_id"`
	TokenId        string         `db:"token_id" gorm:"index" mapstructure:"token_id" json:"tokenId,omitempty" yaml:"token_id"`
	RefreshTokenId sql.NullString `db:"refresh_token_id" gorm:"index" mapstructure:"refresh_token_id" json:"refreshTokenId,omitempty" yaml:"refresh_token_id"`
	IPAddress      string         `db:"ip_addr" gorm:"index" mapstructure:"ip_addr" json:"ipAddr,omitempty" yaml:"ip_addr"`
	UserAgent      string         `db:"user_agent" gorm:"index" mapstructure:"user_agent" json:"userAgent,omitempty" yaml:"user_agent"`
	ExpiredAt      time.Time      `db:"expired_at" gorm:"not null" mapstructure:"expired_at" json:"expiredAt" yaml:"expired_at"`
}

func (Session) TableName() string {
	return "sessions"
}
