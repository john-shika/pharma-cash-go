package controllers

import (
	"github.com/labstack/echo/v4"
	"nokowebapi/apis/extras"
	"nokowebapi/nokocore"
)

func MessageHandler(ctx echo.Context) error {
	return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", nokocore.MapAny{
		"message": "Hay!",
	})
}

func LoginHandler(ctx echo.Context) error {

	return nil
}

func AnonymousController(router *echo.Group) *echo.Group {

	router.GET("/message", MessageHandler)
	router.POST("/login", LoginHandler)

	return router
}
