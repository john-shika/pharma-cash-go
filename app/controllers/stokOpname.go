package controllers

import (
	"errors"
	"fmt"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/utils"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
	repositories2 "pharma-cash-go/app/repositories"
	schemas2 "pharma-cash-go/app/schemas"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateCheckpointOpnameCart(DB *gorm.DB) echo.HandlerFunc {

	stokOpnameRepository := repositories2.NewStockRepository(DB)
	// productRepository := repositories2.NewProductRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var stockOpname *models2.StockOpname
		var stockOpnameNew *models2.StockOpname
		// var products []*models2.Product
		// var cartVerificationOpnames []*models2.CartVerificationOpname
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		nokocore.KeepVoid(err, stockOpname)

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		// stockOpnameBody := new(schemas2.StockOpnameBody)

		// if err = ctx.Bind(stockOpnameBody); err != nil {
		// 	return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		// }

		// if err = ctx.Validate(stockOpnameBody); err != nil {
		// 	return err
		// }

		preloads := []string{"User"}
		if stockOpname, err = stokOpnameRepository.SafePreFirst(preloads, "is_verified = ?", false); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get stcok_opname.", nil)
		}

		if stockOpname != nil {
			stockOpnameResult := schemas2.ToStockOpnameResultCreate(stockOpname)
			return extras.NewMessageBodyBadRequest(ctx, "Data has not been verified yet.", &nokocore.MapAny{
				"stockOpname": stockOpnameResult,
			})
		}

		err = DB.Transaction(func(tx *gorm.DB) error {

			// insert: to table stock_opnames
			stockOpnameNew = &models2.StockOpname{
				IsVerified: false,
				UserID:     uint(jwtAuthInfo.User.ID),
			}
			if err = DB.Create(stockOpnameNew).Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("failed to create stock_opname data")
			}

			// // CREATE CART
			// // select all: from table prodcuts
			// if err = DB.Find(&products).Error; err != nil {
			// 	console.Error(fmt.Sprintf("panic: %s", err.Error()))
			// 	return errors.New("failed to get all products data")

			// }
			// // prepared data
			// for _, product := range products {
			// 	cartVerificationOpnames = append(cartVerificationOpnames, &models2.CartVerificationOpname{
			// 		UserID:     uint(jwtAuthInfo.User.ID),
			// 		ProductID:  product.ID,
			// 		IsMatch:    product.PackageTotal,
			// 		AmountUnit: product.UnitAmount,
			// 		UnitExtra:  product.UnitExtra,
			// 		UnitTotal:  (product.PackageTotal * product.UnitAmount) + product.UnitExtra,
			// 	})
			// }

			// // insert: to table cart_verification_opnames
			// if err = DB.Create(&cartVerificationOpnames).Error; err != nil {
			// 	console.Error(fmt.Sprintf("panic: %s", err.Error()))
			// 	return errors.New("failed to create cart_verification_opnames data")
			// }

			return nil
		})

		if err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Failed to create checkpoint opname cart.", nil)
		}

		// Preload tabel User setelah data dibuat
		if err := DB.Preload("User").First(stockOpnameNew, stockOpnameNew.ID).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to load related user.", nil)
		}

		stockOpnameResult := schemas2.ToStockOpnameResultCreate(stockOpnameNew)
		return extras.NewMessageBodyOk(ctx, "Successfully create checkpoint opname cart.", &nokocore.MapAny{
			"stockOpname": stockOpnameResult,
		})
	}
}

func GetAllStockOpnames(DB *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error
		var stockOpnamesResultGet []schemas2.StockOpnameResultGet
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		query := `
			SELECT
				p.uuid AS product_uuid,
				p.barcode,
				p.product_name,
				p.brand,
				p.package_total,
				p.unit_amount,
				p.unit_extra,
				(p.package_total * p.unit_amount) + p.unit_extra AS unit_total,
				COALESCE(cvo.is_match, TRUE) AS is_match,
				COALESCE(cvo.uuid, NULL) AS cart_stock_opname_id,
				p.created_at,
				p.updated_at
			FROM
				products p
			LEFT JOIN
				cart_verification_opnames cvo
			ON
				p.id = cvo.product_id;
		`

		if err = DB.Raw(query).Scan(&stockOpnamesResultGet).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get stock_opnames.", nil)
		}

		// stockOpnamesResult := schemas2.ToStockOpnamesResult(stockOpnames)
		return extras.NewMessageBodyOk(ctx, "Successfully get all stock_opnames.", &nokocore.MapAny{
			"stockOpnames": stockOpnamesResultGet,
		})
	}
}

func GetProductDetailForPopUpNotMatchVerification(DB *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error
		var product *models2.Product
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		productID := ctx.Param("productId")

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		if err = DB.Preload("Package").Preload("Unit").First(&product, "UUID = ?", productID).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyBadRequest(ctx, "Unable to get product data.", err.Error())
		}

		productResult := schemas2.ToProductResult(product)
		return extras.NewMessageBodyOk(ctx, "Successfully get product detail for pop up not match.", &nokocore.MapAny{
			"product": productResult,
		})
	}
}

