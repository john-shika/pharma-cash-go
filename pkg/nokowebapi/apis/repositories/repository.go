package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
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

	return &schema, nil
}

func (u *BaseRepository[T]) Create(schema *T) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	tableName := sqlx.GetTableName(schema)
	uuid := nokocore.GetValueWithSuperKey(schema, "Model.uuid")

	if check, err = u.Find("uuid = ?", uuid); err != nil {
		return err
	}

	if check != nil {
		return errors.New(fmt.Sprintf("%s already exists", tableName))
	}

	if err = u.DB.Create(schema).Error; err != nil {
		return errors.New(fmt.Sprintf("failed to create %s", tableName))
	}

	return nil
}
