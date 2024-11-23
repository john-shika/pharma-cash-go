package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
	"time"
)

type BaseModel struct {
	ID        int            `db:"id" gorm:"primaryKey;autoIncrement;" mapstructure:"id" json:"id"`
	UUID      uuid.UUID      `db:"uuid" gorm:"unique;index;not null;" mapstructure:"uuid" json:"uuid"`
	CreatedAt time.Time      `db:"created_at" gorm:"not null;" mapstructure:"created_at" json:"createdAt"`
	UpdatedAt time.Time      `db:"updated_at" gorm:"not null;" mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"null;" mapstructure:"deleted_at" json:"deletedAt"`
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

	return nil
}
