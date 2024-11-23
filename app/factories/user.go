package factories

import (
	"fmt"
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"pharma-cash-go/app/repositories"
)

func UserFactory(DB *gorm.DB) []models.User {
	var err error
	nokocore.KeepVoid(err)

	users := []models.User{
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
			Phone:    sqlx.NewString("+62 823-4567-8901"),
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

	return users
}
