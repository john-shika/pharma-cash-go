package repositories

import (
	"errors"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"time"
)

type BaseRepositoryImpl[T any] interface {
	SafeFirst(query string, args ...any) (*T, error)
	First(query string, args ...any) (*T, error)
	SafeMany(offset int, limit int, query string, args ...any) ([]T, error)
	Many(offset int, limit int, query string, args ...any) ([]T, error)
	SafePreFirst(preloads []string, query string, args ...any) (*T, error)
	PreFirst(preloads []string, query string, args ...any) (*T, error)
	SafePreMany(preloads []string, offset int, limit int, query string, args ...any) ([]T, error)
	PreMany(preloads []string, offset int, limit int, query string, args ...any) ([]T, error)
	SafeCheck(schema *T, checkHandler CheckHandler[T]) error
	Check(schema *T, checkHandler CheckHandler[T]) error
	SafeCreate(schema *T) error
	Create(schema *T) error
	SafeUpdate(schema *T, query string, args ...any) error
	Update(schema *T, query string, args ...any) error
	SafeDelete(schema *T, query string, args ...any) error
	Delete(schema *T, query string, args ...any) error
}

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func NewBaseRepository[T any](DB *gorm.DB) BaseRepository[T] {
	return BaseRepository[T]{
		DB: DB,
	}
}

