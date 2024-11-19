package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/schemas"
	"nokowebapi/apis/validators"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/repositories"
)

func MessageHandler(userRepository repositories.UserRepository) echo.HandlerFunc {
	nokocore.KeepVoid(userRepository)

	return func(ctx echo.Context) error {
		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", nokocore.MapAny{
			"message": "Hay!",
		})
	}
}

func LoginHandler(userRepository repositories.UserRepository) echo.HandlerFunc {
	nokocore.KeepVoid(userRepository)

	return func(ctx echo.Context) error {
		var err error
		nokocore.KeepVoid(err)

		user := new(schemas.UserBody)
		if err = ctx.Bind(user); err != nil {
			return err
		}

		if err = ctx.Validate(user); err != nil {
			return err
		}

		if err = validators.ValidatePass(user.Password); err != nil {
			return err
		}

		console.Dir(user)

		return nil
	}
}

func AnonymousController(group *echo.Group, db *gorm.DB) *echo.Group {

	userRepository := repositories.NewUserRepository(db)

	group.GET("/message", MessageHandler(userRepository))
	group.POST("/login", LoginHandler(userRepository))

	return group
}
