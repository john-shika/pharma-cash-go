package repositories

// import (
// 	"fmt"
// 	"github.com/labstack/echo/v4"
// 	"gorm.io/gorm"
// 	"nokowebapi/apis/extras"
// 	"nokowebapi/nokocore"
// 	"pharma-cash-go/app/models"
// 	// "github.com/go-playground/validator/v10"
// )

// func CreateCategory(DB *gorm.DB) echo.HandlerFunc {
// 	var err error
// 	category := new(models.Category)
// 	nokocore.KeepVoid(err, category)

// 	return func(ctx echo.Context) error {
// 		if err = ctx.Bind(&category); err != nil {
// 			return extras.NewMessageBodyBadRequest(ctx, "Invalid request body.", err)
// 		}

// 		if err = DB.Save(&category).Error; err != nil {
// 			fmt.Println(err)

// 			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create category.", nil)
// 		}

// 		return extras.NewMessageBodyOk(ctx, "Successfully create category.", category)
// 	}
// }