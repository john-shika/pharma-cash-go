package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
	"time"
)

type Model struct {
	ID        uint64         `db:"id" gorm:"primaryKey;autoIncrement" mapstructure:"id" json:"id" yaml:"id"`
	UUID      uuid.UUID      `db:"uuid" gorm:"unique;not null;index" mapstructure:"uuid" json:"uuid" yaml:"uuid"`
	CreatedAt time.Time      `db:"created_at" gorm:"not null" mapstructure:"created_at" json:"createdAt" yaml:"created_at"`
	UpdatedAt time.Time      `db:"updated_at" gorm:"not null" mapstructure:"updated_at" json:"updatedAt" yaml:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"index" mapstructure:"deleted_at" json:"deletedAt" yaml:"deleted_at"`
}

func (Model) TableName() string {
	return "models"
}

func (m *Model) BeforeCreate(db *gorm.DB) (err error) {
	nokocore.KeepVoid(db)

	m.UUID = nokocore.Unwrap(uuid.NewV7())
	return
}
