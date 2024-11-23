package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/apis/schemas"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/repositories"
	"strings"
)

func CreateUserHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	userRepository := repositories.NewUserRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var user *models.User
		nokocore.KeepVoid(err, user)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if apis.RoleIsAdmin(jwtAuthInfo) {

			userBody := new(schemas.UserBody)
			if err = ctx.Bind(userBody); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))

				return extras.NewMessageBodyInternalServerError(ctx, "Invalid request body.", nil)
			}

			if err = ctx.Validate(userBody); err != nil {
				return err
			}

			username := strings.TrimSpace(userBody.Username)
			userBody.Username = username
			if username != "" {
				if user, err = userRepository.SafeFirst("username = ?", userBody.Username); user != nil {
					return extras.NewMessageBodyUnprocessableEntity(ctx, "Username already exists.", nil)
				}
			}

			email := strings.TrimSpace(userBody.Email)
			userBody.Email = email
			if email != "" {
				if user, err = userRepository.SafeFirst("email = ?", email); user != nil {
					return extras.NewMessageBodyUnprocessableEntity(ctx, "Email already exists.", nil)
				}
			}

			phone := strings.TrimSpace(userBody.Phone)
			userBody.Phone = phone
			if phone != "" {
				if user, err = userRepository.SafeFirst("phone = ?", phone); user != nil {
					return extras.NewMessageBodyUnprocessableEntity(ctx, "Phone already exists.", nil)
				}
			}

			user = schemas.ToUserModel(userBody)
			if err = userRepository.SafeCreate(user); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))

				return extras.NewMessageBodyInternalServerError(ctx, "Failed to create a new user.", nil)
			}

			return extras.NewMessageBodyOk(ctx, "Successfully created a new user.", &nokocore.MapAny{
				"user": schemas.ToUserResult(user, nil),
			})
		}

		return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
	}
}

func AdminController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/user", CreateUserHandler(DB))

	return group
}
