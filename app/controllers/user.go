package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/apis/repositories"
	"nokowebapi/apis/schemas"
	"nokowebapi/apis/utils"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
	repositories2 "pharma-cash-go/app/repositories"
	schemas2 "pharma-cash-go/app/schemas"
)

func GetProfile(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	employeeRepository := repositories2.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var user *models.User
		var employee *models2.Employee
		nokocore.KeepVoid(err, employee)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)
		roles := jwtAuthInfo.GetRoles()
		user = jwtAuthInfo.User

		preloads := []string{"Shift"}
		if employee, err = employeeRepository.SafePreFirst(preloads, "user_id = ?", user.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get employee.", nil)
		}

		var shift schemas2.ShiftResult
		if employee != nil {
			shift = schemas2.ToShiftResult(&employee.Shift)
		}

		userResult := schemas.ToUserResult(user)
		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
			"user":  userResult,
			"shift": shift,
			"roles": roles,
		})
	}
}

func GetAllSessions(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	sessionRepository := repositories.NewSessionRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var sessions []models.Session
		nokocore.KeepVoid(err, sessions)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)
		session := jwtAuthInfo.Session
		user := jwtAuthInfo.User

		timeUtcNow := nokocore.GetTimeUtcNow()
		if sessions, err = sessionRepository.SafeMany(0, -1, "user_id = ? AND expires > ?", user.ID, timeUtcNow); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get sessions.", nil)
		}

		sessionID := session.ID

		size := len(sessions)
		sessionResults := make([]schemas.SessionResult, size)
		for i, session := range sessions {
			nokocore.KeepVoid(i)

			session.User = *user
			sessionResult := schemas.ToSessionResult(&session)
			sessionResult.Used = session.ID == sessionID
			sessionResults[i] = sessionResult
		}

		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
			"sessions": sessionResults,
		})
	}
}

func SetLogout(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	sessionRepository := repositories.NewSessionRepository(DB)

	return func(ctx echo.Context) error {
		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)
		session := jwtAuthInfo.Session

		sessionID := session.UUID
		if err := sessionRepository.SafeDelete(session, "uuid = ?", sessionID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to log out.", nil)
		}

		return extras.NewMessageBodyOk(ctx, "Successfully logged out.", nil)
	}
}

func GetRefreshToken(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	jwtConfig := globals.GetJwtConfig()
	signingMethod := jwtConfig.GetSigningMethod()
	expiresIn := jwtConfig.GetExpiresIn()

	sessionRepository := repositories.NewSessionRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		nokocore.KeepVoid(err)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)
		session := jwtAuthInfo.Session
		user := jwtAuthInfo.User

		// get user roles
		roles := utils.ToUserRolesArrayString(user.Roles)

		sessionID := session.UUID
		timeUtcNow := nokocore.GetTimeUtcNow()
		expires := timeUtcNow.Add(expiresIn)

		jwtClaimsDataAccess := nokocore.NewEmptyJwtClaimsDataAccess()
		jwtClaimsDataAccess.SetSubject("NokoWebApiToken")
		jwtClaimsDataAccess.SetIssuer(jwtConfig.Issuer)
		jwtClaimsDataAccess.SetAudience(jwtConfig.Audience)
		jwtClaimsDataAccess.SetIssuedAt(timeUtcNow)
		jwtClaimsDataAccess.SetExpiresAt(expires)
		jwtClaimsDataAccess.SetUser(user.Username)
		jwtClaimsDataAccess.SetSessionID(sessionID.String())
		jwtClaimsDataAccess.SetRoles(roles)
		jwtClaimsDataAccess.SetAdmin(user.Admin)
		jwtClaimsDataAccess.SetLevel(user.Level)

		jwtClaims := nokocore.ToJwtClaims(jwtClaimsDataAccess, signingMethod)
		jwtToken := nokocore.GenerateJwtToken(jwtClaims, jwtConfig.SecretKey)

		// don't let gorm update the user
		session.User = models.User{}

		// update the session
		session.RefreshTokenID = sqlx.NewString(jwtClaimsDataAccess.GetIdentity())
		if err = sessionRepository.SafeUpdate(session, "uuid = ?", session.UUID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Failed to update session.", nil)
		}

		return extras.NewMessageBodyOk(ctx, "Successfully refresh token.", &nokocore.MapAny{
			"accessToken": jwtToken,
		})
	}
}

func DeleteOwnUser(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	userRepository := repositories.NewUserRepository(DB)
	employeeRepository := repositories2.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var employee *models2.Employee
		nokocore.KeepVoid(err, employee)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)
		user := jwtAuthInfo.User

		forced := extras.ParseQueryToBool(ctx, "forced")

		if !forced && user.DeletedAt.Valid {
			return extras.NewMessageBodyOk(ctx, "User already deleted.", nil)
		}

		// check di employee juga
		if employee, err = employeeRepository.SafeFirst("user_id = ?", user.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get employee.", nil)
		}

		if employee != nil {
			if forced {
				if err = employeeRepository.Delete(employee, "user_id = ?", user.ID); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to delete employee.", nil)
				}

			} else {
				if err = employeeRepository.SafeDelete(employee, "user_id = ?", user.ID); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to delete employee.", nil)
				}
			}
		}

		if forced {
			if err := userRepository.Delete(user, "id = ?", user.ID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to delete user.", nil)
			}

		} else {
			if err := userRepository.SafeDelete(user, "id = ?", user.ID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to delete user.", nil)
			}
		}

		return extras.NewMessageBodyOk(ctx, "Successfully deleted.", &nokocore.MapAny{
			"user": schemas.ToUserResult(user),
		})
	}
}

func UserController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.GET("/profile", GetProfile(DB))
	group.GET("/sessions", GetAllSessions(DB))
	group.POST("/logout", SetLogout(DB))
	group.GET("/refresh-token", GetRefreshToken(DB))
	group.DELETE("/me", DeleteOwnUser(DB))

	return group
}
