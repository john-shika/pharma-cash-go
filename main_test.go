package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/apis/repositories"
	"nokowebapi/console"
	"nokowebapi/console/zapgorm"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestDB(t *testing.T) {
	var err error
	var DB *gorm.DB
	var user *models.User
	var users []models.User
	nokocore.KeepVoid(err, DB, user, users)

	config := &gorm.Config{
		Logger: zapgorm.New(console.GetLogger("GORM")),
	}

	sqliteFilePath := "migrations/dev.sqlite3"
	nokocore.NoErr(nokocore.CreateEmptyFile(sqliteFilePath))
	if DB, err = gorm.Open(sqlite.Open(sqliteFilePath), config); err != nil {
		panic("failed to connect database")
	}

	tables := []interface{}{
		&models.User{},
		&models.Session{},
	}

	if err = DB.AutoMigrate(tables...); err != nil {
		console.Fatal(fmt.Sprintf("failed to migrate database: %s\n", err.Error()))
	}

	/// dummy data

	users = []models.User{
		{
			Username: "admin",
			Password: "Admin@1234",
			FullName: sqlx.NewString("John, Doe"),
			Email:    sqlx.NewString("admin@example.com"),
			Phone:    sqlx.NewString("+62 812-3456-7890"),
			Admin:    true,
			Level:    1,
		},
		{
			Username: "user",
			Password: "User@1234",
			FullName: sqlx.NewString("Angeline, Rose"),
			Email:    sqlx.NewString("user@example.com"),
			Phone:    sqlx.NewString("+62 823-4567-8901"), // conflict, same phone number
			Admin:    false,
			Level:    1,
		},
	}

	userRepository := repositories.NewUserRepository(DB)

	var check *models.User
	for i, user := range users {
		nokocore.KeepVoid(i)

		if check, err = userRepository.First("username = ?", user.Username); err != nil {
			console.Warn(err.Error())
			continue
		}

		if check != nil {
			console.Warn(fmt.Sprintf("user '%s' already exists", user.Username))
			continue
		}

		if err = userRepository.Create(&user); err != nil {
			console.Warn(err.Error())
			continue
		}

		console.Warn(fmt.Sprintf("user '%s' has been created", user.Username))
	}

	/// unit tests

	// find all users
	if users, err = userRepository.SafeMany(0, 2, "1=1"); err != nil {
		t.Error(err)
		return
	}

	if len(users) != 2 {
		t.Error(errors.New("users count should be 2"))
		return
	}

	// find all users with first order
	if user, err = userRepository.SafeFirst("1=1"); err != nil {
		t.Error(err)
		return
	}

	if user == nil {
		t.Error(errors.New("user is null"))
		return
	}

	// find admin user
	if user, err = userRepository.First("username = ?", "admin"); err != nil {
		t.Error(err)
		return
	}

	if user == nil {
		t.Error(errors.New("user is null"))
		return
	}

	// safe delete admin user
	if err = userRepository.SafeDelete(user, "username = ?", "admin"); err != nil {
		t.Error(errors.New("user can't be deleted"))
		return
	}

	// check admin user
	if user, err = userRepository.SafeFirst("username = ?", "admin"); user != nil {
		t.Error(errors.New("user should be soft deleted"))
		return
	}

	// check admin user
	if user, err = userRepository.First("username = ?", "admin"); err != nil {
		t.Error(errors.New("user not found"))
		return
	}

	// restore admin user
	user.DeletedAt = gorm.DeletedAt{}
	if err = userRepository.Update(user, "username = ?", "admin"); err != nil {
		t.Error(err)
		return
	}

	// find admin user
	if user, err = userRepository.SafeFirst("username = ?", "admin"); err != nil {
		t.Error(errors.New("user not found"))
		return
	}

	// delete admin user
	if err = userRepository.Delete(user, "username = ?", "admin"); err != nil {
		t.Error(errors.New("user can't be deleted"))
		return
	}

	// find admin user
	if user, err = userRepository.First("username = ?", "admin"); user != nil {
		t.Error(errors.New("user should be deleted"))
		return
	}
}
