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

func CreatePackage(DB *gorm.DB) echo.HandlerFunc {
	var err error
	nokocore.KeepVoid(err)

	packageRepository := repositories2.NewPackageRepository(DB)

	return func(ctx echo.Context) error {
		packageBody := new(schemas2.PackageBody)

		if err = ctx.Bind(packageBody); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		if err = ctx.Validate(packageBody); err != nil {
			return err
		}

		packageModel := schemas2.ToPackageModel(packageBody)
		if err = packageRepository.SafeCreate(packageModel); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create package.", nil)
		}

		packageResult := schemas2.ToPackageResult(packageModel)
		return extras.NewMessageBodyOk(ctx, "Successfully create packageBody.", &nokocore.MapAny{
			"package": packageResult,
		})
	}
}

func PackagingController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/packages", CreatePackage(DB))

	return group
}
