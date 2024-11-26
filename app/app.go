package app

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis"
	"nokowebapi/apis/middlewares"
	controllers2 "pharma-cash-go/app/controllers"
	factories2 "pharma-cash-go/app/factories"
	models2 "pharma-cash-go/app/models"
)

func Controllers(group *echo.Group, DB *gorm.DB) []*echo.Group {
	auth := group.Group("/auth")
	auth.Use(middlewares.JWTAuth(DB))

	return []*echo.Group{
		controllers2.GuestController(group, DB),
		controllers2.UserController(auth, DB),
		controllers2.AdminController(auth, DB),
		controllers2.ProductController(auth, DB),
		controllers2.UnitController(auth, DB),
		controllers2.PackagingController(auth, DB),
	}
}

func Factories(DB *gorm.DB) []any {
	return []any{
		factories2.UserFactory(DB),
		factories2.ShiftFactory(DB),
	}
}

func DBAutoMigrations(DB *gorm.DB) {
	tables := []any{
		&models2.Product{},
		&models2.Employee{},
		&models2.Shift{},
		&models2.Category{},
		&models2.Product{},
		&models2.ProductCategory{},
		&models2.Package{},
		&models2.Unit{},
	}

	apis.DBAutoMigrations(DB, tables)
}