func (b *BaseRepository[T]) SafeFirst(query string, args ...any) (*T, error) {
	var err error
	var schema T
	nokocore.KeepVoid(err, schema)

	stmt := b.DB.Where("deleted_at IS NULL")
	tx := stmt.Where(query, args...).Limit(1).Find(&schema)
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

func (b *BaseRepository[T]) First(query string, args ...any) (*T, error) {
	var err error
	var schema T
	nokocore.KeepVoid(err, schema)

	stmt := b.DB.Unscoped()
	tx := stmt.Where(query, args...).Limit(1).Find(&schema)
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

func (b *BaseRepository[T]) SafeMany(offset int, limit int, query string, args ...any) ([]T, error) {
	var err error
	var schemas []T
	nokocore.KeepVoid(err, schemas)

	stmt := b.DB.Where("deleted_at IS NULL")
	tx := stmt.Where(query, args...).Offset(offset).Limit(limit).Find(&schemas)
	if err = tx.Error; err != nil {
		return nil, err
	}

	return schemas, nil
}

func (b *BaseRepository[T]) Many(offset int, limit int, query string, args ...any) ([]T, error) {
	var err error
	var schemas []T
	nokocore.KeepVoid(err, schemas)

	stmt := b.DB.Unscoped()
	tx := stmt.Where(query, args...).Offset(offset).Limit(limit).Find(&schemas)
	if err = tx.Error; err != nil {
		return nil, err
	}

	return schemas, nil
}

func (b *BaseRepository[T]) SafePreFirst(preloads []string, query string, args ...any) (*T, error) {
	var err error
	var schema T
	nokocore.KeepVoid(err, schema)

	stmt := b.DB.Where("deleted_at IS NULL")
	for i, preload := range preloads {
		nokocore.KeepVoid(i)

		stmt = stmt.Preload(preload)
	}

	tx := stmt.Where(query, args...).Limit(1).Find(&schema)
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

func (b *BaseRepository[T]) PreFirst(preloads []string, query string, args ...any) (*T, error) {
	var err error
	var schema T
	nokocore.KeepVoid(err, schema)

	stmt := b.DB.Unscoped()
	for i, preload := range preloads {
		nokocore.KeepVoid(i)

		stmt = stmt.Preload(preload)
	}

	tx := stmt.Where(query, args...).Limit(1).Find(&schema)
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

func (b *BaseRepository[T]) SafePreMany(preloads []string, offset int, limit int, query string, args ...any) ([]T, error) {
	var err error
	var schemas []T
	nokocore.KeepVoid(err, schemas)

	stmt := b.DB.Where("deleted_at IS NULL")
	for i, preload := range preloads {
		nokocore.KeepVoid(i)

		stmt = stmt.Preload(preload)
	}

	tx := stmt.Where(query, args...).Offset(offset).Limit(limit).Find(&schemas)
	if err = tx.Error; err != nil {
		return nil, err
	}

	return schemas, nil
}

func (b *BaseRepository[T]) PreMany(preloads []string, offset int, limit int, query string, args ...any) ([]T, error) {
	var err error
	var schemas []T
	nokocore.KeepVoid(err, schemas)

	stmt := b.DB.Unscoped()
	for i, preload := range preloads {
		nokocore.KeepVoid(i)

		stmt = stmt.Preload(preload)
	}

	tx := stmt.Where(query, args...).Offset(offset).Limit(limit).Find(&schemas)
	if err = tx.Error; err != nil {
		return nil, err
	}

	return schemas, nil
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
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(schema))
		id := nokocore.GetValueWithSuperKey(schema, "BaseModel.id").(int)
		if id != 0 {
			if check, err = b.SafeFirst("id = ?", id); err != nil {
				return errors.New(fmt.Sprintf("failed to search '%s' table", tableNameType))
			}
		}

		identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
		if identity != uuid.Nil {
			if check, err = b.SafeFirst("uuid = ?", identity); err != nil {
				return errors.New(fmt.Sprintf("failed to search '%s' table", tableNameType))
			}
		}

		if check != nil {
			if checkHandler != nil {
				return checkHandler.Call(check)
			}

			return errors.New(fmt.Sprintf("'%s' table already exists", tableNameType))
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
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(schema))
		id := nokocore.GetValueWithSuperKey(schema, "BaseModel.id").(int)
		if id != 0 {
			if check, err = b.First("id = ?", id); err != nil {
				return errors.New(fmt.Sprintf("failed to search '%s' table", tableNameType))
			}
		}

		identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
		if identity != uuid.Nil {
			if check, err = b.First("uuid = ?", identity); err != nil {
				return errors.New(fmt.Sprintf("failed to search '%s' table", tableNameType))
			}
		}

		if check != nil {
			if checkHandler != nil {
				return checkHandler.Call(check)
			}

			return errors.New(fmt.Sprintf("'%s' table already exists", tableNameType))
		}

		return nil
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) baseInit(schema *T) error {
	var err error
	nokocore.KeepVoid(err)

	tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(schema))

	// using mapstructure to inject any values
	identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
	if identity != uuid.Nil {
		return nil

	}

	identity = nokocore.NewUUID()
	deletedAt := gorm.DeletedAt{}

	timeUtcNow := nokocore.GetTimeUtcNow()
	err = mapstructure.Decode(nokocore.MapAny{
		"BaseModel": nokocore.MapAny{
			"uuid":       identity,
			"created_at": timeUtcNow,
			"updated_at": timeUtcNow,
			"deleted_at": deletedAt,
		},
	}, schema)

	if err != nil {
		return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
	}

	return nil
}

func (b *BaseRepository[T]) SafeCreate(schema *T) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	if schema != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(schema))
		if err = b.SafeCheck(schema, nil); err != nil {
			return err
		}

		if err = b.baseInit(schema); err != nil {
			return err
		}

		tx := b.DB.Create(schema)
		if err = tx.Error; err != nil {
			return errors.New(fmt.Sprintf("failed to create '%s' table", tableNameType))
		}

		if tx.RowsAffected > 0 {
			return nil
		}

		return fmt.Errorf("no rows affected in '%s' table", tableNameType)
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) Create(schema *T) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	if schema != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(schema))
		if err = b.Check(schema, nil); err != nil {
			return err
		}

		if err = b.baseInit(schema); err != nil {
			return err
		}

		tx := b.DB.Unscoped().Create(schema)
		if err = tx.Error; err != nil {
			return errors.New(fmt.Sprintf("failed to create '%s' table", tableNameType))
		}

		if tx.RowsAffected > 0 {
			return nil
		}

		return fmt.Errorf("no rows affected in '%s' table", tableNameType)
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) SafeUpdate(schema *T, query string, args ...any) error {
	var err error
	var id int
	var identity uuid.UUID
	var check *T
	nokocore.KeepVoid(err, id, identity, check)

	if schema != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(schema))

		id = nokocore.GetValueWithSuperKey(schema, "BaseModel.id").(int)
		identity = nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)

		if id != 0 && identity != uuid.Nil {
			tx := b.DB.Save(schema)
			if err = tx.Error; err != nil {
				return errors.New(fmt.Sprintf("failed to update '%s' table", tableNameType))
			}

			if tx.RowsAffected > 0 {
				return nil
			}

			return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
		}

		if check, err = b.SafeFirst(query, args...); err != nil {
			return err
		}

		if check != nil {
			id = nokocore.GetValueWithSuperKey(check, "BaseModel.id").(int)
			identity = nokocore.GetValueWithSuperKey(check, "BaseModel.uuid").(uuid.UUID)

			// keep it
			createdAt := nokocore.GetValueWithSuperKey(check, "BaseModel.created_at").(time.Time)
			deletedAt := gorm.DeletedAt{}

			// using mapstructure to inject any values
			timeUtcNow := nokocore.GetTimeUtcNow()
			err = mapstructure.Decode(nokocore.MapAny{
				"BaseModel": nokocore.MapAny{
					"id":         id,
					"uuid":       identity,
					"created_at": createdAt,
					"updated_at": timeUtcNow,
					"deleted_at": deletedAt,
				},
			}, schema)

			if err != nil {
				return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
			}

			tx := b.DB.Save(schema)
			if err = tx.Error; err != nil {
				return errors.New(fmt.Sprintf("failed to update '%s' table", tableNameType))
			}

			if tx.RowsAffected > 0 {
				return nil
			}

			return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
		}

		return errors.New(fmt.Sprintf("failed to search '%s' table", tableNameType))
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) Update(schema *T, query string, args ...any) error {
	var err error
	var identity uuid.UUID
	var id int
	var check *T
	nokocore.KeepVoid(err, id, identity, check)

	if schema != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(schema))

		identity = nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)

		if identity != uuid.Nil {
			tx := b.DB.Unscoped().Save(schema)
			if err = tx.Error; err != nil {
				return errors.New(fmt.Sprintf("failed to update '%s' table", tableNameType))
			}

			if tx.RowsAffected > 0 {
				return nil
			}

			return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
		}

		if check, err = b.First(query, args...); err != nil {
			return err
		}

		if check != nil {
			id = nokocore.GetValueWithSuperKey(check, "BaseModel.id").(int)
			identity = nokocore.GetValueWithSuperKey(check, "BaseModel.uuid").(uuid.UUID)

			// keep it
			createdAt := nokocore.GetValueWithSuperKey(check, "BaseModel.created_at").(time.Time)
			deletedAt := nokocore.GetValueWithSuperKey(check, "BaseModel.deleted_at").(gorm.DeletedAt)

			// using mapstructure to inject any values
			timeUtcNow := nokocore.GetTimeUtcNow()
			err = mapstructure.Decode(nokocore.MapAny{
				"BaseModel": nokocore.MapAny{
					"id":         id,
					"uuid":       identity,
					"created_at": createdAt,
					"updated_at": timeUtcNow,
					"deleted_at": deletedAt,
				},
			}, schema)

			if err != nil {
				return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
			}

			tx := b.DB.Unscoped().Save(schema)
			if err = tx.Error; err != nil {
				return errors.New(fmt.Sprintf("failed to update '%s' table", tableNameType))
			}

			if tx.RowsAffected > 0 {
				return nil
			}

			return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
		}

		return errors.New(fmt.Sprintf("failed to search '%s' table", tableNameType))
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) SafeDelete(schema *T, query string, args ...any) error {
	var err error
	nokocore.KeepVoid(err)

	if schema != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(schema))
		timeUtcNow := nokocore.GetTimeUtcNow()

		// using mapstructure to inject any values
		err = mapstructure.Decode(nokocore.MapAny{
			"BaseModel": nokocore.MapAny{
				"deleted_at": sqlx.NewDeletedAt(timeUtcNow),
			},
		}, schema)

		if err != nil {
			return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
		}

		identity := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
		if identity != uuid.Nil {
			if err = b.SafeUpdate(schema, "uuid = ?", identity); err != nil {
				return errors.New(fmt.Sprintf("unable to delete '%s' table", tableNameType))
			}

			return nil
		}

		if err = b.SafeUpdate(schema, query, args...); err != nil {
			return errors.New(fmt.Sprintf("unable to delete '%s' table", tableNameType))
		}

		return nil
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) Delete(schema *T, query string, args ...any) error {
	var err error
	nokocore.KeepVoid(err)

	if schema != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(schema))
		tx := b.DB.Unscoped().Where(query, args...).Delete(schema)
		if err = tx.Error; err != nil {
			return errors.New(fmt.Sprintf("unable to delete '%s' table", tableNameType))
		}

		if tx.RowsAffected > 0 {
			return nil
		}

		return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
	}

	return errors.New("invalid value")
}