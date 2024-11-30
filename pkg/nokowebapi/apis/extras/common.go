package extras

import (
	"errors"
	"github.com/labstack/echo/v4"
	"io/fs"
	"net/http"
	"nokowebapi/apis/schemas"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
)

type EchoGroupImpl interface {
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
	File(path string, file string)
	RouteNotFound(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Add(method string, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
	Static(pathPrefix string, fsRoot string)
	StaticFS(pathPrefix string, filesystem fs.FS)
	FileFS(path string, file string, filesystem fs.FS, m ...echo.MiddlewareFunc) *echo.Route
}

func EchoHTTPErrorHandler() echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		req := ctx.Request()

		if req.Method == echo.HEAD {
			err = ctx.NoContent(http.StatusInternalServerError)
			return
		}

		var httpError *echo.HTTPError
		if errors.As(err, &httpError) {
			message := nokocore.ToStringReflect(httpError.Message)
			httpStatusCode := nokocore.HttpStatusCode(httpError.Code)
			httpStatusCodeValue := httpStatusCode.ToString()
			messageBody := schemas.NewMessageBody(false, httpError.Code, string(httpStatusCodeValue), message, nil)
			err = ctx.JSON(httpError.Code, messageBody)
			return
		}

		var validationError *sqlx.ValidateError
		if errors.As(err, &validationError) {
			message := "Validation failed."
			fields := validationError.Fields()
			err = NewMessageBodyUnprocessableEntity(ctx, message, &nokocore.MapAny{
				"fields": fields,
			})
			return
		}

		err = NewMessageBodyInternalServerError(ctx, err.Error(), nil)
	}
}

type URLQueryPagination struct {
	Size   int
	Page   int
	Offset int
	Limit  int
}

func NewURLQueryPagination(size int, page int) *URLQueryPagination {
	var offset int
	var limit int
	nokocore.KeepVoid(offset, limit)

	switch {
	// FIXME: this call lots of resources
	case size < 0:
		size = -1
		page = 1
		offset = 0
		limit = -1
		break

	case size > 0 && page > 0:
		offset = (page - 1) * size
		limit = size
		break

	default:
		size = 0
		page = 1
		offset = 0
		limit = 0
		break
	}

	return &URLQueryPagination{
		Size:   size,
		Page:   page,
		Offset: offset,
		Limit:  limit,
	}
}

func NewURLQueryPaginationFromEchoContext(ctx echo.Context) *URLQueryPagination {
	size := ParseQueryToInt(ctx, "size")
	page := ParseQueryToInt(ctx, "page")

	return NewURLQueryPagination(size, page)
}
