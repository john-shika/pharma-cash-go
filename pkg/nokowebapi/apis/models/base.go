package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
	"time"
)

type BaseModel struct {
	ID        uint           `db:"id" gorm:"primaryKey;autoIncrement;" mapstructure:"id" json:"id"`
	UUID      uuid.UUID      `db:"uuid" gorm:"unique;index;not null;" mapstructure:"uuid" json:"uuid"`
	CreatedAt time.Time      `db:"created_at" gorm:"not null;" mapstructure:"created_at" json:"createdAt"`
	UpdatedAt time.Time      `db:"updated_at" gorm:"not null;" mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"null;" mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}

func (BaseModel) TableName() string {
	return "empty"
}

func (b *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	nokocore.KeepVoid(tx)

	if b.UUID == uuid.Nil {
		b.UUID = nokocore.NewUUID()
	}

	timeUtcNow := nokocore.GetTimeUtcNow()
	if b.CreatedAt.IsZero() {
		b.CreatedAt = timeUtcNow
	}

	if b.UpdatedAt.IsZero() {
		b.UpdatedAt = timeUtcNow
	}

	return nil
}
