package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/schemas"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/models"
	"pharma-cash-go/app/repositories"
	schemas2 "pharma-cash-go/app/schemas"
)

func ProfileHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	employeeRepository := repositories.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var employee *models.Employee
		nokocore.KeepVoid(err, employee)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		preloads := []string{"Shift"}
		if employee, err = employeeRepository.SafePreFirst(preloads, "user_id = ?", jwtAuthInfo.User.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get employee.", nil)
		}

		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
			"user":  schemas.ToUserResult(&jwtAuthInfo.Session.User, nil),
			"shift": schemas2.ToShiftResult(&employee.Shift),
		})
	}
}

func SessionHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	employeeRepository := repositories.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var employee *models.Employee
		nokocore.KeepVoid(err, employee)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		preloads := []string{"Shift"}
		if employee, err = employeeRepository.SafePreFirst(preloads, "user_id = ?", jwtAuthInfo.User.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get employee.", nil)
		}

		var shift schemas2.ShiftResult
		if employee != nil {
			shift = schemas2.ToShiftResult(&employee.Shift)
		}

		userResult := schemas.ToUserResult(jwtAuthInfo.User, nil)
		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
			"session": schemas.ToSessionResult(jwtAuthInfo.Session, userResult),
			"shift":   shift,
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
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to log out.", nil)
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
