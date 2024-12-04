package controllers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/utils"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
	repositories2 "pharma-cash-go/app/repositories"
	schemas2 "pharma-cash-go/app/schemas"
)

func GetAllCarts(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	cartRepository := repositories2.NewCartRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var carts []models2.Cart
		nokocore.KeepVoid(err, carts)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		user := jwtAuthInfo.User
		pagination := extras.NewURLQueryPaginationFromEchoContext(ctx)
		preloads := []string{"Product", "Product.Categories", "Product.Package", "Product.Unit"}
		if carts, err = cartRepository.SafePreMany(preloads, pagination.Offset, pagination.Limit, "user_id = ? AND closed = FALSE", user.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get carts.", nil)
		}

		size := len(carts)
		cartResults := make([]schemas2.CartResult, size)
		for i, cart := range carts {
			nokocore.KeepVoid(i)
			cartResult := schemas2.ToCartResult(&cart)
			cartResults[i] = cartResult
		}

		return extras.NewMessageBodyOk(ctx, "Successfully get carts.", &nokocore.MapAny{
			"carts": cartResults,
		})
	}
}

func ProductCheckout(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	productRepository := repositories2.NewProductRepository(DB)
	cartRepository := repositories2.NewCartRepository(DB)
	transactionRepository := repositories2.NewTransactionRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var productID string
		var transactionID string
		var product *models2.Product
		var transaction *models2.Transaction
		nokocore.KeepVoid(err, transactionID, product, transaction)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)
		user := jwtAuthInfo.User
		userID := user.ID

		if !utils.RoleIsAdmin(jwtAuthInfo) && !utils.RoleIs(jwtAuthInfo, nokocore.RoleOfficer) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		if productID = extras.ParseQueryToString(ctx, "product_id"); productID != "" {
			if err = sqlx.ValidateUUID(productID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Invalid parameter 'product_id'.", nil)
			}
		}

		if transactionID = extras.ParseQueryToString(ctx, "transaction_id"); transactionID != "" {
			if err = sqlx.ValidateUUID(transactionID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Invalid parameter 'transaction_id'.", nil)
			}
		}

		cartBody := new(schemas2.CartBody)

		if err = ctx.Bind(cartBody); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to bind cart.", nil)
		}

		// passing validation
		if productID != "" {
			cartBody.ProductID = uuid.MustParse(productID)
		}

		// passing validation
		if transactionID != "" {
			cartBody.TransactionID = uuid.MustParse(transactionID)
		}

		fmt.Println(nokocore.ShikaYamlEncode(cartBody))

		if err = ctx.Validate(cartBody); err != nil {
			return err
		}

		preloads := []string{"Categories", "Package", "Unit"}
		if product, err = productRepository.SafePreFirst(preloads, "uuid = ?", productID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get product.", nil)
		}

		if product == nil {
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Product not found.", nil)
		}

		if transactionID != "" {
			if transaction, err = transactionRepository.SafeFirst("uuid = ? AND closed = FALSE", transactionID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get transaction.", nil)
			}

			if transaction == nil {
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Transaction not found.", nil)
			}

			if transaction.UserID != userID {
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Transaction not found.", nil)
			}
		}

		if transaction == nil {
			if transaction, err = transactionRepository.SafeFirst("user_id = ? AND closed = FALSE", userID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get transaction.", nil)
			}
		}

		if transaction == nil {
			transaction = &models2.Transaction{
				UserID: userID,
				Pay:    decimal.NewFromInt(0),
				Signed: false,
				Closed: false,
			}

			if err = transactionRepository.Create(transaction); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to create transaction.", nil)
			}
		}

		cart := schemas2.ToCartModelWithProductModel(cartBody, product)

		// set owner and transaction
		cart.UserID = userID
		cart.TransactionID = transaction.ID

		unitTotal := product.UnitScale * cart.PackageTotal
		unitTotal += cart.UnitExtra

		// inject current product
		cart.ProductID = product.ID
		cart.Product = *product

		var check *models2.Cart

		if check, err = cartRepository.SafeFirst("user_id = ? AND product_id = ? AND closed = FALSE", userID, product.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get cart.", nil)
		}

		status := "add"

		if check != nil {
			// inject base model values
			cart.ID = check.ID
			cart.UUID = check.UUID
			cart.CreatedAt = check.CreatedAt

			if err = cartRepository.SafeUpdate(cart, "id = ?", cart.ID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to update cart.", nil)
			}

			status = "update"

		} else {
			if err = cartRepository.Create(cart); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to create cart.", nil)
			}
		}

		qty := decimal.NewFromInt(int64(unitTotal))
		pay := product.SalePrice.Mul(qty)

		transaction.Pay = pay
		if err = transactionRepository.SafeUpdate(transaction, "id = ?", transaction.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to update transaction.", nil)
		}

		cartResult := schemas2.ToCartResult(cart)
		transactionResult := schemas2.ToTransactionResult(transaction)
		return extras.NewMessageBodyOk(ctx, fmt.Sprintf("Successfully %s cart.", status), &nokocore.MapAny{
			"cart":        cartResult,
			"transaction": transactionResult,
			"unitTotal":   unitTotal,
			"pay":         pay,
		})
	}
}

func ShopController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.GET("/carts", GetAllCarts(DB))
	group.POST("/product/checkout", ProductCheckout(DB))

	return group
}
