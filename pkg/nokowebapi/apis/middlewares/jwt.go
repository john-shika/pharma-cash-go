package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
)

func JWTAuth(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			var err error
			var token string
			var jwtToken nokocore.JwtTokenImpl
			var session *models.Session
			nokocore.KeepVoid(err, token, jwtToken, session)

			// don't panic
			defer nokocore.HandlePanic(func(err error) {
				nokocore.NoErr(extras.NewMessageBodyUnauthorized(ctx, "Recovery.", nil))
				nokocore.KeepVoid(err)
			})

			if token, err = extras.GetJwtTokenFromEchoContext(ctx); err != nil {
				return extras.NewMessageBodyUnauthorized(ctx, err.Error(), nil)
			}

			jwtConfig := globals.GetJwtConfig()
			if jwtToken, err = nokocore.ParseJwtToken(token, jwtConfig.SecretKey, jwtConfig.GetSigningMethod()); err != nil {
				return extras.NewMessageBodyUnauthorized(ctx, err.Error(), nil)
			}

			jwtClaims := nokocore.Unwrap(nokocore.GetJwtClaimsFromJwtToken(jwtToken))
			jwtClaimsAccessData, jwtSigningMethod := nokocore.CvtJwtClaimsToJwtClaimsAccessData(jwtClaims)

			nokocore.KeepVoid(jwtClaimsAccessData, jwtSigningMethod)
			fmt.Println(jwtClaimsAccessData)

			ctx.Set("token", token)
			ctx.Set("jwt_token", jwtToken)
			ctx.Set("jwt_claims", jwtClaims)
			ctx.Set("jwt_claims_access_data", jwtClaimsAccessData)
			return next(ctx)
		}
	}
}
