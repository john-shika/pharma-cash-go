package repositories

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/models"
	// "github.com/go-playground/validator/v10"
)

func CreatePackaging(DB *gorm.DB) echo.HandlerFunc {
	var err error
	packaging := new(models.Packaging)
	nokocore.KeepVoid(err, packaging)

	return func(ctx echo.Context) error {
		if err = ctx.Bind(&packaging); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		if err = DB.Create(&packaging).Error; err != nil {
			fmt.Println(err)

			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create packaging.", nil)
		}

		return extras.NewMessageBodyOk(ctx, "Successfully create packaging.", packaging)
	}
}