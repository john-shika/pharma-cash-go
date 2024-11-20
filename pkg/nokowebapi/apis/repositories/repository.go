package repositories

import (
	"errors"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
)

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func NewBaseRepository[T any](DB *gorm.DB) BaseRepository[T] {
	return BaseRepository[T]{
		DB: DB,
	}
}

func (u *BaseRepository[T]) Find(wheres ...any) (*T, error) {
	var err error
	var schema T
	nokocore.KeepVoid(err, schema)

	if err = u.DB.Find(&schema, wheres...).Error; err != nil {
		return nil, err
	}

	// the schema was initialized but not updated from the database
	identity := nokocore.GetValueWithSuperKey(schema, "Model.uuid").(uuid.UUID)
	if identity != uuid.Nil {
		return &schema, nil
	}

	return nil, nil
}

func (u *BaseRepository[T]) Create(schema *T) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	tableName := sqlx.GetTableName(schema)
	identity := nokocore.GetValueWithSuperKey(schema, "Model.uuid").(uuid.UUID)
	if identity != uuid.Nil {
		if check, err = u.Find("uuid = ?", identity); err != nil {
			return err
		}

		if check != nil {
			return errors.New(fmt.Sprintf("%s already exists", tableName))
		}
	}

	// using mapstructure to inject any values

	err = mapstructure.Decode(nokocore.MapAny{
		"Model": models.Model{
			UUID:      nokocore.NewUUID(),
			CreatedAt: nokocore.GetTimeUtcNow(),
			UpdatedAt: nokocore.GetTimeUtcNow(),
			DeletedAt: gorm.DeletedAt{},
		},
	}, schema)

	if err != nil {
		return err
	}

	tx := u.DB.Create(schema)
	if err = tx.Error; err != nil {
		return errors.New(fmt.Sprintf("failed to create %s", tableName))
	}

	if tx.RowsAffected > 0 {
		return nil
	}

	return fmt.Errorf("no rows affected in %s", tableName)
}
