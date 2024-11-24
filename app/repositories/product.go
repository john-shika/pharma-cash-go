package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nokowebapi/apis/extras"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/models"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateProduct(DB *gorm.DB) echo.HandlerFunc {
	var err error
	var product models.Product
	nokocore.KeepVoid(err, product)

	return func(ctx echo.Context) error {
		// Read the raw request body
		requestBody := make(map[string]interface{})
		if err = ctx.Bind(&requestBody); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		// Check if "expired" exists and transform it
		if expiredRaw, ok := requestBody["expired"].(string); ok {
			layout := "2006-01-02"
			parsedTime, err := time.Parse(layout, expiredRaw)
			if err != nil {
				return extras.NewMessageBodyBadRequest(ctx, "Invalid date format for expired. Use YYYY-MM-DD.", err)
			}
			// Replace the "expired" field with the correct RFC3339 format
			requestBody["expired"] = parsedTime
		} else {
			return extras.NewMessageBodyBadRequest(ctx, "Expired field is missing or invalid.", nil)
		}

		// Manually inject the updated body back into the context
		ctx.SetRequest(ctx.Request().Clone(ctx.Request().Context()))
		newBody, _ := json.Marshal(requestBody)
		ctx.Request().Body = ioutil.NopCloser(bytes.NewReader(newBody))

		if err = ctx.Bind(&product); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body when binding.", err)
		}

		// validasi: isPackagingID, isUnitID
		var packaging models.Packaging
		if err := DB.First(&packaging, product.PackagingID).Error; err != nil {
			return extras.NewMessageBodyNotFound(ctx, "Packaging not found.", err)
		}
		var unit models.Unit
		if err := DB.First(&unit, product.UnitID).Error; err != nil {
			return extras.NewMessageBodyNotFound(ctx, "Unit not found.", err)
		}

		// validasi: isEmpty
		if err := validate.Struct(product); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body when validate.", err.Error())
		}

		// query: create
		if err = DB.Create(&product).Error; err != nil {
			fmt.Println(err)

			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create product.", nil)
		}

		return extras.NewMessageBodyOk(ctx, "Successfully create product.", product)
	}
}
