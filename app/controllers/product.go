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

	// productRepository := repositories2.NewProductRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		nokocore.KeepVoid(err)

		keywords := extras.ParseQueryToString(ctx, "keywords")

		pagination := extras.NewURLQueryPaginationFromEchoContext(ctx)

		var products []models2.Product
		query := "brand LIKE ? OR product_name LIKE ? OR barcode LIKE ?"
		args := []any{"%" + keywords + "%", "%" + keywords + "%", "%" + keywords + "%"}
		tx := DB.Preload("Package").Preload("Unit").Where(query, args...).Offset(pagination.Offset).Limit(pagination.Limit).Find(&products)
		if err = tx.Error; err != nil {
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

	return func(ctx echo.Context) error {
		paramProductId := ctx.Param("productId")

		// product, err := productRepository.SafeFirst("uuid = ?", paramProductId)
		var product models2.Product
		if err := DB.Where("uuid = ?", paramProductId).Preload("Categories").Preload("Package").Preload("Unit").First(&product).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get product detail.", nil)
		}

		productResult := schemas2.ToProductResult(&product)
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
		paramProductId := ctx.Param("productId")

		productBody := new(schemas2.ProductBody)
		if err := ctx.Bind(productBody); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Failed to bind product.", err.Error())
		}

		var err error
		nokocore.KeepVoid(err)

		product, err := productRepository.SafeFirst("uuid = ?", paramProductId)
		if err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get product.", err.Error())
		}
		fmt.Println("= productBody", productBody)

		packages, err := packageRepository.SafeFirst("uuid = ?", productBody.PackageID)
		if err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get package.", err.Error())
		}

		unit, err := unitRepository.SafeFirst("uuid = ?", productBody.UnitID)
		if err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to get unit.", err.Error())
		}

		newProductBody := schemas2.ToProductModel(productBody)
		newProductBody.PackageID = packages.ID
		newProductBody.UnitID = unit.ID
		fmt.Println("= newProductBody", newProductBody)

		product.Barcode = newProductBody.Barcode
		product.Brand = newProductBody.Brand
		product.ProductName = newProductBody.ProductName
		product.Supplier = newProductBody.Supplier
		product.Description = newProductBody.Description
		product.Categories = newProductBody.Categories
		product.Expires = newProductBody.Expires
		product.PurchasePrice = newProductBody.PurchasePrice
		product.SupplierDiscount = newProductBody.SupplierDiscount
		product.VAT = newProductBody.VAT
		product.ProfitMargin = newProductBody.ProfitMargin
		product.PackageID = newProductBody.PackageID
		product.PackageTotal = newProductBody.PackageTotal
		product.UnitID = newProductBody.UnitID
		product.UnitAmount = newProductBody.UnitAmount
		product.UnitExtra = newProductBody.UnitExtra

		fmt.Println("= product", product)
		if err := productRepository.SafeUpdate(product, "uuid = ?", paramProductId); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to update product.", err.Error())
		}

		return extras.NewMessageBodyOk(ctx, "Successfully update product.", &nokocore.MapAny{
			"product": schemas2.ToProductResult(product),
		})
	}
}

func DeleteProduct(DB *gorm.DB) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		paramProductId := ctx.Param("productId")
		// productId, _ := strconv.Atoi(paramProductId)

		result := DB.Model(models2.Product{}).Where("uuid = ?", paramProductId).Delete(&models2.Product{})
		if result.Error != nil {
			console.Error(fmt.Sprintf("panic: %s", result.Error.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to delete product.", nil)
		}

		if result.RowsAffected < 1 {
			return extras.NewMessageBodyNotFound(ctx, "Product not found.", nil)
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
