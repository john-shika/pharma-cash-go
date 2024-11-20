package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
)

func JWTAuth(DB *gorm.DB) echo.MiddlewareFunc {
	nokocore.KeepVoid(DB)

	jwtConfig := globals.GetJwtConfig()

	//sessionRepository := repositories.NewSessionRepository(DB)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			var err error
			var token string
			var jwtToken nokocore.JwtTokenImpl
			var session *models.Session
			nokocore.KeepVoid(err, token, jwtToken, session)

			// don't panic
			defer nokocore.HandlePanic(func(err error) {
				nokocore.NoErr(extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt recovered.", nil))
				nokocore.KeepVoid(err)
			})

			if token, err = extras.GetJwtTokenFromEchoContext(ctx); err != nil {
				return extras.NewMessageBodyUnauthorized(ctx, "Unable to get JWT token.", nil)
			}

			if jwtToken, err = nokocore.ParseJwtToken(token, jwtConfig.SecretKey, jwtConfig.GetSigningMethod()); err != nil {
				return extras.NewMessageBodyUnauthorized(ctx, "Invalid JWT token.", nil)
			}

			jwtClaims := nokocore.Unwrap(nokocore.GetJwtClaimsFromJwtToken(jwtToken))
			jwtClaimsDataAccess, jwtSigningMethod := nokocore.CvtJwtClaimsToJwtClaimsDataAccess(jwtClaims)

			nokocore.KeepVoid(jwtClaimsDataAccess, jwtSigningMethod)
			fmt.Println(jwtClaimsDataAccess)

			sessionId := jwtClaimsDataAccess.GetSessionId()
			identity := jwtClaimsDataAccess.GetIdentity()

			session = new(models.Session)
			tx := DB.Preload("User").Where("uuid = ? AND (token_id = ? OR refresh_token_id = ?)", sessionId, identity, identity).Find(session)
			if err = tx.Error; err != nil {
				console.Error(err.Error())

				return extras.NewMessageBodyUnauthorized(ctx, "Invalid JWT token.", nil)
			}

			if tx.RowsAffected == 0 {
				console.Error("session not found")
				return extras.NewMessageBodyUnauthorized(ctx, "Invalid JWT token.", nil)
			}

			//if session, err = sessionRepository.SafeFirst("uuid = ? AND (token_id = ? OR refresh_token_id = ?)", sessionId, identity, identity); err != nil {
			//	console.Error(err.Error())
			//
			//	return extras.NewMessageBodyUnauthorized(ctx, "Invalid JWT token.", nil)
			//}

			if session != nil {
				ctx.Set("token", token)
				ctx.Set("jwt_token", jwtToken)
				ctx.Set("jwt_claims", jwtClaims)
				ctx.Set("jwt_claims_data_access", jwtClaimsDataAccess)
				ctx.Set("session", session)

				return next(ctx)
			}

			console.Error("session not found")
			return extras.NewMessageBodyUnauthorized(ctx, "Invalid JWT token.", nil)
		}
	}
}
