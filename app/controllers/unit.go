package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"pharma-cash-go/app/repositories"
)

func UnitController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/units", repositories.CreateUnit(DB))

	return group
}