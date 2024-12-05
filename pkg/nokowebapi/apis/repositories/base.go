package repositories

import (
	"errors"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"nokowebapi/apis/utils"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"strings"
)

type checkHandler[T any] func(schema *T) error

func (t checkHandler[T]) Call(schema *T) error {
	return t(schema)
}

type stmtConstraintFields struct {
	Query  string
	Values []any
}

type HookFunc func(tx *gorm.DB) (*gorm.DB, error)

type BaseRepositoryImpl[T any] interface {
	SafeFirst(query string, args ...any) (*T, error)
	SafeFirstHook(hook HookFunc) (*T, error)
	First(query string, args ...any) (*T, error)
	FirstHook(hook HookFunc) (*T, error)
	SafeMany(offset int, limit int, query string, args ...any) ([]T, error)
	SafeManyHook(hook HookFunc) ([]T, error)
	Many(offset int, limit int, query string, args ...any) ([]T, error)
	ManyHook(hook HookFunc) ([]T, error)
	SafePreFirst(preloads []string, query string, args ...any) (*T, error)
	PreFirst(preloads []string, query string, args ...any) (*T, error)
	SafePreMany(preloads []string, offset int, limit int, query string, args ...any) ([]T, error)
	PreMany(preloads []string, offset int, limit int, query string, args ...any) ([]T, error)
	TrySafeFind(model *T) (*T, error)
	TryFind(model *T) (*T, error)
	SafeTryManyFind(models []T) ([]T, error)
	TryManyFind(models []T) ([]T, error)
	SafeCheck(model *T, handler checkHandler[T]) error
	Check(model *T, handler checkHandler[T]) error
	Create(model *T) error
	SafeUpdate(model *T, query string, args ...any) error
	SafeUpdateHook(model *T, hook HookFunc) error
	Update(model *T, query string, args ...any) error
	UpdateHook(model *T, hook HookFunc) error
	SafeDelete(model *T, query string, args ...any) error
	SafeDeleteHook(model *T, hook HookFunc) error
	Delete(model *T, query string, args ...any) error
	DeleteHook(model *T, hook HookFunc) error
	Fields() ([]schema.Field, error)
}

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func NewBaseRepository[T any](DB *gorm.DB) BaseRepositoryImpl[T] {
	return &BaseRepository[T]{
		DB: DB,
	}
}

func (b *BaseRepository[T]) isRegis(schema *T) bool {
	schemaID := nokocore.GetValueWithSuperKey(schema, "BaseModel.id").(uint)
	schemaUUID := nokocore.GetValueWithSuperKey(schema, "BaseModel.uuid").(uuid.UUID)
	return schemaID != 0 && schemaUUID != uuid.Nil
}

func (b *BaseRepository[T]) SafeFirst(query string, args ...any) (*T, error) {
	return b.SafeFirstHook(func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(query, args...), nil
	})
}

func (b *BaseRepository[T]) SafeFirstHook(hook HookFunc) (*T, error) {
	var err error
	var model T

	stmt := b.DB.Where("deleted_at IS NULL")
	if stmt, err = hook(stmt); err != nil {
		return nil, err
	}

	tx := stmt.Limit(1).Find(&model)
	if err = tx.Error; err != nil {
		console.Error(fmt.Sprintf("panic: %s", err.Error()))
		return nil, errors.New("failed to find table")
	}

	if b.isRegis(&model) {
		return &model, nil
	}

	return nil, nil
}

func (b *BaseRepository[T]) First(query string, args ...any) (*T, error) {
	return b.FirstHook(func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(query, args...), nil
	})
}

func (b *BaseRepository[T]) FirstHook(hook HookFunc) (*T, error) {
	var err error
	var model T

	stmt := b.DB.Unscoped()
	if stmt, err = hook(stmt); err != nil {
		return nil, err
	}

	tx := stmt.Limit(1).Find(&model)
	if err = tx.Error; err != nil {
		console.Error(fmt.Sprintf("panic: %s", err.Error()))
		return nil, errors.New("failed to find table")
	}

	if b.isRegis(&model) {
		return &model, nil
	}

	return nil, nil
}

func (b *BaseRepository[T]) SafeMany(offset int, limit int, query string, args ...any) ([]T, error) {
	return b.SafeManyHook(func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(query, args...).Offset(offset).Limit(limit), nil
	})
}

