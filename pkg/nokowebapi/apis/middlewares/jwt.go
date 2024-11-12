package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/apis/schemas"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
)

func AuthJWT(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error
			var token string
			var jwtToken nokocore.JwtTokenImpl
			var session *models.Session
			nokocore.KeepVoid(err, token, jwtToken, session)

			if token, err = extras.GetJwtTokenFromEchoContext(c); err != nil {
				return c.JSON(http.StatusUnauthorized, schemas.NewMessageBodyUnauthorized(err.Error(), nil))
			}

			jwtConfig := globals.GetJwtConfig()
			if jwtToken, err = nokocore.ParseJwtToken(token, jwtConfig.SecretKey, jwt.SigningMethodHS256); err != nil {
				return c.JSON(http.StatusUnauthorized, schemas.NewMessageBodyUnauthorized(err.Error(), nil))
			}

			jwtClaims := nokocore.Unwrap(nokocore.GetJwtClaimsFromJwtToken(jwtToken, jwt.SigningMethodHS256))
			jwtClaimsAccessData := nokocore.CvtJwtClaimsToJwtClaimsAccessData(jwtClaims)

			fmt.Println(jwtClaimsAccessData)

			c.Set("token", token)
			c.Set("jwt_token", jwtToken)
			c.Set("jwt_claims", jwtClaims)
			c.Set("jwt_claims_access_data", jwtClaimsAccessData)
			return next(c)
		}
	}
}
