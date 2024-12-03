package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/utils"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
	repositories2 "pharma-cash-go/app/repositories"
	schemas2 "pharma-cash-go/app/schemas"
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

		// tax product price and income
		//margin := product.PurchasePrice.Mul(decimal.NewFromFloat(product.ProfitMargin))
		//result := product.PurchasePrice.Add(margin)
		//tax := result.Mul(decimal.NewFromFloat(product.VAT))
		//product.SalePrice = result.Add(tax)

		// only tax product price
		margin := product.PurchasePrice.Mul(decimal.NewFromFloat(product.ProfitMargin))
		tax := product.PurchasePrice.Mul(decimal.NewFromFloat(product.VAT))
		product.SalePrice = product.PurchasePrice.Add(margin).Add(tax)

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

func GetAllProductsByName(DB *gorm.DB) echo.HandlerFunc {

	productRepository := repositories2.NewProductRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var products []models2.Product
		nokocore.KeepVoid(err, products)

		keywords := extras.ParseQueryToString(ctx, "keywords")

		pagination := extras.NewURLQueryPaginationFromEchoContext(ctx)

		preloads := []string{"Categories", "Package", "Unit"}
		query := "brand LIKE ? OR product_name LIKE ? OR barcode LIKE ?"
		args := []any{"%" + keywords + "%", "%" + keywords + "%", "%" + keywords + "%"}
		if products, err = productRepository.SafePreMany(preloads, pagination.Offset, pagination.Limit, query, args...); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get products.", nil)
		}

		var productResults []schemas2.ProductResult
		for i, product := range products {
			nokocore.KeepVoid(i)
			productResults = append(productResults, schemas2.ToProductResult(&product))
		}

		return extras.NewMessageBodyOk(ctx, "Successfully get products.", &nokocore.MapAny{
			"products": productResults,
		})
	}
}

func GetProductDetailByProductId(DB *gorm.DB) echo.HandlerFunc {

	productRepository := repositories2.NewProductRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var product *models2.Product
		nokocore.KeepVoid(err, product)

		productID := ctx.Param("productId")

		preloads := []string{"Categories", "Package", "Unit"}
		if product, err = productRepository.SafePreFirst(preloads, "uuid = ?", productID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get product.", nil)
		}

		productResult := schemas2.ToProductResult(product)
		return extras.NewMessageBodyOk(ctx, "Successfully get product.", &nokocore.MapAny{
			"product": productResult,
		})
	}
}

func UpdateProduct(DB *gorm.DB) echo.HandlerFunc {

	productRepository := repositories2.NewProductRepository(DB)
	packageRepository := repositories2.NewPackageRepository(DB)
	unitRepository := repositories2.NewUnitRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var product *models2.Product
		var newProduct *models2.Product
		var packageModel *models2.Package
		var unit *models2.Unit
		nokocore.KeepVoid(err, product, packageModel, unit)

		productID := ctx.Param("productId")

		productBody := new(schemas2.ProductBody)
		if err = ctx.Bind(productBody); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Failed to bind product.", err.Error())
		}

		if err = ctx.Validate(productBody); err != nil {
			return err
		}

		newProduct = schemas2.ToProductModel(productBody)

		preloads := []string{"Categories", "Package", "Unit"}
		if product, err = productRepository.SafePreFirst(preloads, "uuid = ?", productID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get product.", err.Error())
		}

		if packageID := productBody.PackageID; packageID != "" {
			if packageModel, err = packageRepository.SafeFirst("uuid = ?", packageID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyInternalServerError(ctx, "Failed to get package.", err.Error())
			}
		}

		if packageModel == nil {
			if packageType := nokocore.ToTitleCase(productBody.PackageType); packageType != "" {
				if packageModel, err = packageRepository.SafeFirst("package_type = ?", packageType); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyInternalServerError(ctx, "Failed to get package.", err.Error())
				}

				if packageModel == nil {
					packageModel = &models2.Package{
						PackageType: packageType,
					}

					if err = packageRepository.SafeCreate(packageModel); err != nil {
						console.Error(fmt.Sprintf("panic: %s", err.Error()))
						return extras.NewMessageBodyInternalServerError(ctx, "Failed to create package.", err.Error())
					}
				}
			} else {
				return extras.NewMessageBodyNotFound(ctx, "Package not found.", nil)
			}
		}

		newProduct.PackageID = packageModel.ID
		newProduct.Package = *packageModel

		if unitID := productBody.UnitID; unitID != "" {
			if unit, err = unitRepository.SafeFirst("uuid = ?", unitID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyInternalServerError(ctx, "Failed to get unit.", err.Error())
			}
		}

		if unit == nil {
			if unitType := nokocore.ToTitleCase(productBody.UnitType); unitType != "" {
				if unit, err = unitRepository.SafeFirst("unit_type = ?", unitType); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyInternalServerError(ctx, "Failed to get unit.", err.Error())
				}

				if unit == nil {
					unit = &models2.Unit{
						UnitType: productBody.UnitType,
					}

					if err = unitRepository.SafeCreate(unit); err != nil {
						console.Error(fmt.Sprintf("panic: %s", err.Error()))
						return extras.NewMessageBodyInternalServerError(ctx, "Failed to create unit.", err.Error())
					}
				}
			} else {
				return extras.NewMessageBodyNotFound(ctx, "Unit not found.", nil)
			}
		}

		newProduct.UnitID = unit.ID
		newProduct.Unit = *unit

		newProduct.ID = product.ID
		newProduct.UUID = product.UUID
		newProduct.CreatedAt = product.CreatedAt

		if err = productRepository.SafeUpdate(newProduct, "id = ?", product.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to update product.", err.Error())
		}

		return extras.NewMessageBodyOk(ctx, "Successfully update product.", &nokocore.MapAny{
			"product": schemas2.ToProductResult(product),
		})
	}
}

func DeleteProduct(DB *gorm.DB) echo.HandlerFunc {

	productRepository := repositories2.NewProductRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var product *models2.Product
		nokocore.KeepVoid(err, product)

		productId := ctx.Param("productId")
		permanent := extras.ParseQueryToBool(ctx, "permanent")

		if product, err = productRepository.First("uuid = ?", productId); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get product.", nil)
		}

		if !permanent && product.DeletedAt.Valid {
			return extras.NewMessageBodyOk(ctx, "Product already deleted.", nil)
		}

		if product != nil {
			if !permanent {
				if err = productRepository.SafeDelete(product, "uuid = ?", productId); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyInternalServerError(ctx, "Failed to delete product.", nil)
				}
			} else {
				if err = productRepository.Delete(product, "uuid = ?", productId); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyInternalServerError(ctx, "Failed to delete product.", nil)
				}
			}
		}

		return extras.NewMessageBodyOk(ctx, "Successfully delete product.", nil)
	}
}

func ProductController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.GET("/products", GetAllProductsByName(DB))
	group.POST("/product", CreateProduct(DB))
	group.GET("/product/:productId", GetProductDetailByProductId(DB))
	group.PUT("/product/:productId", UpdateProduct(DB))
	group.DELETE("/product/:productId", DeleteProduct(DB))

	return group
}
