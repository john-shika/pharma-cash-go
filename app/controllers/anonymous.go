package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/apis/schemas"
	"nokowebapi/apis/validators"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/repositories"
)

func MessageHandler(userRepository repositories.UserRepositoryImpl) echo.HandlerFunc {
	nokocore.KeepVoid(userRepository)

	return func(ctx echo.Context) error {
		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", nokocore.MapAny{
			"message": "Hay!",
		})
	}
}

func LoginHandler(userRepository repositories.UserRepositoryImpl, sessionRepository repositories.SessionRepositoryImpl) echo.HandlerFunc {
	nokocore.KeepVoid(userRepository)

	jwtConfig := globals.GetJwtConfig()
	signingMethod := jwtConfig.GetSigningMethod()
	expiresIn := jwtConfig.GetExpiresIn()

	return func(ctx echo.Context) error {
		var err error
		var user *models.User
		nokocore.KeepVoid(err, user)

		req := ctx.Request()
		ipAddr := ctx.RealIP()
		userAgent := req.UserAgent()

		userBody := new(schemas.UserBody)
		if err = ctx.Bind(userBody); err != nil {
			return err
		}

		if err = ctx.Validate(userBody); err != nil {
			return err
		}

		if err = validators.ValidatePass(userBody.Password); err != nil {
			return err
		}

		if user, err = userRepository.Find("username = ?", userBody.Username); err != nil {
			return extras.NewMessageBodyUnauthorized(ctx, err.Error(), nil)
		}

		password := nokocore.NewPassword(userBody.Password)
		if !password.Equals(user.Password) {
			return extras.NewMessageBodyUnauthorized(ctx, "Invalid password.", nil)
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
		jwtClaimsDataAccess.SetRole(user.Role)
		jwtClaimsDataAccess.SetAdmin(user.Admin)
		jwtClaimsDataAccess.SetLevel(user.Level)

		jwtClaims := nokocore.CvtJwtClaimsDataAccessToJwtClaims(jwtClaimsDataAccess, signingMethod)
		jwtToken := nokocore.GenerateJwtToken(jwtClaims, jwtConfig.SecretKey)

		err = sessionRepository.Create(&models.Session{
			UserID:         user.ID,
			TokenId:        jwtClaimsDataAccess.GetIdentity(),
			RefreshTokenId: sql.NullString{},
			IPAddress:      ipAddr,
			UserAgent:      userAgent,
			Expires:        expires,
		})

		if err != nil {
			return extras.NewMessageBodyInternalServerError(ctx, err.Error(), nil)
		}

		return extras.NewMessageBodyOk(ctx, "Successfully logged in.", nokocore.MapAny{
			"token": jwtToken,
		})
	}
}

func AnonymousController(group *echo.Group, DB *gorm.DB) *echo.Group {

	userRepository := repositories.NewUserRepository(DB)
	sessionRepository := repositories.NewSessionRepository(DB)

	group.GET("/message", MessageHandler(userRepository))
	group.POST("/login", LoginHandler(userRepository, sessionRepository))

	return group
}
