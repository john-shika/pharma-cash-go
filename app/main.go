package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/middlewares"
	"nokowebapi/apis/models"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"pharma-cash-go/app/controllers"
	m "pharma-cash-go/app/models"
	"pharma-cash-go/app/repositories"
	"time"
)

func Main(args []string) nokocore.ExitCode {
	var err error
	var DB *gorm.DB
	nokocore.KeepVoid(DB, err, args)

	e := echo.New()
	e.Use(middlewares.Recovery())

	/// Echo Configs Start

	e.Validator = nokocore.NewValidator()
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

	/// Echo Configs End

	config := &gorm.Config{
		Logger: console.GetLogger("App").GORMLogger(),
	}

	sqliteFilePath := "migrations/dev.sqlite3"
	nokocore.NoErr(nokocore.CreateEmptyFile(sqliteFilePath))
	if DB, err = gorm.Open(sqlite.Open(sqliteFilePath), config); err != nil {
		panic("failed to connect database")
	}

	tables := []any{

		// cores
		&models.User{},
		&models.Session{},

		// app locals
		&m.Product{},
	}

	if err = DB.AutoMigrate(tables...); err != nil {
		console.Fatal(fmt.Sprintf("failed to migrate database: %s\n", err.Error()))
	}

	/// dummy data

	users := []models.User{
		{
			Username: "admin",
			Password: "Admin@1234",
			FullName: sqlx.NewString("John, Doe"),
			Email:    sqlx.NewString("admin@example.com"),
			Phone:    sqlx.NewString("+62 812-3456-7890"),
			Admin:    true,
			Level:    1,
		},
		{
			Username: "user",
			Password: "User@1234",
			FullName: sqlx.NewString("Angeline, Rose"),
			Email:    sqlx.NewString("user@example.com"),
			Phone:    sqlx.NewString("+62 823-4567-8901"),
			Admin:    false,
			Level:    1,
		},
	}

	userRepository := repositories.NewUserRepository(DB)

	var check *models.User
	for i, user := range users {
		nokocore.KeepVoid(i)

		if check, err = userRepository.First("username = ?", user.Username); err != nil {
			console.Warn(err.Error())
			continue
		}

		if check != nil {
			console.Warn(fmt.Sprintf("user '%s' already exists", user.Username))
			continue
		}

		if err = userRepository.Create(&user); err != nil {
			console.Warn(err.Error())
			continue
		}

		console.Warn(fmt.Sprintf("user '%s' has been created", user.Username))
	}

	/// dummy data

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
			"Accept-Language",
			"Content-Language",
			"Content-Length",
			"Content-Type",
		},
		Credentials: true,
		MaxAge:      86400,
	}))

	g := e.Group("/api/v1")

	gAuth := g.Group("/auth")
	gAuth.Use(middlewares.JWTAuth(DB))

	/// Controllers Start

	controllers.AnonymousController(g, DB)
	controllers.UserController(gAuth, DB)
	controllers.AdminController(gAuth, DB)

	/// Controllers End

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
