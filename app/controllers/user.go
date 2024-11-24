package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/apis/repositories"
	"nokowebapi/apis/schemas"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
	repositories2 "pharma-cash-go/app/repositories"
	schemas2 "pharma-cash-go/app/schemas"
)

func ProfileHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	employeeRepository := repositories2.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var employee *models2.Employee
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

		userResult := schemas.ToUserResult(&jwtAuthInfo.Session.User, nil)
		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
			"user":  userResult,
			"shift": shift,
		})
	}
}

func SessionHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	employeeRepository := repositories2.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var employee *models2.Employee
		nokocore.KeepVoid(err, employee)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		preloads := []string{"Shift"}
		if employee, err = employeeRepository.SafePreFirst(preloads, "user_id = ?", jwtAuthInfo.User.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get employee.", nil)
		}

		userResult := schemas.ToUserResult(jwtAuthInfo.User, nil)
		sessionResult := schemas.ToSessionResult(jwtAuthInfo.Session, userResult)
		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
			"session": sessionResult,
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

func RefreshTokenHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	jwtConfig := globals.GetJwtConfig()
	signingMethod := jwtConfig.GetSigningMethod()
	expiresIn := jwtConfig.GetExpiresIn()

	sessionRepository := repositories.NewSessionRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		nokocore.KeepVoid(err)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		user := jwtAuthInfo.User
		session := jwtAuthInfo.Session

		// get user roles
		roles := nokocore.RolesUnpack(user.Roles)

		sessionId := session.UUID
		timeUtcNow := nokocore.GetTimeUtcNow()
		expires := timeUtcNow.Add(expiresIn)

		jwtClaimsDataAccess := nokocore.NewEmptyJwtClaimsDataAccess()
		jwtClaimsDataAccess.SetSubject("NokoWebApiToken")
		jwtClaimsDataAccess.SetIssuer(jwtConfig.Issuer)
		jwtClaimsDataAccess.SetAudience(jwtConfig.Audience)
		jwtClaimsDataAccess.SetIssuedAt(timeUtcNow)
		jwtClaimsDataAccess.SetExpiresAt(expires)
		jwtClaimsDataAccess.SetUser(user.Username)
		jwtClaimsDataAccess.SetSessionId(sessionId.String())
		jwtClaimsDataAccess.SetRoles(roles)
		jwtClaimsDataAccess.SetAdmin(user.Admin)
		jwtClaimsDataAccess.SetLevel(user.Level)

		jwtClaims := nokocore.ToJwtClaims(jwtClaimsDataAccess, signingMethod)
		jwtToken := nokocore.GenerateJwtToken(jwtClaims, jwtConfig.SecretKey)

		// don't let gorm update the user
		session.User = models.User{}

		// update the session
		session.RefreshTokenId = sqlx.NewString(jwtClaimsDataAccess.GetIdentity())
		if err = sessionRepository.SafeUpdate(session, "uuid = ?", session.UUID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to update session.", nil)
		}

		return extras.NewMessageBodyOk(ctx, "Successfully refresh token.", &nokocore.MapAny{
			"accessToken": jwtToken,
		})
	}
}

func UserController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.GET("/profile", ProfileHandler(DB))
	group.GET("/session", SessionHandler(DB))
	group.POST("/logout", LogoutHandler(DB))
	group.GET("/refresh-token", RefreshTokenHandler(DB))

	return group
}
