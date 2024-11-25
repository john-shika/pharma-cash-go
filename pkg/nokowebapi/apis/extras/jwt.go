package extras

import (
	"github.com/labstack/echo/v4"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
	"strings"
)

type JwtAuthInfo struct {
	Token               string                           `mapstructure:"token" json:"token" yaml:"token"`
	JwtToken            nokocore.JwtTokenImpl            `mapstructure:"jwt_token" json:"jwtToken" yaml:"jwt_token"`
	JwtClaims           nokocore.JwtClaimsImpl           `mapstructure:"jwt_claims" json:"jwtClaims" yaml:"jwt_claims"`
	JwtClaimsDataAccess nokocore.JwtClaimsDataAccessImpl `mapstructure:"jwt_claims_data_access" json:"jwtClaimsDataAccess" yaml:"jwt_claims_data_access"`
	Session             *models.Session                  `mapstructure:"session" json:"session" yaml:"session"`
	User                *models.User                     `mapstructure:"user" json:"user" yaml:"user"`
	Roles               []models.Role                    `mapstructure:"roles" json:"roles" yaml:"roles"`
}

func (j *JwtAuthInfo) GetRoles() []string {
	var temp []string
	for i, role := range j.Roles {
		nokocore.KeepVoid(i)
		temp = append(temp, role.RoleName)
	}
	return temp
}

func GetJwtTokenFromEchoContext(ctx echo.Context) (string, error) {
	var ok bool
	var token string
	nokocore.KeepVoid(ok, token)

	req := ctx.Request()
	authorization := req.Header.Get("Authorization")
	if token = strings.Trim(authorization, " "); token == "" {
		return "", nokocore.ErrJwtTokenInvalid
	}

	if token, ok = strings.CutPrefix(token, "Bearer "); !ok {
		return "", nokocore.ErrJwtTokenInvalid
	}
	if token = strings.Trim(token, " "); token == "" {
		return "", nokocore.ErrJwtTokenInvalid
	}
	return token, nil
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
