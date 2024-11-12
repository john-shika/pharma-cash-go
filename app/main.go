package app

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
	"nokowebapi/apis/middlewares"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"pharma-cash-go/app/controllers"
	"time"
)

func Main(args []string) nokocore.ExitCode {

	e := echo.New()
	//e.HideBanner = true
	//e.HidePort = false

	// Alt-Svc: h3-25=":443"; ma=3600, h2=":443"; ma=3600
	// "Accept", "Accept-Language", "Content-Language", "Content-Type"
	// "Authorization", "Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma"

	e.Use(middlewares.CORSWithConfig(&middlewares.CORSConfig{
		Origins:     []string{"*"},
		Methods:     []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"},
		Headers:     []string{"Accept", "Accept-Language", "Content-Language", "Content-Type"},
		Credentials: true,
		MaxAge:      86400,
	}))

	g := e.Group("/api/v1")

	controllers.AnonymousController(g)

	h2s := &http2.Server{
		MaxConcurrentStreams: 100,
		MaxReadFrameSize:     16384,
		IdleTimeout:          10 * time.Second,
	}

	if taskConfig := globals.GetTasks().GetTask("self"); taskConfig != nil {
		if taskConfig.Network != nil {
			host := taskConfig.Network.GetHost()
			nokocore.NoErr(e.StartH2CServer(host, h2s))
			return nokocore.ExitCodeSuccess
		}
	}

	console.Error("failed to start server.")
	return nokocore.ExitCodeFailure
}