func (b *BaseRepository[T]) SafeManyHook(hook HookFunc) ([]T, error) {
	var err error
	var models []T

	stmt := b.DB.Where("deleted_at IS NULL")
	if stmt, err = hook(stmt); err != nil {
		return nil, err
	}

	tx := stmt.Find(&models)
	if err = tx.Error; err != nil {
		console.Error(fmt.Sprintf("panic: %s", err.Error()))
		return nil, errors.New("failed to find table")
	}

	return models, nil
}

func (b *BaseRepository[T]) Many(offset int, limit int, query string, args ...any) ([]T, error) {
	return b.ManyHook(func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(query, args...).Offset(offset).Limit(limit), nil
	})
}

func (b *BaseRepository[T]) ManyHook(hook HookFunc) ([]T, error) {
	var err error
	var models []T

	stmt := b.DB.Unscoped()
	if stmt, err = hook(stmt); err != nil {
		return nil, err
	}

	tx := stmt.Find(&models)
	if err = tx.Error; err != nil {
		console.Error(fmt.Sprintf("panic: %s", err.Error()))
		return nil, errors.New("failed to find table")
	}

	return models, nil
}

func (b *BaseRepository[T]) SafePreFirst(preloads []string, query string, args ...any) (*T, error) {
	return b.SafeFirstHook(func(tx *gorm.DB) (*gorm.DB, error) {
		for i, preload := range preloads {
			nokocore.KeepVoid(i)
			tx = tx.Preload(preload)
		}

		return tx.Where(query, args...), nil
	})
}

func (b *BaseRepository[T]) PreFirst(preloads []string, query string, args ...any) (*T, error) {
	return b.FirstHook(func(tx *gorm.DB) (*gorm.DB, error) {
		for i, preload := range preloads {
			nokocore.KeepVoid(i)
			tx = tx.Preload(preload)
		}

		return tx.Where(query, args...), nil
	})
}

func (b *BaseRepository[T]) SafePreMany(preloads []string, offset int, limit int, query string, args ...any) ([]T, error) {
	return b.SafeManyHook(func(tx *gorm.DB) (*gorm.DB, error) {
		for i, preload := range preloads {
			nokocore.KeepVoid(i)
			tx = tx.Preload(preload)
		}

		return tx.Where(query, args...).Offset(offset).Limit(limit), nil
	})
}

func (b *BaseRepository[T]) PreMany(preloads []string, offset int, limit int, query string, args ...any) ([]T, error) {
	return b.ManyHook(func(tx *gorm.DB) (*gorm.DB, error) {
		for i, preload := range preloads {
			nokocore.KeepVoid(i)
			tx = tx.Preload(preload)
		}

		return tx.Where(query, args...).Offset(offset).Limit(limit), nil
	})
}

func (b *BaseRepository[T]) TrySafeFind(model *T) (*T, error) {
	var err error
	var check *T
	var stmt *stmtConstraintFields
	nokocore.KeepVoid(err, check, stmt)

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))

		if schemaID := nokocore.GetValueWithSuperKey(model, "BaseModel.id").(uint); schemaID != 0 {
			if check, err = b.SafeFirst("id = ?", schemaID); err != nil {
				return nil, errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
			}
		}

		if check == nil {
			if schemaUUID := nokocore.GetValueWithSuperKey(model, "BaseModel.uuid").(uuid.UUID); schemaUUID != uuid.Nil {
				if check, err = b.SafeFirst("uuid = ?", schemaUUID); err != nil {
					return nil, errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
				}
			}
		}

		if check == nil {
			if stmt, err = b.makeStmtWithConstraintFields(model); err != nil {
				return nil, fmt.Errorf("failed to find '%s' table, %w", tableNameType, err)
			}

			if stmt != nil {
				if check, err = b.SafeFirst(stmt.Query, stmt.Values...); err != nil {
					return nil, errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
				}
			}
		}

		// pass away
		return check, nil
	}

	return nil, errors.New("invalid value")
}

