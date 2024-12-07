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
	"strings"
	"time"

	"github.com/google/uuid"
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
				pkg.uuid AS package_uuid,
				pkg.package_type AS package_type,
				p.package_total,
				u.uuid AS unit_uuid,
    			u.unit_type AS unit_type,
				p.unit_scale,
				p.unit_extra,
				(p.package_total * p.unit_scale) + p.unit_extra AS unit_total,
				COALESCE(cvo.is_match, TRUE) AS is_match,
				COALESCE(cvo.uuid, NULL) AS cart_stock_opname_id,
				COALESCE(cvo.not_match_reason, NULL) AS not_match_reason,
				p.created_at,
				p.updated_at
			FROM
				products p
			LEFT JOIN
				cart_verification_opnames cvo ON p.id = cvo.product_id
			LEFT JOIN
				packages pkg ON p.package_id = pkg.id
			LEFT JOIN 
    			units u ON p.unit_id = u.id;
		`

		if err = DB.Raw(query).Scan(&stockOpnamesResultGet).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get stock_opnames.", err.Error())
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
		if err = cartVerificationOpnameRepository.Create(&models2.CartVerificationOpname{
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
			return extras.NewMessageBodyBadRequest(ctx, "Failed to load cart_verification_opnames data with related cartVerificationOpnameId.", err.Error())
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

		if err = DB.Save(&cartVerificationOpnames).Error; err != nil {
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

func VerifyStockOpname(DB *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error
		var verificationOpnames []*models2.VerificationOpname
		var stockOpnamesResultGetVerfies []schemas2.StockOpnameResultGetVerify
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		// // check: is table cart_verification_opname empty
		// if err = DB.Find(&cartVerificationOpnames).Error; err != nil {
		// 	console.Error(fmt.Sprintf("panic: %s", err.Error()))
		// 	return extras.NewMessageBodyInternalServerError(ctx, "Failed to get cart_verification_opnames data.", err.Error())
		// }

		// if len(cartVerificationOpnames) == 0 {
		// 	return extras.NewMessageBodyBadRequest(ctx, "Failed to verify stock opname, there is not cart_verification_opnames data.", nil)
		// }

		// check:is there no table stock_opname data with isVerified == false
		var stockOpname models2.StockOpname
		if err = DB.First(&stockOpname, "is_verified = ?", false).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyBadRequest(ctx, "There is no stock_opnames data, create checkpoint first.", err.Error())
		}

		// get all
		query := `
			SELECT
				p.uuid AS product_uuid,
				p.barcode,
				p.product_name,
				p.brand,
				p.package_total AS system_package_total,
				p.unit_scale AS system_unit_scale,
				p.unit_extra AS system_unit_extra,
				(p.package_total * p.unit_scale) + p.unit_extra AS system_unit_total,
				COALESCE(cvo.is_match, TRUE) AS is_match,
				COALESCE(cvo.uuid, NULL) AS cart_stock_opname_id,
				COALESCE(cvo.not_match_reason, NULL) AS not_match_reason,
				COALESCE(cvo.real_package_total, NULL) AS real_package_total,
				COALESCE(cvo.real_unit_extra, NULL) AS real_unit_extra,
				(real_package_total * p.unit_scale) + real_unit_extra AS real_unit_total,
				p.created_at,
				p.updated_at
			FROM
				products p
			LEFT JOIN
				cart_verification_opnames cvo
			ON
				p.id = cvo.product_id;
		`

		if err = DB.Raw(query).Scan(&stockOpnamesResultGetVerfies).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get stock_opnames.", err.Error())
		}

		// // CREATE CART
		// // select all: from table prodcuts
		// if err = DB.Find(&products).Error; err != nil {
		// 	console.Error(fmt.Sprintf("panic: %s", err.Error()))
		// 	return errors.New("failed to get all products data")
		// }

		err = DB.Transaction(func(tx *gorm.DB) error {

			// prepared data
			//
			ids := []uuid.UUID{}
			casesPackageTotal := "CASE uuid"
			casesUnitExtra := "CASE uuid"
			//
			for _, stockOpnamesResultGetVerify := range stockOpnamesResultGetVerfies {
				if !stockOpnamesResultGetVerify.IsMatch {
					ids = append(ids, stockOpnamesResultGetVerify.ProductUUID)
					casesPackageTotal += fmt.Sprintf(" WHEN '%s' THEN '%d'", stockOpnamesResultGetVerify.ProductUUID, stockOpnamesResultGetVerify.RealPackageTotal)
					casesUnitExtra += fmt.Sprintf(" WHEN '%s' THEN '%d'", stockOpnamesResultGetVerify.ProductUUID, stockOpnamesResultGetVerify.RealUnitExtra)

				}
				verificationOpnames = append(verificationOpnames, &models2.VerificationOpname{
					ProductID:          stockOpnamesResultGetVerify.ProductUUID,
					StockOpnameID:      stockOpname.ID,
					SystemPackageTotal: stockOpnamesResultGetVerify.SystemPackageTotal,
					SystemUnitExtra:    stockOpnamesResultGetVerify.SystemUnitExtra,
					SystemUnitTotal:    stockOpnamesResultGetVerify.SystemUnitTotal,
					IsMatch:            stockOpnamesResultGetVerify.IsMatch,
					NotMatchReason:     stockOpnamesResultGetVerify.NotMatchReason,
					RealPackageTotal:   stockOpnamesResultGetVerify.RealPackageTotal,
					RealUnitExtra:      stockOpnamesResultGetVerify.RealUnitExtra,
					RealUnitTotal:      stockOpnamesResultGetVerify.RealUnitTotal,
					UserID:             jwtAuthInfo.User.ID,
				})

			}
			casesPackageTotal += " END"
			casesUnitExtra += " END"

			var idList []string
			for _, id := range ids {
				idList = append(idList, fmt.Sprintf("'%s'", id))
			}
			idsString := strings.Join(idList, ",")

			// insert: to table verification_opnames
			if err = DB.Create(&verificationOpnames).Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("failed to create cart_verification_opnames data")
			}

			// bulk update: change the package_total and unit_extra from table products
			updatedAt := time.Now().UTC().Format("2006-01-02 15:04:05.9999999-07:00")
			rawQuery := fmt.Sprintf(`
				UPDATE products
				SET 
					package_total = %s,
					unit_extra = %s,
					updated_at = '%s'
				WHERE uuid IN (%s) 
			`,
				casesPackageTotal, casesUnitExtra, updatedAt, idsString)
			fmt.Println("= rawQuery: ", rawQuery)

			if err := DB.Exec(rawQuery).Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("failed to update product package_total and unit_extra data")
			}

			// empty: delete all data from table cart_verification_opnames
			if err = DB.Unscoped().Where("1 = 1").Delete(&models2.CartVerificationOpname{}).Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("failed to delete all cart_verification_opnames data")
			}

			// update: submited_at and is_verified from table stock_opname
			if err = DB.Model(&stockOpname).Updates(map[string]interface{}{"submited_at": time.Now(), "is_verified": true}).Error; err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("failed to update submited_at and is_verified from table stock_opname")
			}

			return nil
		})

		if err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Failed to verify all stock opname.", err.Error())
		}

		return extras.NewMessageBodyOk(ctx, "Successfully update verify all product in table product and all nor match product in cart_verification_opnames, then inserted to verification_opname.", &nokocore.MapAny{
			"lengthProductVerified": len(verificationOpnames),
		})
	}
}

func GetHistoryStockOpnameDates(DB *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		type DateEntry struct {
			StockOpnameId uuid.UUID `json:"stockOpnameId"`
			Date          string    `json:"date"`
		}

		type YearlyData struct {
			Year  string      `json:"year"`
			Dates []DateEntry `json:"dates"`
		}

		// Mendapatkan parameter year dari query
		year := ctx.QueryParam("year")

		// Konstruksi kueri
		query := DB.Table("stock_opnames").Select("strftime('%Y', created_at) AS year, uuid AS stock_opname_id, created_at AS date")
		if year != "" {
			query = query.Where("strftime('%Y', created_at) = ?", year)
		}

		var rawResults []struct {
			Year          string
			StockOpnameId uuid.UUID
			Date          string
		}

		if err = query.Find(&rawResults).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Failed to get history stock opname dates.", err.Error())
		}

		groupedData := make(map[string][]DateEntry)
		for _, row := range rawResults {
			groupedData[row.Year] = append(groupedData[row.Year], DateEntry{
				StockOpnameId: row.StockOpnameId,
				Date:          row.Date,
			})
		}

		var result []YearlyData
		for year, dates := range groupedData {
			result = append(result, YearlyData{
				Year:  year,
				Dates: dates,
			})
		}

		return extras.NewMessageBodyOk(ctx, "Successfully get history stock opname dates.", result)
	}
}

func StokOpnameController(group *echo.Group, DB *gorm.DB) *echo.Group {

	// submenu verification
	group.POST("/warehouse/checkpoint", CreateCheckpointOpnameCart(DB))
	group.GET("/warehouse/cart", GetAllStockOpnames(DB))
	group.GET("/warehouse/stock/:productId", GetProductDetailForPopUpNotMatchVerification(DB))
	group.POST("/warehouse/cart/not-match/:productId", NotMatchVerification(DB))
	group.GET("/warehouse/cart/not-match/:cartVerificationOpnameId", GetNotMatchVerificationByCartVerificationOpnameId(DB))
	group.PUT("/warehouse/cart/not-match/:cartVerificationOpnameId", UpdateNotMatchVerificationByCartVerificationOpnameId(DB))
	group.DELETE("/warehouse/cart/not-match/:cartVerificationOpnameId", DeleteCartVerificationOpnameByCartVerificationOpnameId(DB))
	group.POST("/warehouse/cart/verify", VerifyStockOpname(DB))

	// submenu history
	group.GET("/warehouse/history/dates", GetHistoryStockOpnameDates(DB))

	return group
}
