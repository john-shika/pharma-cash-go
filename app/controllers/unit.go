package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	repositories2 "pharma-cash-go/app/repositories"
	schemas2 "pharma-cash-go/app/schemas"
)

func CreateUnit(DB *gorm.DB) echo.HandlerFunc {
	var err error
	nokocore.KeepVoid(err)

	unitRepository := repositories2.NewUnitRepository(DB)

	return func(ctx echo.Context) error {
		unitBody := new(schemas2.UnitBody)

		if err = ctx.Bind(unitBody); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		if err = ctx.Validate(unitBody); err != nil {
			return err
		}

		unitModel := schemas2.ToUnitModel(unitBody)
		if err = unitRepository.SafeCreate(unitModel); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create unit.", nil)
		}

		unitResult := schemas2.ToUnitResult(unitModel)
		return extras.NewMessageBodyOk(ctx, "Successfully create unit.", &nokocore.MapAny{
			"unit": unitResult,
		})
	}
}

func UnitController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/units", CreateUnit(DB))

	return group
}
