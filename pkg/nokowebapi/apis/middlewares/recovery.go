package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"nokowebapi/apis/extras"
	"nokowebapi/console"
	"nokowebapi/nokocore"
)

func Recovery() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			// don't panic
			defer nokocore.HandlePanic(func(err error) {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))

				// send internal server error
				nokocore.NoErr(extras.NewMessageBodyInternalServerError(ctx, "Internal server error.", nil))
			})

			return next(ctx)
		}
	}
}
