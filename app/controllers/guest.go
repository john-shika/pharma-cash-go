package controllers

import (
	"database/sql"
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
	models2 "pharma-cash-go/app/models"
	repositories2 "pharma-cash-go/app/repositories"
	schemas2 "pharma-cash-go/app/schemas"
)

func MessageHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	return func(ctx echo.Context) error {
		return extras.NewMessageBodyOk(ctx, "Hai\x21", nil)
	}
}

func LoginHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	jwtConfig := globals.GetJwtConfig()
	signingMethod := jwtConfig.GetSigningMethod()
	expiresIn := jwtConfig.GetExpiresIn()

	userRepository := repositories.NewUserRepository(DB)
	sessionRepository := repositories.NewSessionRepository(DB)
	employeeRepository := repositories2.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var username string
		var user *models.User
		var employee *models2.Employee
		nokocore.KeepVoid(err, username, user, employee)

		req := ctx.Request()
		ipAddr := ctx.RealIP()
		userAgent := req.UserAgent()

		userBody := new(schemas.UserBody)
		if err = ctx.Bind(userBody); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Invalid request body.", nil)
		}

		if err = ctx.Validate(userBody); err != nil {
			return err
		}

		if user, err = userRepository.SafeLogin(userBody.Username, userBody.Password); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnauthorized(ctx, "Invalid username or password.", nil)
		}

		// get user roles
		roles := nokocore.RolesUnpack(user.Roles)

		sessionId := nokocore.NewUUID()
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

		session := &models.Session{
			UserID:         user.ID,
			TokenId:        jwtClaimsDataAccess.GetIdentity(),
			RefreshTokenId: sql.NullString{},
			IPAddress:      ipAddr,
			UserAgent:      userAgent,
			Expires:        expires,
		}

		session.UUID = sessionId
		if err = sessionRepository.SafeCreate(session); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create session.", nil)
		}

		preloads := []string{"Shift"}
		if employee, err = employeeRepository.SafePreFirst(preloads, "user_id = ?", user.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get employee.", nil)
		}

		var shift schemas2.ShiftResult
		if employee != nil {
			shift = schemas2.ToShiftResult(&employee.Shift)
		}

		roles = nokocore.RolesUnpack(user.Roles)
		userResult := schemas.ToUserResult(user, nil)
		return extras.NewMessageBodyOk(ctx, "Successfully logged in.", &nokocore.MapAny{
			"accessToken": jwtToken,
			"user":        userResult,
			"shift":       shift,
		})
	}
}

func GuestController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.GET("/message", MessageHandler(DB))
	group.POST("/login", LoginHandler(DB))

	return group
}