package utils

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
	"time"
)

func CopyBaseModel[T any](schema T, source T) (T, error) {
	err := mapstructure.Decode(nokocore.MapAny{
		"BaseModel": nokocore.MapAny{
			"id":         nokocore.GetValueWithSuperKey(source, "BaseModel.id").(uint),
			"uuid":       nokocore.GetValueWithSuperKey(source, "BaseModel.uuid").(uuid.UUID),
			"created_at": nokocore.GetValueWithSuperKey(source, "BaseModel.created_at").(time.Time),
			"updated_at": nokocore.GetValueWithSuperKey(source, "BaseModel.updated_at").(time.Time),
			"deleted_at": nokocore.GetValueWithSuperKey(source, "BaseModel.deleted_at").(gorm.DeletedAt),
		},
	}, schema)

	return schema, err
}
