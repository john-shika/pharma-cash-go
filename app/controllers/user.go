package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/schemas"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/repositories"
)

func ProfileHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	return func(ctx echo.Context) error {
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
			"user": schemas.ToUserResp(&jwtAuthInfo.Session.User),
		})
	}
}

func SessionHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	return func(ctx echo.Context) error {
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
			"session": schemas.ToSessionResp(jwtAuthInfo.Session),
		})
	}
}

func LogoutHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	sessionRepository := repositories.NewSessionRepository(DB)

	return func(ctx echo.Context) error {
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if err := sessionRepository.SafeDelete(jwtAuthInfo.Session); err != nil {
			return err
		}

		return extras.NewMessageBodyOk(ctx, "Successfully logged out.", nil)
	}
}

func UserController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.GET("/profile", ProfileHandler(DB))
	group.GET("/session", SessionHandler(DB))
	group.POST("/logout", LogoutHandler(DB))

	return group
}
