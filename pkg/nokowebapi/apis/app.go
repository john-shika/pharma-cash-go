package apis

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/console"
	"nokowebapi/nokocore"
)

func SqliteOpenFile(path string, config *gorm.Config) (*gorm.DB, error) {
	var err error
	var DB *gorm.DB
	nokocore.KeepVoid(err, DB)

	if err = nokocore.CreateEmptyFile(path); err != nil {
		return nil, err
	}

	if DB, err = gorm.Open(sqlite.Open(path), config); err != nil {
		return nil, fmt.Errorf("failed to open database, %w", err)
	}

	return DB, nil
}

func DBAutoMigrations(DB *gorm.DB, tables []any) {
	var err error
	nokocore.KeepVoid(err)

	defaults := []any{
		&models.User{},
		&models.Role{},
		&models.UserRoles{},
		&models.Session{},
	}

	if err = DB.AutoMigrate(defaults...); err != nil {
		console.Fatal(fmt.Sprintf("failed to migrate database: %s\n", err.Error()))
	}

	if err = DB.AutoMigrate(tables...); err != nil {
		console.Fatal(fmt.Sprintf("failed to migrate database: %s\n", err.Error()))
	}
}
