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
	SafeFirst(query string, args ...any) (*T, error)
	First(query string, args ...any) (*T, error)
	SafeMany(query string, args ...any) ([]T, error)
	Many(query string, args ...any) ([]T, error)
	SafePreFirst(preloads []string, query string, args ...any) (*T, error)
	PreFirst(preloads []string, query string, args ...any) (*T, error)
	SafePreMany(preloads []string, query string, args ...any) ([]T, error)
	PreMany(preloads []string, query string, args ...any) ([]T, error)
	SafeCheck(schema *T, checkHandler CheckHandler[T]) error
	Check(schema *T, checkHandler CheckHandler[T]) error
	SafeCreate(schema *T) error
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

// careful with SQL injection, can be conflict, you must be known what you are doing!
func injectQuery(query string, conditions ...string) string {
	var ok bool
	nokocore.KeepVoid(ok)

	query = strings.TrimSpace(strings.Split(query, ";")[0])
	for i, condition := range conditions {
		nokocore.KeepVoid(i)

		condition = strings.TrimSpace(condition)
		if !strings.Contains(query, condition) {
			if query != "" {
				query = fmt.Sprintf("%s  %s", query, condition)
				continue
			}

			query = condition
		}
	}

	return query
}

func (b *BaseRepository[T]) SafeFirst(query string, args ...any) (*T, error) {
	query = injectQuery(query, "AND deleted_at IS NULL", "LIMIT 1")
	return b.First(query, args...)
}

func (b *BaseRepository[T]) First(query string, args ...any) (*T, error) {
	var err error
	var schema T
	nokocore.KeepVoid(err, schema)

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

func (b *BaseRepository[T]) SafeMany(query string, args ...any) ([]T, error) {
	query = injectQuery(query, "AND deleted_at IS NULL")
	return b.Many(query, args...)
}

func (b *BaseRepository[T]) Many(query string, args ...any) ([]T, error) {
	var err error
	var schemas []T
	nokocore.KeepVoid(err, schemas)

	tx := b.DB.Unscoped().Where(query, args...).Find(&schemas)
	if err = tx.Error; err != nil {
		return nil, err
	}

	temp := make([]T, 0)
	for i, schema := range schemas {
		nokocore.KeepVoid(i)

		// the schemas were initialized but not updated from the database
		identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
		if identity != uuid.Nil {
			temp = append(temp, schema)
		}
	}

	return temp, nil
}

func (b *BaseRepository[T]) SafePreFirst(preloads []string, query string, args ...any) (*T, error) {
	query = injectQuery(query, "AND deleted_at IS NULL", "LIMIT 1")
	return b.PreFirst(preloads, query, args...)
}

func (b *BaseRepository[T]) PreFirst(preloads []string, query string, args ...any) (*T, error) {
	var err error
	var schema T
	nokocore.KeepVoid(err, schema)

	unscoped := b.DB.Unscoped()
	for i, preload := range preloads {
		nokocore.KeepVoid(i)

		unscoped = unscoped.Preload(preload)
	}

	tx := unscoped.Where(query, args...).Find(&schema)
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

func (b *BaseRepository[T]) SafePreMany(preloads []string, query string, args ...any) ([]T, error) {
	query = injectQuery(query, "AND deleted_at IS NULL")
	return b.PreMany(preloads, query, args...)
}

func (b *BaseRepository[T]) PreMany(preloads []string, query string, args ...any) ([]T, error) {
	var err error
	var schemas []T
	nokocore.KeepVoid(err, schemas)

	unscoped := b.DB.Unscoped()
	for i, preload := range preloads {
		nokocore.KeepVoid(i)

		unscoped = unscoped.Preload(preload)
	}

	tx := unscoped.Where(query, args...).Find(&schemas)
	if err = tx.Error; err != nil {
		return nil, err
	}

	temp := make([]T, 0)
	for i, schema := range schemas {
		nokocore.KeepVoid(i)

		// the schemas were initialized but not updated from the database
		identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
		if identity != uuid.Nil {
			temp = append(temp, schema)
		}
	}

	return temp, nil
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
			if check, err = b.SafeFirst("id = ?", id); err != nil {
				return errors.New(fmt.Sprintf("failed to search %s", tableName))
			}
		}

		identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
		if identity != uuid.Nil {
			if check, err = b.SafeFirst("uuid = ?", identity); err != nil {
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
			if check, err = b.First("id = ?", id); err != nil {
				return errors.New(fmt.Sprintf("failed to search %s", tableName))
			}
		}

		identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
		if identity != uuid.Nil {
			if check, err = b.First("uuid = ?", identity); err != nil {
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

func (b *BaseRepository[T]) baseInit(schema *T) error {
	var err error
	nokocore.KeepVoid(err)

	tableName := sqlx.GetTableName(schema)
	identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)

	// using mapstructure to inject any values
	if identity == uuid.Nil {
		timeUtcNow := nokocore.GetTimeUtcNow()
		identity = nokocore.NewUUID()

		err = mapstructure.Decode(&nokocore.MapAny{
			"BaseModel": nokocore.MapAny{
				"uuid":       identity,
				"created_at": timeUtcNow,
				"updated_at": timeUtcNow,
			},
		}, schema)

		if err != nil {
			return errors.New(fmt.Sprintf("failed to inject values into %s", tableName))
		}
	}

	return nil
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

		if err = b.baseInit(schema); err != nil {
			return err
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

		if err = b.baseInit(schema); err != nil {
			return err
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
		if check, err = b.SafeFirst(query, args...); err != nil {
			return err
		}

		if check != nil {
			id := nokocore.GetValueWithSuperKey(check, "BaseModel.id").(uint64)
			identity := nokocore.GetValueWithSuperKey(check, "BaseModel.uuid").(uuid.UUID)
			timeUtcNow := nokocore.GetTimeUtcNow()

			// using mapstructure to inject any values
			err = mapstructure.Decode(nokocore.MapAny{
				"BaseModel": nokocore.MapAny{
					"id":         id,
					"uuid":       identity,
					"updated_at": timeUtcNow,
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
		if check, err = b.First(query, args...); err != nil {
			return err
		}

		if check != nil {
			id := nokocore.GetValueWithSuperKey(check, "BaseModel.id").(uint64)
			identity := nokocore.GetValueWithSuperKey(check, "BaseModel.uuid").(uuid.UUID)
			timeUtcNow := nokocore.GetTimeUtcNow()

			// using mapstructure to inject any values
			err = mapstructure.Decode(nokocore.MapAny{
				"BaseModel": nokocore.MapAny{
					"id":         id,
					"uuid":       identity,
					"updated_at": timeUtcNow,
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
