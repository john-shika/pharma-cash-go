package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/utils"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	models2 "pharma-cash-go/app/models"
	repositories2 "pharma-cash-go/app/repositories"
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
		preloads := []string{"Product"}
		if carts, err = cartRepository.SafePreMany(preloads, pagination.Offset, pagination.Limit, "user_id = ?", user.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get carts.", nil)
		}

		return nil
	}
}

func ShopController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.GET("/carts", GetAllCarts(DB))

	return group
}
