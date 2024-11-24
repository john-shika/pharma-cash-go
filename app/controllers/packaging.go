package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"pharma-cash-go/app/repositories"
)

func PackagingController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/packagings", repositories.CreatePackaging(DB))

	return group
}