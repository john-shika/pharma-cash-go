package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"pharma-cash-go/app/repositories"
)

func ProductController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/products", repositories.CreateProduct(DB))

	return group
}