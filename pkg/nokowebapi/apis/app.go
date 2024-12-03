package apis

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"nokowebapi/apis/models"
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

type FactoryData map[int][]any

func (FactoryData) GetNameType() string {
	return "FactoryData"
}

type FactoryHook func(tx *gorm.DB) []any

func (FactoryHook) GetNameType() string {
	return "FactoryHook"
}

func Factories(DB *gorm.DB, factories []FactoryHook) FactoryData {
	temp := make(FactoryData)
	for i, factory := range factories {
		nokocore.KeepVoid(i)
		temp[i] = factory(DB)
	}

	return temp
}

func Migrations(DB *gorm.DB, tables []any) error {
	var err error
	nokocore.KeepVoid(err)

	defaults := []any{
		&models.User{},
		&models.Role{},
		&models.UserRoles{},
		&models.Session{},
	}

	if err = DB.AutoMigrate(defaults...); err != nil {
		return err
	}

	if err = DB.AutoMigrate(tables...); err != nil {
		return err
	}

	return err
}
