package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
)

func ProfileHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	//sessionRepository := repositories.NewSessionRepository(DB)

	return func(ctx echo.Context) error {
		jwtClaimsDataAccess := ctx.Get("jwt_claims_data_access").(nokocore.JwtClaimsDataAccessImpl)
		session := ctx.Get("session").(*models.Session)

		//identity := jwtClaimsDataAccess.GetIdentity()
		//sessionRepository.SafeFirst("uuid = ?", identity)

		// Guest;Admin;Enterprise;TeamKit;Developer

		// validate session

		fmt.Println(session.User)

		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", jwtClaimsDataAccess)
	}
}

func UserController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.GET("/profile", ProfileHandler(DB))

	return group
}
