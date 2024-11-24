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

func CreateUnit(DB *gorm.DB) echo.HandlerFunc {
	var err error
	unit := new(models.Unit)
	nokocore.KeepVoid(err, unit)

	return func(ctx echo.Context) error {
		if err = ctx.Bind(&unit); err != nil {
			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
		}

		if err = DB.Save(&unit).Error; err != nil {
			fmt.Println(err)

			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create unit.", nil)
		}

		return extras.NewMessageBodyOk(ctx, "Successfully create unit.", unit)
	}
}