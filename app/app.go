package app

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/middlewares"
	"pharma-cash-go/app/controllers"
	"pharma-cash-go/app/factories"
	"pharma-cash-go/app/models"
)

func Controllers(group *echo.Group, DB *gorm.DB) []*echo.Group {
	auth := group.Group("/auth")
	auth.Use(middlewares.JWTAuth(DB))

	return []*echo.Group{
		controllers.GuestController(group, DB),
		controllers.UserController(auth, DB),
		controllers.AdminController(auth, DB),
	}
}

func Factories(DB *gorm.DB) []any {
	return []any{
		factories.UserFactory(DB),
		factories.ShiftFactory(DB),
	}
}

func Tables() []any {
	return []any{
		&models.Product{},
		&models.Employee{},
		&models.Shift{},
	}
}
