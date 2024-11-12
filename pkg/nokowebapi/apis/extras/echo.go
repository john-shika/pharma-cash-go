package extras

import (
	"github.com/labstack/echo/v4"
	"nokowebapi/nokocore"
	"strings"
)

type EchoRouterImpl interface {
	Use(middleware ...echo.MiddlewareFunc)
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Any(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route
	Match(methods []string, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route
	Group(prefix string, middleware ...echo.MiddlewareFunc) (sg *echo.Group)
	RouteNotFound(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
}

func GetJwtTokenFromEchoContext(c echo.Context) (string, error) {
	var ok bool
	var token string
	nokocore.KeepVoid(ok, token)

	if token = strings.Trim(c.Request().Header.Get("Authorization"), " "); token == "" {
		return "", nokocore.ErrJwtTokenNotFound
	}

	if token, ok = strings.CutPrefix(token, "Bearer "); !ok {
		return "", nokocore.ErrJwtTokenNotFound
	}
	if token = strings.Trim(token, " "); token == "" {
		return "", nokocore.ErrJwtTokenNotFound
	}
	return token, nil
}
