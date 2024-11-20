package controllers

import (
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

	jwtConfig := globals.GetJwtConfig()
	signingMethod := jwtConfig.GetSigningMethod()
	expiresIn := jwtConfig.GetExpiresIn()

	return func(ctx echo.Context) error {
		var err error
		var user *models.User
		nokocore.KeepVoid(err, user)

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

		timeUtcNow := nokocore.GetTimeUtcNow()
		expires := timeUtcNow.Add(expiresIn)

		jwtClaimsDataAccess := nokocore.NewEmptyJwtClaimsDataAccess()
		jwtClaimsDataAccess.SetSubject("NokoWebApiToken")
		jwtClaimsDataAccess.SetIssuer(jwtConfig.Issuer)
		jwtClaimsDataAccess.SetAudience(jwtConfig.Audience)
		jwtClaimsDataAccess.SetIssuedAt(timeUtcNow)
		jwtClaimsDataAccess.SetExpiresAt(expires)
		jwtClaimsDataAccess.SetUser(user.Username)
		jwtClaimsDataAccess.SetSessionId(nokocore.NewUUID().String())
		jwtClaimsDataAccess.SetRole(user.Role)
		jwtClaimsDataAccess.SetAdmin(user.Admin)
		jwtClaimsDataAccess.SetLevel(user.Level)

		jwtClaims := nokocore.CvtJwtClaimsDataAccessToJwtClaims(jwtClaimsDataAccess, signingMethod)

		return extras.NewMessageBodyOk(ctx, "Successfully logged in.", nokocore.MapAny{
			"token": nokocore.GenerateJwtToken(jwtClaims, jwtConfig.SecretKey),
		})
	}
}

func AnonymousController(group *echo.Group, DB *gorm.DB) *echo.Group {

	userRepository := repositories.NewUserRepository(DB)

	group.GET("/message", MessageHandler(userRepository))
	group.POST("/login", LoginHandler(userRepository))

	return group
}
