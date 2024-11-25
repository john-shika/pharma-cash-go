package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
	"gorm.io/gorm"
	"net/http"
	"nokowebapi/apis"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/middlewares"
	"nokowebapi/console"
	"nokowebapi/console/echozap"
	"nokowebapi/console/zapgorm"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"time"
)

func Main(args []string) nokocore.ExitCode {
	var err error
	var DB *gorm.DB
	nokocore.KeepVoid(DB, err, args)

	e := echo.New()
	e.Use(middlewares.Recovery())
	e.Use(echozap.New(console.GetLogger("Echo")))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			header := c.Response().Header()

			header.Set("Alt-Svc", "h3-25=\":443\"; ma=3600, h2=\":443\"; ma=3600")
			header.Set("Access-Control-Expose-Headers", "Date, Vary")
			header.Set("Date", time.Now().Format(time.RFC1123))
			header.Set("Server", "Apache/2.4.62 (Ubuntu)")
			header.Set("Vary", "Accept-Encoding")

			return next(c)
		}
	})
	e.Use(middlewares.CORSWithConfig(&middlewares.CORSConfig{
		Origins: []string{
			"*",
		},
		Methods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			//http.MethodPatch,
			http.MethodDelete,
			//http.MethodConnect,
			http.MethodOptions,
			//http.MethodTrace,
		},
		Headers: []string{
			"Accept",
			"Accept-Encoding",
			"Accept-Language",
			"Authorization",
			"Content-Language",
			"Content-Length",
			"Content-Type",
		},
		Credentials: true,
		MaxAge:      86400,
	}))

	/// Echo Configs Start

	e.Validator = sqlx.NewValidator()
	//e.IPExtractor = echo.ExtractIPDirect()
	//e.IPExtractor = echo.ExtractIPFromXFFHeader(
	//	echo.TrustLoopback(false),
	//	echo.TrustLinkLocal(false),
	//	echo.TrustPrivateNet(false),
	//	echo.TrustIPRange(lbIPRange),
	//)

	e.HideBanner = false
	e.HidePort = false

	// http error handling
	e.HTTPErrorHandler = extras.EchoHTTPErrorHandler()

	// create group rest api
	group := e.Group("/api/v1")

	/// Echo Configs End

	config := &gorm.Config{
		Logger: zapgorm.New(console.GetLogger("GORM")),
	}

	sqliteFilePath := "migrations/dev.sqlite3"
	if DB, err = apis.SqliteOpenFile(sqliteFilePath, config); err != nil {
		console.Error(fmt.Sprintf("panic: %s", err.Error()))
		return nokocore.ExitCodeFailure
	}

	// START DATABASE Auto Migrations

	DBAutoMigrations(DB)

	// END DATABASE Auto Migrations

	/// START FACTORIES

	Factories(DB)

	/// END FACTORIES

	// START CONTROLLERS

	Controllers(group, DB)

	// END CONTROLLERS

	h2s := &http2.Server{
		MaxConcurrentStreams: 100,
		MaxReadFrameSize:     16384,
		IdleTimeout:          10 * time.Second,
	}

	if taskConfig := globals.GetTaskConfig("self"); taskConfig != nil {
		if network := taskConfig.GetNetwork(); network != nil {
			host := network.GetHost()
			nokocore.NoErr(e.StartH2CServer(host, h2s))
			return nokocore.ExitCodeSuccess
		}
	}

	console.Error("failed to start server.")
	return nokocore.ExitCodeFailure
}