func (b *BaseRepository[T]) TryFind(model *T) (*T, error) {
	var err error
	var check *T
	var stmt *stmtConstraintFields
	nokocore.KeepVoid(err, check, stmt)

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))

		if schemaID := nokocore.GetValueWithSuperKey(model, "BaseModel.id").(uint); schemaID != 0 {
			if check, err = b.First("id = ?", schemaID); err != nil {
				return nil, errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
			}
		}

		if check == nil {
			if schemaUUID := nokocore.GetValueWithSuperKey(model, "BaseModel.uuid").(uuid.UUID); schemaUUID != uuid.Nil {
				if check, err = b.First("uuid = ?", schemaUUID); err != nil {
					return nil, errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
				}
			}
		}

		if check == nil {
			if stmt, err = b.makeStmtWithConstraintFields(model); err != nil {
				return nil, fmt.Errorf("failed to find '%s' table, %w", tableNameType, err)
			}

			if stmt != nil {
				if check, err = b.First(stmt.Query, stmt.Values...); err != nil {
					return nil, errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
				}
			}
		}

		// pass away
		return check, nil
	}

	return nil, errors.New("invalid value")
}

func (b *BaseRepository[T]) SafeTryManyFind(models []T) ([]T, error) {
	var err error
	var check *T
	var checks []T
	nokocore.KeepVoid(err, check, checks)

	if models != nil {
		for i, model := range models {
			nokocore.KeepVoid(i)
			if check, err = b.TrySafeFind(&model); err != nil {
				return nil, err
			}

			if check != nil {
				checks = append(checks, *check)
			}
		}

		return checks, nil
	}

	return nil, errors.New("invalid value")
}

func (b *BaseRepository[T]) TryManyFind(models []T) ([]T, error) {
	var err error
	var check *T
	var checks []T
	nokocore.KeepVoid(err, check, checks)

	if models != nil {
		for i, model := range models {
			nokocore.KeepVoid(i)
			if check, err = b.TryFind(&model); err != nil {
				return nil, err
			}

			if check != nil {
				checks = append(checks, *check)
			}
		}

		return checks, nil
	}

	return nil, errors.New("invalid value")
}

func (b *BaseRepository[T]) SafeCheck(model *T, handler checkHandler[T]) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))
		if check, err = b.TrySafeFind(model); err != nil {
			return err
		}

		if check != nil && handler != nil {
			return handler(check)
		}

		return errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) Check(model *T, handler checkHandler[T]) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))
		if check, err = b.TryFind(model); err != nil {
			return err
		}

		if check != nil && handler != nil {
			return handler(check)
		}

		return errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) baseInit(model *T) error {
	var err error
	nokocore.KeepVoid(err)

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))

		if b.isRegis(model) {
			return nil
		}

		timeUtcNow := nokocore.GetTimeUtcNow()
		err = mapstructure.Decode(nokocore.MapAny{
			"BaseModel": nokocore.MapAny{
				"uuid":       nokocore.NewUUID(),
				"created_at": timeUtcNow,
				"updated_at": timeUtcNow,
				"deleted_at": gorm.DeletedAt{},
			},
		}, model)

		if err != nil {
			return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
		}

		return nil
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) Create(model *T) error {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))

		if check, err = b.TryFind(model); err != nil {
			return err
		}

		if check != nil {

			// restore and replace schema existing
			timeUtcNow := nokocore.GetTimeUtcNow()
			err = mapstructure.Decode(nokocore.MapAny{
				"BaseModel": nokocore.MapAny{
					"created_at": timeUtcNow,
					"updated_at": timeUtcNow,
				},
			}, check)
			if err != nil {
				return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
			}

			schemaID := nokocore.GetValueWithSuperKey(check, "BaseModel.id").(uint)
			if err = b.Update(check, "id = ?", schemaID); err != nil {
				return err
			}

			return nil
		}

		if err = b.baseInit(model); err != nil {
			return err
		}

		tx := b.DB.Create(model)
		if err = tx.Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return errors.New(fmt.Sprintf("failed to create '%s' table", tableNameType))
		}

		if tx.RowsAffected < 1 {
			return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
		}

		return nil
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) SafeUpdate(model *T, query string, args ...any) error {
	return b.SafeUpdateHook(model, func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(query, args...), nil
	})
}

