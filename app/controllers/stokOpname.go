package controllers

import (
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

	return func(ctx echo.Context) error {
		var err error
		var stockOpname *models2.StockOpname
		nokocore.KeepVoid(err, stockOpname)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

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
			return extras.NewMessageBodyBadRequest(ctx, "Data has not been verified.", &nokocore.MapAny{
				"stockOpname": stockOpnameResult,
			})
		}

		stockOpnameNew := &models2.StockOpname{
			IsVerified: false,
			UserID:     uint(jwtAuthInfo.User.ID),
		}

		if err := DB.Create(stockOpnameNew).Error; err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create unit.", nil)
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

func StokOpnameController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/warehouse/checkpoint", CreateCheckpointOpnameCart(DB))

	return group
}
