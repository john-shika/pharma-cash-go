package repositories

import (
	"errors"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"strings"
)

type BaseRepositoryImpl[T any] interface {
	SafeFind(query string, args ...any) (*T, error)
	Find(query string, args ...any) (*T, error)
	SafeCheck(schema *T, checkHandler CheckHandler[T]) error
	Check(schema *T, checkHandler CheckHandler[T]) error
	Create(schema *T) error
	SafeUpdate(schema *T, query string, args ...any) error
	Update(schema *T, query string, args ...any) error
	SafeDelete(schema *T) error
	Delete(schema *T) error
}

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func NewBaseRepository[T any](DB *gorm.DB) BaseRepository[T] {
	return BaseRepository[T]{
		DB: DB,
	}
}

func (b *BaseRepository[T]) SafeFind(query string, args ...any) (*T, error) {
	if !strings.Contains(query, "deleted_at IS NULL") {
		words := []string{strings.TrimSpace(query), "AND deleted_at IS NULL"}
		query = strings.Join(words, " ")
	}
	return b.Find(query, args...)
}

func (b *BaseRepository[T]) Find(query string, args ...any) (*T, error) {
	var err error
	var schema T
	nokocore.KeepVoid(err, schema)

	//conditions := []any{query}
	//conditions = append(conditions, args...)
	//tx := b.DB.Unscoped().Find(&schema, conditions)
	
	tx := b.DB.Unscoped().Where(query, args...).Find(&schema)
	if err = tx.Error; err != nil {
		return nil, err
	}

	// the schema was initialized but not updated from the database
	identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
	if identity != uuid.Nil {
		return &schema, nil
	}

	return nil, nil
}

type CheckHandler[T any] func(schema *T) error

func (t CheckHandler[T]) Call(schema *T) error {
	return t(schema)
}