func (b *BaseRepository[T]) SafeUpdateHook(model *T, hook HookFunc) error {
	var err error
	var check *T

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))
		timeUtcNow := nokocore.GetTimeUtcNow()

		// using mapstructure to inject any values
		err = mapstructure.Decode(nokocore.MapAny{
			"BaseModel": nokocore.MapAny{
				"updated_at": timeUtcNow,
			},
		}, model)

		if err != nil {
			return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
		}

		if b.isRegis(model) {
			tx := b.DB.Save(model)
			if err = tx.Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New(fmt.Sprintf("failed to update '%s' table", tableNameType))
			}

			if tx.RowsAffected < 1 {
				return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
			}

			return nil
		}

		if check, err = b.SafeFirstHook(hook); err != nil {
			return err
		}

		if check != nil {
			// injecting check values into the current model
			err = mapstructure.Decode(nokocore.MapAny{
				"BaseModel": nokocore.MapAny{
					"id":   nokocore.GetValueWithSuperKey(check, "BaseModel.id").(uint),
					"uuid": nokocore.GetValueWithSuperKey(check, "BaseModel.uuid").(uuid.UUID),
				},
			}, model)

			if err != nil {
				return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
			}

			tx := b.DB.Save(model)
			if err = tx.Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New(fmt.Sprintf("failed to update '%s' table", tableNameType))
			}

			if tx.RowsAffected < 1 {
				return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
			}

			return nil
		}

		return errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) Update(model *T, query string, args ...any) error {
	return b.UpdateHook(model, func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where(query, args...), nil
	})
}

func (b *BaseRepository[T]) UpdateHook(model *T, hook HookFunc) error {
	var err error
	var check *T

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))
		timeUtcNow := nokocore.GetTimeUtcNow()

		// using mapstructure to inject any values
		err = mapstructure.Decode(nokocore.MapAny{
			"BaseModel": nokocore.MapAny{
				"updated_at": timeUtcNow,
			},
		}, model)

		if err != nil {
			return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
		}

		if b.isRegis(model) {
			tx := b.DB.Unscoped().Save(model)
			if err = tx.Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New(fmt.Sprintf("failed to update '%s' table", tableNameType))
			}

			if tx.RowsAffected < 1 {
				return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
			}

			return nil
		}

		if check, err = b.FirstHook(hook); err != nil {
			return err
		}

		if check != nil {
			// injecting check values into the current model
			err = mapstructure.Decode(nokocore.MapAny{
				"BaseModel": nokocore.MapAny{
					"id":   nokocore.GetValueWithSuperKey(check, "BaseModel.id").(uint),
					"uuid": nokocore.GetValueWithSuperKey(check, "BaseModel.uuid").(uuid.UUID),
				},
			}, model)

			if err != nil {
				return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
			}

			tx := b.DB.Unscoped().Save(model)
			if err = tx.Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New(fmt.Sprintf("failed to update '%s' table", tableNameType))
			}

			if tx.RowsAffected < 1 {
				return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
			}

			return nil
		}

		return errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) SafeDelete(model *T, query string, args ...any) error {
	var err error

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))
		timeUtcNow := nokocore.GetTimeUtcNow()

		// using mapstructure to inject any values
		err = mapstructure.Decode(nokocore.MapAny{
			"BaseModel": nokocore.MapAny{
				"deleted_at": sqlx.NewDeletedAt(timeUtcNow),
			},
		}, model)

		if err != nil {
			return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
		}

		if err = b.SafeUpdate(model, query, args...); err != nil {
			return errors.New(fmt.Sprintf("unable to delete '%s' table", tableNameType))
		}

		return nil
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) SafeDeleteHook(model *T, hook HookFunc) error {
	var err error

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))
		timeUtcNow := nokocore.GetTimeUtcNow()

		// using mapstructure to inject any values
		err = mapstructure.Decode(nokocore.MapAny{
			"BaseModel": nokocore.MapAny{
				"deleted_at": sqlx.NewDeletedAt(timeUtcNow),
			},
		}, model)

		if err != nil {
			return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
		}

		if err = b.SafeUpdateHook(model, hook); err != nil {
			return errors.New(fmt.Sprintf("unable to delete '%s' table", tableNameType))
		}

		return nil
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) Delete(model *T, query string, args ...any) error {
	var err error
	var check *T

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))

		if b.isRegis(model) {
			tx := b.DB.Unscoped().Delete(model)
			if err = tx.Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New(fmt.Sprintf("unable to delete '%s' table", tableNameType))
			}

			if tx.RowsAffected < 1 {
				return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
			}

			return nil
		}

		if check, err = b.First(query, args...); err != nil {
			return err
		}

		if check != nil {
			// injecting check values into the current model
			err = mapstructure.Decode(nokocore.MapAny{
				"BaseModel": nokocore.MapAny{
					"id":   nokocore.GetValueWithSuperKey(check, "BaseModel.id").(uint),
					"uuid": nokocore.GetValueWithSuperKey(check, "BaseModel.uuid").(uuid.UUID),
				},
			}, model)

			if err != nil {
				return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
			}

			tx := b.DB.Unscoped().Delete(model)
			if err = tx.Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New(fmt.Sprintf("unable to delete '%s' table", tableNameType))
			}

			if tx.RowsAffected < 1 {
				return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
			}

			return nil
		}

		return errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) DeleteHook(model *T, hook HookFunc) error {
	var err error
	var check *T

	if model != nil {
		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(model))

		if b.isRegis(model) {
			tx := b.DB.Unscoped().Delete(model)
			if err = tx.Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New(fmt.Sprintf("unable to delete '%s' table", tableNameType))
			}

			if tx.RowsAffected < 1 {
				return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
			}

			return nil
		}

		if check, err = b.FirstHook(hook); err != nil {
			return err
		}

		if check != nil {
			// injecting check values into the current model
			err = mapstructure.Decode(nokocore.MapAny{
				"BaseModel": nokocore.MapAny{
					"id":   nokocore.GetValueWithSuperKey(check, "BaseModel.id").(uint),
					"uuid": nokocore.GetValueWithSuperKey(check, "BaseModel.uuid").(uuid.UUID),
				},
			}, model)

			if err != nil {
				return errors.New(fmt.Sprintf("failed to inject values into '%s' table", tableNameType))
			}

			tx := b.DB.Unscoped().Delete(model)
			if err = tx.Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New(fmt.Sprintf("unable to delete '%s' table", tableNameType))
			}

			if tx.RowsAffected < 1 {
				return errors.New(fmt.Sprintf("no rows affected in '%s' table", tableNameType))
			}

			return nil
		}

		return errors.New(fmt.Sprintf("failed to find '%s' table", tableNameType))
	}

	return errors.New("invalid value")
}

