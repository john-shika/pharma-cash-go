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
			"user": schemas.ToUserResult(&jwtAuthInfo.Session.User, nil),
		})
	}
}

func SessionHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	return func(ctx echo.Context) error {
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		userResult := schemas.ToUserResult(jwtAuthInfo.User, nil)
		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
			"session": schemas.ToSessionResult(jwtAuthInfo.Session, userResult),
		})
	}
}

func LogoutHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	sessionRepository := repositories.NewSessionRepository(DB)

	return func(ctx echo.Context) error {
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		sessionId := jwtAuthInfo.Session.UUID
		if err := sessionRepository.SafeDelete(jwtAuthInfo.Session, "uuid = ?", sessionId); err != nil {
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