func NotMatchVerification(DB *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error
		var product *models2.Product
		var cartVerificationOpname *models2.CartVerificationOpname
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		productID := ctx.Param("productId")

		cartVerificationOpnameRepository := repositories2.NewCartVerificationOpnameRepository(DB)

		// check: is productId exist
		if err = DB.First(&product, "UUID = ?", productID).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyBadRequest(ctx, "Unable to get product data.", err.Error())
		}

		// check: is productId already registered
		if err = DB.Preload("Product").First(&cartVerificationOpname, "product_id = ?", product.ID).Error; err == nil {
			console.Error(fmt.Sprintf("panic: %s", err))
			return extras.NewMessageBodyBadRequest(ctx, "Already exists cart_verification_opnames data with productId.", err)
		}

		// binding json
		cartVerificationOpnameBody := &schemas2.CartVerificationOpnameBody{}
		if err = ctx.Bind(cartVerificationOpnameBody); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		// insert: to table cart_verification_opnames
		if err = cartVerificationOpnameRepository.SafeCreate(&models2.CartVerificationOpname{
			UserID:           uint(jwtAuthInfo.User.ID),
			ProductID:        product.ID,
			IsMatch:          false,
			NotMatchReason:   cartVerificationOpnameBody.NotMatchReason,
			RealPackageTotal: cartVerificationOpnameBody.RealPackageTotal,
			RealUnitExtra:    cartVerificationOpnameBody.RealUnitExtra,
			// RealUnitTotal:    (cartVerificationOpnameBody.RealPackageTotal * product.UnitAmount) + cartVerificationOpnameBody.RealUnitExtra,
		}); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Failed to create cart_verification_opnames data.", err.Error())
		}

		// Preload tabel User setelah data dibuat
		var newCartVerificationOpname *models2.CartVerificationOpname
		if err := DB.Preload("User").Preload("Product").First(&newCartVerificationOpname, "product_id = ?", product.ID).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to load cart_verification_opnames data with related productId.", err.Error())
		}

		cartVerificationOpnameResult := schemas2.ToCartVerificationOpnameResult(newCartVerificationOpname)
		return extras.NewMessageBodyOk(ctx, "Successfully create cart_verification_opnames data.", cartVerificationOpnameResult)

	}
}

func GetNotMatchVerificationByCartVerificationOpnameId(DB *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error
		var cartVerificationOpnames *models2.CartVerificationOpname
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)
		cartVerificationOpnameId := ctx.Param("cartVerificationOpnameId")

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		if err = DB.Preload("User").Preload("Product").First(&cartVerificationOpnames, "uuid = ?", cartVerificationOpnameId).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to load cart_verification_opnames data with related cartVerificationOpnameId.", err.Error())
		}

		cartVerificationOpnameResult := schemas2.ToCartVerificationOpnameResult(cartVerificationOpnames)
		return extras.NewMessageBodyOk(ctx, "Successfully load cart_verification_opnames data.", cartVerificationOpnameResult)

	}
}

func UpdateNotMatchVerificationByCartVerificationOpnameId(DB *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error
		var cartVerificationOpnames *models2.CartVerificationOpname
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)
		cartVerificationOpnameId := ctx.Param("cartVerificationOpnameId")

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		if err = DB.Preload("User").Preload("Product").First(&cartVerificationOpnames, "uuid = ?", cartVerificationOpnameId).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to load cart_verification_opnames data with related cartVerificationOpnameId.", err.Error())
		}

		cartVerificationOpnameBody := new(schemas2.CartVerificationOpnameBody)
		if err = ctx.Bind(&cartVerificationOpnameBody); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		cartVerificationOpnames.NotMatchReason = cartVerificationOpnameBody.NotMatchReason
		cartVerificationOpnames.RealPackageTotal = cartVerificationOpnameBody.RealPackageTotal
		cartVerificationOpnames.RealUnitExtra = cartVerificationOpnameBody.RealUnitExtra

		if err = DB.Model(&cartVerificationOpnames).Update("is_match", true).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to update cart_verification_opnames data with related cartVerificationOpnameId.", err.Error())
		}

		cartVerificationOpnameResult := schemas2.ToCartVerificationOpnameResult(cartVerificationOpnames)
		return extras.NewMessageBodyOk(ctx, "Successfully update cart_verification_opnames data with related cartVerificationOpnameId.", cartVerificationOpnameResult)
	}
}

func DeleteCartVerificationOpnameByCartVerificationOpnameId(DB *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error
		var cartVerificationOpnames *models2.CartVerificationOpname
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)
		cartVerificationOpnameId := ctx.Param("cartVerificationOpnameId")

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		if err = DB.Preload("User").Preload("Product").First(&cartVerificationOpnames, "uuid = ?", cartVerificationOpnameId).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to load cart_verification_opnames data with related cartVerificationOpnameId.", err.Error())
		}

		if err = DB.Unscoped().Delete(&cartVerificationOpnames).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to delete cart_verification_opnames data with related cartVerificationOpnameId.", err.Error())
		}

		return extras.NewMessageBodyOk(ctx, "Successfully delete cart_verification_opnames data with related cartVerificationOpnameId.", nil)
	}
}

func StokOpnameController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/warehouse/checkpoint", CreateCheckpointOpnameCart(DB))
	group.GET("/warehouse/cart", GetAllStockOpnames(DB))
	group.GET("/warehouse/stock/:productId", GetProductDetailForPopUpNotMatchVerification(DB))
	group.POST("/warehouse/cart/not-match/:productId", NotMatchVerification(DB))
	group.GET("/warehouse/cart/not-match/:cartVerificationOpnameId", GetNotMatchVerificationByCartVerificationOpnameId(DB))
	group.PUT("/warehouse/cart/not-match/:cartVerificationOpnameId", UpdateNotMatchVerificationByCartVerificationOpnameId(DB))
	group.DELETE("/warehouse/cart/not-match/:cartVerificationOpnameId", DeleteCartVerificationOpnameByCartVerificationOpnameId(DB))

	return group
}
