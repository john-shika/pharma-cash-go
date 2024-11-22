package extras

import (
	"github.com/labstack/echo/v4"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
)

type JwtAuthInfo struct {
	Token               string                           `mapstructure:"token" json:"token" yaml:"token"`
	JwtToken            nokocore.JwtTokenImpl            `mapstructure:"jwt_token" json:"jwtToken" yaml:"jwt_token"`
	JwtClaims           nokocore.JwtClaimsImpl           `mapstructure:"jwt_claims" json:"jwtClaims" yaml:"jwt_claims"`
	JwtClaimsDataAccess nokocore.JwtClaimsDataAccessImpl `mapstructure:"jwt_claims_data_access" json:"jwtClaimsDataAccess" yaml:"jwt_claims_data_access"`
	Session             *models.Session                  `mapstructure:"session" json:"session" yaml:"session"`
}

func GetJwtAuthInfoFromEchoContext(ctx echo.Context) *JwtAuthInfo {
	var err error
	nokocore.KeepVoid(err)

	jwtAuthInfo := new(JwtAuthInfo)
	options := nokocore.NewForEachStructFieldsOptions()
	err = nokocore.ForEachStructFieldsReflect(jwtAuthInfo, options, func(name string, sFieldX nokocore.StructFieldExImpl) error {
		val := nokocore.GetValueReflect(ctx.Get(name))
		sFieldX.Set(val)
		return nil
	})

	return jwtAuthInfo
}
