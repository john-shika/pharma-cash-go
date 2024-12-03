package factories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/factories"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
)

func UserFactory(DB *gorm.DB) []any {
	users := []any{
		models.User{
			Username: "admin",
			Password: "Admin@1234",
			FullName: sqlx.NewString("John, Doe"),
			Email:    sqlx.NewString("admin@example.com"),
			Phone:    sqlx.NewString("0 000-0000-0000-0000"),
			Admin:    true,
			Level:    1,
		},
		models.User{
			Username: "user",
			Password: "User@1234",
			FullName: sqlx.NewString("Angeline, Rose"),
			Email:    sqlx.NewString("user@example.com"),
			Phone:    sqlx.NewString("0 000-0000-0000-0001"),
			Admin:    false,
			Level:    0,
		},
	}

	return factories.BaseFactory(DB, users, "username = ?", func(user any) []any {
		return []any{
			nokocore.GetValueWithSuperKey(user, "username"),
		}
	})
}
