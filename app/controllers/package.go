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

func CreatePackage(DB *gorm.DB) echo.HandlerFunc {

	packageRepository := repositories2.NewPackageRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var packageModel *models2.Package
		nokocore.KeepVoid(err)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		packageBody := new(schemas2.PackageBody)

		if err = ctx.Bind(packageBody); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		if err = ctx.Validate(packageBody); err != nil {
			return err
		}

		// normalized text
		packageBody.PackageType = nokocore.ToTitleCase(packageBody.PackageType)

		packageType := packageBody.PackageType
		if packageModel, err = packageRepository.SafeFirst("package_type = ?", packageType); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get package.", nil)
		}

		if packageModel != nil {
			packageResult := schemas2.ToPackageResult(packageModel)
			return extras.NewMessageBodyOk(ctx, "Package already exists.", &nokocore.MapAny{
				"package": packageResult,
			})
		}

		packageModel = schemas2.ToPackageModel(packageBody)
		if err = packageRepository.Create(packageModel); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create package.", nil)
		}

		packageResult := schemas2.ToPackageResult(packageModel)
		return extras.NewMessageBodyOk(ctx, "Successfully create packageBody.", &nokocore.MapAny{
			"package": packageResult,
		})
	}
}

func GetAllPackages(DB *gorm.DB) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		var err error
		nokocore.KeepVoid(err)

		pagination := extras.NewURLQueryPaginationFromEchoContext(ctx)

		var packages []models2.Package
		tx := DB.Offset(pagination.Offset).Limit(pagination.Limit).Find(&packages)
		if err = tx.Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get packages.", nil)
		}

		var packageResults []schemas2.PackageResult
		for i, packageModel := range packages {
			nokocore.KeepVoid(i)
			packageResults = append(packageResults, schemas2.ToPackageResult(&packageModel))
		}

		return extras.NewMessageBodyOk(ctx, "Successfully get packages.", &nokocore.MapAny{
			"packages": packageResults,
		})
	}
}

func PackagingController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/package", CreatePackage(DB))
	group.GET("/packages", GetAllPackages(DB))

	return group
}
