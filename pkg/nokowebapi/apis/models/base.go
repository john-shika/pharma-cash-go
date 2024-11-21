package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
	"time"
)

type BaseModel struct {
	ID        uint64         `db:"id" gorm:"primaryKey;autoIncrement;" mapstructure:"id" json:"id" yaml:"id"`
	UUID      uuid.UUID      `db:"uuid" gorm:"unique;index;not null;" mapstructure:"uuid" json:"uuid" yaml:"uuid"`
	CreatedAt time.Time      `db:"created_at" gorm:"not null;" mapstructure:"created_at" json:"createdAt" yaml:"created_at"`
	UpdatedAt time.Time      `db:"updated_at" gorm:"not null;" mapstructure:"updated_at" json:"updatedAt" yaml:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:";" mapstructure:"deleted_at" json:"deletedAt" yaml:"deleted_at"`
}

func (m *BaseModel) TableName() string {
	return "empty"
}

func (m *BaseModel) BeforeCreate(db *gorm.DB) (err error) {
	nokocore.KeepVoid(db)

	if m.UUID == uuid.Nil {
		m.UUID = nokocore.NewUUID()
	}

	timeUtcNow := nokocore.GetTimeUtcNow()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = timeUtcNow
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = timeUtcNow
	}

	return
}
