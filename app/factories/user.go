package factories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/factories"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
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

	return factories.BaseFactory[models.User](DB, users, "username = ?", func(user models.User) []any {
		return []any{
			user.Username,
		}
	})
}
