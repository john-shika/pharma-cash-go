package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
	repositories2 "pharma-cash-go/app/repositories"
	schemas2 "pharma-cash-go/app/schemas"
)

var validate = validator.New()

func CreateProduct(DB *gorm.DB) echo.HandlerFunc {

	packageRepository := repositories2.NewPackageRepository(DB)
	unitRepository := repositories2.NewUnitRepository(DB)
	productRepository := repositories2.NewProductRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var packageModel *models2.Package
		var unit *models2.Unit
		nokocore.KeepVoid(err, packageModel, unit)

		productBody := new(schemas2.ProductBody)
		if err = ctx.Bind(&productBody); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		if err = ctx.Validate(productBody); err != nil {
			return err
		}

		if packageModel, err = packageRepository.SafeFirst("uuid = ?", productBody.PackageID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get package.", nil)
		}

		// can be automatic build
		if packageModel == nil {
			return extras.NewMessageBodyNotFound(ctx, "Package not found.", nil)
		}

		if unit, err = unitRepository.SafeFirst("uuid = ?", productBody.UnitID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get unit.", nil)
		}

		// can be automatic build
		if unit == nil {
			return extras.NewMessageBodyNotFound(ctx, "Unit not found.", nil)
		}

		product := schemas2.ToProductModel(productBody, packageModel, unit)

		if err = productRepository.SafeCreate(product); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create product.", nil)
		}

		productResult := schemas2.ToProductResult(product)
		return extras.NewMessageBodyOk(ctx, "Successfully create product.", &nokocore.MapAny{
			"product": productResult,
		})
	}
}

func ProductController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/product", CreateProduct(DB))

	return group
}
