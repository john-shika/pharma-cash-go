package middlewares

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/repositories"
)

func JWTAuth(DB *gorm.DB) echo.MiddlewareFunc {
	nokocore.KeepVoid(DB)

	jwtConfig := globals.GetJwtConfig()
	signingMethod := jwtConfig.GetSigningMethod()

	sessionRepository := repositories.NewSessionRepository(DB)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			var err error
			var token string
			var jwtToken nokocore.JwtTokenImpl
			var session *models.Session
			nokocore.KeepVoid(err, token, jwtToken, session)

			if token, err = extras.GetJwtTokenFromEchoContext(ctx); err != nil {
				return extras.NewMessageBodyUnauthorized(ctx, "Unable to get JWT token.", nil)
			}

			if jwtToken, err = nokocore.ParseJwtToken(token, jwtConfig.SecretKey, jwtConfig.GetSigningMethod()); err != nil {
				return extras.NewMessageBodyUnauthorized(ctx, "Invalid JWT token.", nil)
			}

			jwtClaims := nokocore.Unwrap(nokocore.GetJwtClaimsFromJwtToken(jwtToken))
			jwtClaimsDataAccess, jwtSigningMethod := nokocore.ToJwtClaimsDataAccess(jwtClaims)

			if jwtSigningMethod.Alg() != signingMethod.Alg() {
				return extras.NewMessageBodyUnauthorized(ctx, fmt.Sprintf("Invalid JWT token. Expected signing method: %s. Actual signing method: %s.", signingMethod.Alg(), jwtSigningMethod.Alg()), nil)
			}

			sessionId := jwtClaimsDataAccess.GetSessionId()
			identity := jwtClaimsDataAccess.GetIdentity()

			// initial session
			session = new(models.Session)

			// get current session
			preloads := []string{"User"}
			if session, err = sessionRepository.SafePreFirst(preloads, "uuid = ? AND (token_id = ? OR refresh_token_id = ?)", sessionId, identity, identity); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))

				return extras.NewMessageBodyUnauthorized(ctx, "Invalid JWT token.", nil)
			}

			if session != nil {

				// update refresh token id
				if session.RefreshTokenId.String == identity {
					session.TokenId = identity
					session.RefreshTokenId = sql.NullString{}
					if err = sessionRepository.SafeUpdate(session, "uuid = ?", sessionId); err != nil {
						console.Error(fmt.Sprintf("panic: %s", err.Error()))
					}
				}

				// get user data
				user := &session.User

				// set echo context
				ctx.Set("token", token)
				ctx.Set("jwt_token", jwtToken)
				ctx.Set("jwt_claims", jwtClaims)
				ctx.Set("jwt_claims_data_access", jwtClaimsDataAccess)
				ctx.Set("session", session)
				ctx.Set("user", user)

				return next(ctx)
			}

			console.Error("session not found")
			return extras.NewMessageBodyUnauthorized(ctx, "Invalid JWT token.", nil)
		}
	}
}