func (b *BaseRepository[T]) makeStmtWithConstraintFields(model *T) (*stmtConstraintFields, error) {
	var err error
	var fields []schema.Field
	nokocore.KeepVoid(err, fields)

	if model != nil {
		if fields, err = b.Fields(); err != nil {
			return nil, err
		}

		size := 0
		var values []any
		var queries []string
		for i, field := range fields {
			nokocore.KeepVoid(i)

			dbName := field.DBName
			fieldName := nokocore.ToSnakeCase(field.Name)
			isConstraint := field.PrimaryKey || field.Unique
			if isConstraint {
				var value any
				if dbName != "id" && dbName != "uuid" {
					value = nokocore.GetValueWithSuperKey(model, fieldName)
					if utils.SqlValueIsNull(value) {
						if field.NotNull {
							return nil, errors.New(fmt.Sprintf("field '%s' is required", fieldName))
						}

						continue
					}
				} else {
					fieldName = fmt.Sprintf("BaseModel.%s", fieldName)
					value = nokocore.GetValueWithSuperKey(model, fieldName)
					if utils.SqlValueIsNull(value) {
						continue
					}
				}

				queries = append(queries, fmt.Sprintf("%s = ?", dbName))
				values = append(values, value)
				size += 1
			}
		}

		if size > 0 {
			query := strings.Join(queries, " OR ")
			stmt := &stmtConstraintFields{
				Query:  query,
				Values: values,
			}

			return stmt, nil
		}

		// pass away
		return nil, nil
	}

	return nil, errors.New("invalid value")
}

func (b *BaseRepository[T]) Fields() ([]schema.Field, error) {
	var err error
	var model T
	nokocore.KeepVoid(err, model)

	stmt := &gorm.Statement{
		DB: b.DB,
	}

	if err = stmt.Parse(&model); err != nil {
		return nil, err
	}

	// make unreferenced copy of fields
	size := len(stmt.Schema.Fields)
	temp := make([]schema.Field, size)
	for i, field := range stmt.Schema.Fields {
		temp[i] = *field
		//field.PrimaryKey
		//field.Unique
		//field.HasDefaultValue
		//field.AutoIncrement
		//field.Creatable
		//field.NotNull
		//field.Readable
		//field.Updatable
	}

	return temp, nil
}
