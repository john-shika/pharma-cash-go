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

func Controllers(group *echo.Group, DB *gorm.DB) {
	auth := group.Group("/auth")
	auth.Use(middlewares.JWTAuth(DB))

	controllers2.GuestController(group, DB)
	controllers2.UserController(auth, DB)
	controllers2.AdminController(auth, DB)
	controllers2.ProductController(auth, DB)
	controllers2.UnitController(auth, DB)
	controllers2.PackagingController(auth, DB)
	controllers2.ShopController(auth, DB)
}

func Factories(DB *gorm.DB) apis.FactoryData {
	return apis.Factories(DB, []apis.FactoryHook{
		factories2.UserFactory,
		factories2.ShiftFactory,
	})
}

func Migrations(DB *gorm.DB) error {
	return apis.Migrations(DB, []any{
		&models2.Barcode{},
		&models2.Cart{},
		&models2.Category{},
		&models2.Employee{},
		&models2.Package{},
		&models2.Product{},
		&models2.ProductCategory{},
		&models2.Shift{},
		&models2.Transaction{},
		&models2.Unit{},
	})
}
