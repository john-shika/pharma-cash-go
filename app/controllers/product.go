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
	"strconv"
)

func CreateProduct(DB *gorm.DB) echo.HandlerFunc {

	packageRepository := repositories2.NewPackageRepository(DB)
	unitRepository := repositories2.NewUnitRepository(DB)
	productRepository := repositories2.NewProductRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var packageModel *models2.Package
		var unit *models2.Unit
		nokocore.KeepVoid(err, packageModel, unit)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		productBody := new(schemas2.ProductBody)
		if err = ctx.Bind(&productBody); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		if err = ctx.Validate(productBody); err != nil {
			return err
		}

		product := schemas2.ToProductModel(productBody)

		if packageID := productBody.PackageID; packageID != "" {
			if packageModel, err = packageRepository.SafeFirst("uuid = ?", packageID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyInternalServerError(ctx, "Failed to get package.", nil)
			}
		}

		// can be automatic build
		if packageModel == nil {
			if packageType := nokocore.ToTitleCase(productBody.PackageType); packageType != "" {
				if packageModel, err = packageRepository.SafeFirst("package_type = ?", packageType); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyInternalServerError(ctx, "Failed to get package.", nil)
				}

				if packageModel == nil {
					packageModel = &models2.Package{
						PackageType: packageType,
					}
					if err = packageRepository.SafeCreate(packageModel); err != nil {
						console.Error(fmt.Sprintf("panic: %s", err.Error()))
						return extras.NewMessageBodyInternalServerError(ctx, "Failed create package.", nil)
					}
				}

			} else {
				return extras.NewMessageBodyNotFound(ctx, "Package not found.", nil)
			}
		}

		product.PackageID = packageModel.ID
		product.Package = *packageModel

		if unitID := productBody.UnitID; unitID != "" {
			if unit, err = unitRepository.SafeFirst("uuid = ?", unitID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyInternalServerError(ctx, "Failed to get unit.", nil)
			}
		}

		// can be automatic build
		if unit == nil {
			if unitType := nokocore.ToTitleCase(productBody.UnitType); unitType != "" {
				if unit, err = unitRepository.SafeFirst("unit_type = ?", unitType); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyInternalServerError(ctx, "Failed to get unit.", nil)
				}

				if unit == nil {
					unit = &models2.Unit{
						UnitType: unitType,
					}
					if err = unitRepository.SafeCreate(unit); err != nil {
						console.Error(fmt.Sprintf("panic: %s", err.Error()))
						return extras.NewMessageBodyInternalServerError(ctx, "Failed create unit.", nil)
					}
				}

			} else {
				return extras.NewMessageBodyNotFound(ctx, "Unit not found.", nil)
			}
		}

		product.UnitID = unit.ID
		product.Unit = *unit

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

func GetAllProductByName(DB *gorm.DB) echo.HandlerFunc {

	// productRepository := repositories2.NewProductRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		nokocore.KeepVoid(err)

		keywords := extras.ParseQueryToString(ctx, "keywords")
		size, _ := strconv.Atoi(extras.ParseQueryToString(ctx, "size"))
		page, _ := strconv.Atoi(extras.ParseQueryToString(ctx, "page"))

		offset := (page - 1) * size

		var products []models2.Product
		if err = DB.Where("brand LIKE ? OR product_name LIKE ? OR barcode LIKE ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").Preload("Package").Preload("Unit").Limit(size).Offset(offset).Find(&products).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get products.", nil)
		}
		fmt.Println("= products", products)

		var productResults []schemas2.ProductResult
		for _, product := range products {
			productResults = append(productResults, schemas2.ToProductResult(&product))
		}

		return extras.NewMessageBodyOk(ctx, "Successfully get products.", &nokocore.MapAny{
			"products": productResults,
		})
	}
}

func ProductController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/products", CreateProduct(DB))
	group.GET("/products", GetAllProductByName(DB))

	return group
}
