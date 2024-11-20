package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/repositories"
)

func ProfileHandler(userRepository repositories.UserRepositoryImpl) echo.HandlerFunc {
	nokocore.KeepVoid(userRepository)

	return func(ctx echo.Context) error {
		jwtClaimsDataAccess := ctx.Get("jwt_claims_data_access").(nokocore.JwtClaimsDataAccessImpl)

		fmt.Println(jwtClaimsDataAccess)

		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", jwtClaimsDataAccess)
	}
}

func UserController(group *echo.Group, DB *gorm.DB) *echo.Group {

	userRepository := repositories.NewUserRepository(DB)

	group.GET("/profile", ProfileHandler(userRepository))

	return group
}
