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
			stockOpnameResult := schemas2.ToStockOpnameResult(stockOpname)
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

		stockOpnameResult := schemas2.ToStockOpnameResult(stockOpnameNew)
		return extras.NewMessageBodyOk(ctx, "Successfully create checkpoint opname cart.", &nokocore.MapAny{
			"stockOpname": stockOpnameResult,
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
		fmt.Println("productID", productID)
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
			RealPackageTotal: cartVerificationOpnameBody.RealPackageTotal,
			RealUnitExtra:    cartVerificationOpnameBody.RealUnitExtra,
			RealUnitTotal:    (cartVerificationOpnameBody.RealPackageTotal * product.UnitAmount) + cartVerificationOpnameBody.RealUnitExtra,
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

func StokOpnameController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/warehouse/checkpoint", CreateCheckpointOpnameCart(DB))
	group.POST("/warehouse/cart/not-match/:productId", NotMatchVerification(DB))

	return group
}
