package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/utils"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
	repositories2 "pharma-cash-go/app/repositories"
	schemas2 "pharma-cash-go/app/schemas"
)

func CreateUnit(DB *gorm.DB) echo.HandlerFunc {

	unitRepository := repositories2.NewUnitRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var unit *models2.Unit
		nokocore.KeepVoid(err, unit)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		unitBody := new(schemas2.UnitBody)

		if err = ctx.Bind(unitBody); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		if err = ctx.Validate(unitBody); err != nil {
			return err
		}

		// normalized text
		unitBody.UnitType = nokocore.ToTitleCase(unitBody.UnitType)

		unitType := unitBody.UnitType
		if unit, err = unitRepository.SafeFirst("unit_type = ?", unitType); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get unit.", nil)
		}

		if unit != nil {
			unitResult := schemas2.ToUnitResult(unit)
			return extras.NewMessageBodyOk(ctx, "Unit already exists.", &nokocore.MapAny{
				"unit": unitResult,
			})
		}

		unit = schemas2.ToUnitModel(unitBody)
		if err = unitRepository.SafeCreate(unit); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create unit.", nil)
		}

		unitResult := schemas2.ToUnitResult(unit)
		return extras.NewMessageBodyOk(ctx, "Successfully create unit.", &nokocore.MapAny{
			"unit": unitResult,
		})
	}
}

func GetAllUnits(DB *gorm.DB) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		var err error
		nokocore.KeepVoid(err)

		pagination := extras.NewURLQueryPaginationFromEchoContext(ctx)

		var units []models2.Unit
		tx := DB.Offset(pagination.Offset).Limit(pagination.Limit).Find(&units)
		if err = tx.Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get units.", nil)
		}

		var unitResults []schemas2.UnitResult
		for i, unit := range units {
			nokocore.KeepVoid(i)
			unitResults = append(unitResults, schemas2.ToUnitResult(&unit))
		}

		return extras.NewMessageBodyOk(ctx, "Successfully get units.", &nokocore.MapAny{
			"units": unitResults,
		})
	}
}

func UnitController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.GET("/units", GetAllUnits(DB))
	group.POST("/unit", CreateUnit(DB))

	return group
}
