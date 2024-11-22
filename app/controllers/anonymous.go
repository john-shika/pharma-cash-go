package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/apis/schemas"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/repositories"
)

func MessageHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	return func(ctx echo.Context) error {
		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", nokocore.MapAny{
			"message": "Hay!",
		})
	}
}

func LoginHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	jwtConfig := globals.GetJwtConfig()
	signingMethod := jwtConfig.GetSigningMethod()
	expiresIn := jwtConfig.GetExpiresIn()

	userRepository := repositories.NewUserRepository(DB)
	sessionRepository := repositories.NewSessionRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var user *models.User
		var username string
		nokocore.KeepVoid(err, user, username)

		req := ctx.Request()
		ipAddr := ctx.RealIP()
		userAgent := req.UserAgent()

		userBody := new(schemas.UserBody)
		if err = ctx.Bind(userBody); err != nil {
			console.Error(err.Error())

			return extras.NewMessageBodyInternalServerError(ctx, "Invalid request body.", nil)
		}

		if err = ctx.Validate(userBody); err != nil {
			return err
		}

		if err = nokocore.ValidatePassword(userBody.Password); err != nil {
			return err
		}

		if user, err = userRepository.SafeLogin(userBody.Username, userBody.Password); err != nil {
			console.Error(err.Error())

			return extras.NewMessageBodyUnauthorized(ctx, "Invalid username or password.", nil)
		}

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
		jwtClaimsDataAccess.SetRoles(user.GetRoles())
		jwtClaimsDataAccess.SetAdmin(user.Admin)
		jwtClaimsDataAccess.SetLevel(user.Level)

		jwtClaims := nokocore.CvtJwtClaimsDataAccessToJwtClaims(jwtClaimsDataAccess, signingMethod)
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
			console.Error(err.Error())

			return extras.NewMessageBodyInternalServerError(ctx, "Failed to create session.", nil)
		}

		return extras.NewMessageBodyOk(ctx, "Successfully logged in.", nokocore.MapAny{
			"token": jwtToken,
			"user": nokocore.MapAny{
				"fullname": user.FullName.String,
				"username": user.Username,
				"email":    user.Email.String,
				"phone":    user.Phone.String,
				"admin":    user.Admin,
				"roles":    user.GetRoles(),
				"level":    user.Level,
			},
		})
	}
}

func AnonymousController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.GET("/message", MessageHandler(DB))
	group.POST("/login", LoginHandler(DB))

	return group
}