func (b *BaseRepository[T]) SafeCheck(schema *T, checkHandler CheckHandler[T]) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	if schema != nil {
		tableName := sqlx.GetTableName(schema)
		id := nokocore.GetValueWithSuperKey(schema, "BaseModel.id").(uint64)
		if id != 0 {
			if check, err = b.SafeFind("id = ?", id); err != nil {
				return errors.New(fmt.Sprintf("failed to search %s", tableName))
			}
		}

		identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
		if identity != uuid.Nil {
			if check, err = b.SafeFind("uuid = ?", identity); err != nil {
				return errors.New(fmt.Sprintf("failed to search %s", tableName))
			}
		}

		if check != nil {
			if checkHandler != nil {
				return checkHandler.Call(check)
			}

			return errors.New(fmt.Sprintf("%s already exists", tableName))
		}

		return nil
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) Check(schema *T, checkHandler CheckHandler[T]) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	if schema != nil {
		tableName := sqlx.GetTableName(schema)
		id := nokocore.GetValueWithSuperKey(schema, "BaseModel.id").(uint64)
		if id != 0 {
			if check, err = b.Find("id = ?", id); err != nil {
				return errors.New(fmt.Sprintf("failed to search %s", tableName))
			}
		}

		identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
		if identity != uuid.Nil {
			if check, err = b.Find("uuid = ?", identity); err != nil {
				return errors.New(fmt.Sprintf("failed to search %s", tableName))
			}
		}

		if check != nil {
			if checkHandler != nil {
				return checkHandler.Call(check)
			}

			return errors.New(fmt.Sprintf("%s already exists", tableName))
		}

		return nil
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) SafeCreate(schema *T) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	if schema != nil {
		tableName := sqlx.GetTableName(schema)
		if err = b.SafeCheck(schema, nil); err != nil {
			return err
		}

		// using mapstructure to inject any values
		err = mapstructure.Decode(nokocore.MapAny{
			"BaseModel": nokocore.MapAny{
				"uuid": nokocore.NewUUID(),
			},
		}, schema)

		if err != nil {
			return errors.New(fmt.Sprintf("failed to inject values into %s", tableName))
		}

		tx := b.DB.Create(schema)
		if err = tx.Error; err != nil {
			return errors.New(fmt.Sprintf("failed to create %s", tableName))
		}

		if tx.RowsAffected > 0 {
			return nil
		}

		return fmt.Errorf("no rows affected in %s", tableName)
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) Create(schema *T) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	if schema != nil {
		tableName := sqlx.GetTableName(schema)
		if err = b.Check(schema, nil); err != nil {
			return err
		}

		// using mapstructure to inject any values
		err = mapstructure.Decode(nokocore.MapAny{
			"BaseModel": nokocore.MapAny{
				"uuid": nokocore.NewUUID(),
			},
		}, schema)

		if err != nil {
			return errors.New(fmt.Sprintf("failed to inject values into %s", tableName))
		}

		tx := b.DB.Unscoped().Create(schema)
		if err = tx.Error; err != nil {
			return errors.New(fmt.Sprintf("failed to create %s", tableName))
		}

		if tx.RowsAffected > 0 {
			return nil
		}

		return fmt.Errorf("no rows affected in %s", tableName)
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) SafeUpdate(schema *T, query string, args ...any) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	if schema != nil {
		tableName := sqlx.GetTableName(schema)
		if check, err = b.SafeFind(query, args...); err != nil {
			return err
		}

		if check != nil {
			id := nokocore.GetValueWithSuperKey(check, "BaseModel.id").(uint64)
			identity := nokocore.GetValueWithSuperKey(check, "BaseModel.uuid").(uuid.UUID)

			// using mapstructure to inject any values
			err = mapstructure.Decode(nokocore.MapAny{
				"BaseModel": nokocore.MapAny{
					"id":   id,
					"uuid": identity,
				},
			}, schema)

			if err != nil {
				return errors.New(fmt.Sprintf("failed to inject values into %s", tableName))
			}

			tx := b.DB.Save(schema)
			if err = tx.Error; err != nil {
				return errors.New(fmt.Sprintf("failed to update %s", tableName))
			}

			if tx.RowsAffected > 0 {
				return nil
			}

			return errors.New(fmt.Sprintf("no rows affected in %s", tableName))
		}

		return errors.New(fmt.Sprintf("%s not found", tableName))
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) Update(schema *T, query string, args ...any) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	if schema != nil {
		tableName := sqlx.GetTableName(schema)
		if check, err = b.Find(query, args...); err != nil {
			return err
		}

		if check != nil {
			id := nokocore.GetValueWithSuperKey(check, "BaseModel.id").(uint64)
			identity := nokocore.GetValueWithSuperKey(check, "BaseModel.uuid").(uuid.UUID)

			// using mapstructure to inject any values
			err = mapstructure.Decode(nokocore.MapAny{
				"BaseModel": nokocore.MapAny{
					"id":   id,
					"uuid": identity,
				},
			}, schema)

			if err != nil {
				return errors.New(fmt.Sprintf("failed to inject values into %s", tableName))
			}

			tx := b.DB.Unscoped().Save(schema)
			if err = tx.Error; err != nil {
				return errors.New(fmt.Sprintf("failed to update %s", tableName))
			}

			if tx.RowsAffected > 0 {
				return nil
			}

			return errors.New(fmt.Sprintf("no rows affected in %s", tableName))
		}

		return errors.New(fmt.Sprintf("%s not found", tableName))
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) SafeDelete(schema *T) error {
	var err error
	nokocore.KeepVoid(err)

	if schema != nil {
		tableName := sqlx.GetTableName(schema)
		timeUtcNow := nokocore.GetTimeUtcNow()

		// using mapstructure to inject any values
		err = mapstructure.Decode(nokocore.MapAny{
			"BaseModel": nokocore.MapAny{
				"deleted_at": sqlx.NewDeletedAt(timeUtcNow),
			},
		}, schema)

		if err != nil {
			return errors.New(fmt.Sprintf("failed to inject values into %s", tableName))
		}

		id := nokocore.GetValueWithSuperKey(schema, "BaseModel.id").(uint64)
		if id != 0 {
			if err = b.Update(schema, "id = ?", id); err != nil {
				return errors.New(fmt.Sprintf("unable to delete %s", tableName))
			}
		}

		identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
		if identity != uuid.Nil {
			if err = b.Update(schema, "uuid = ?", identity); err != nil {
				return errors.New(fmt.Sprintf("unable to delete %s", tableName))
			}
		}

		return nil
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) Delete(schema *T) error {
	var err error
	nokocore.KeepVoid(err)

	if schema != nil {
		tableName := sqlx.GetTableName(schema)
		tx := b.DB.Unscoped().Delete(schema)
		if err = tx.Error; err != nil {
			return errors.New(fmt.Sprintf("unable to delete %s", tableName))
		}

		if tx.RowsAffected > 0 {
			return nil
		}

		return errors.New(fmt.Sprintf("no rows affected in %s", tableName))
	}

	return errors.New("invalid value")
}
