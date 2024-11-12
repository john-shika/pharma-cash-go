package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
	"nokowebapi/apis/middlewares"
	"nokowebapi/nokocore"
	"os"
	"time"
)

func Main(args []string) nokocore.ExitCode {

	e := echo.New()
	//e.HideBanner = true
	//e.HidePort = false

	corsConfig := middlewares.CORSConfig{
		Origins:     []string{"*"},
		Methods:     []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"},
		Headers:     []string{"Accept", "Accept-Language", "Content-Language", "Content-Type"},
		Credentials: true,
		MaxAge:      86400,
	}

	// Alt-Svc: h3-25=":443"; ma=3600, h2=":443"; ma=3600
	// "Accept", "Accept-Language", "Content-Language", "Content-Type"
	// "Authorization", "Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma"

	e.Use(middlewares.CORSWithConfig(corsConfig))

	r := e.Router()

	r.Add("GET", "/", func(c echo.Context) error {
		return c.JSON(200, nokocore.MapAny{
			"message": "Hello, World!",
		})
	})

	h2s := &http2.Server{
		MaxConcurrentStreams: 100,
		MaxReadFrameSize:     16384,
		IdleTimeout:          10 * time.Second,
	}
	
	fmt.Println("Current working directory:", nokocore.Unwrap(os.Getwd()))

	nokocore.NoErr(e.StartH2CServer(":8080", h2s))

	return nokocore.ExitCodeSuccess
}
