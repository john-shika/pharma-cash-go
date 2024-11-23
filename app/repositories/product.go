package repositories

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/models"
)

func CreateProduct(DB *gorm.DB) echo.HandlerFunc {
	var err error
	product := new(models.Product)
	nokocore.KeepVoid(err, product)

	return func(ctx echo.Context) error {

		fmt.Println(product)
		if err = ctx.Bind(&product); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		fmt.Println(product)

		if err = ctx.Validate(&product); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err.Error())
		}

		if err = DB.Save(&product).Error; err != nil {
			fmt.Println(err)

			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create product.", nil)
		}

		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", product)
	}
}
