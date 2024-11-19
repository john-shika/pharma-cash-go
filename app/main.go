package app

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/middlewares"
	"nokowebapi/apis/models"
	"nokowebapi/apis/schemas"
	"nokowebapi/apis/validators"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	"pharma-cash-go/app/controllers"
	"time"
)

func Main(args []string) nokocore.ExitCode {
	var err error
	var db *gorm.DB
	nokocore.KeepVoid(db, err, args)

	e := echo.New()

	/// Echo Configs Start

	e.Validator = validators.NewEchoValidator()

	e.HideBanner = false
	e.HidePort = false

	// http error handling
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		req := ctx.Request()

		if req.Method == echo.HEAD {
			err = ctx.NoContent(http.StatusInternalServerError)
			return
		}

		var httpError *echo.HTTPError
		if errors.As(err, &httpError) {
			message := nokocore.GetStringValueReflect(httpError.Message)
			httpStatusCodeValue := nokocore.GetValueFromHttpStatusCode(nokocore.HttpStatusCode(httpError.Code))
			messageBody := schemas.NewMessageBody(false, httpError.Code, string(httpStatusCodeValue), message, nil)
			err = ctx.JSON(httpError.Code, messageBody)
			return
		}

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			message := "Validation failed."
			var data []string

			fields := []validator.FieldError(validationErrors)
			for i, field := range fields {
				nokocore.KeepVoid(i)

				name := nokocore.ToSnakeCase(field.StructField())
				data = append(data, fmt.Sprintf("Field '%s' contains an invalid value.", name))
			}

			err = extras.NewMessageBodyUnprocessableEntity(ctx, message, data)
			return
		}

		var validatePassErr *validators.ValidatePassError
		if errors.As(err, &validatePassErr) {
			message := "Validation failed."
			data := validatePassErr.Fields()
			err = extras.NewMessageBodyUnprocessableEntity(ctx, message, data)
			return
		}

		err = extras.NewMessageBodyInternalServerError(ctx, err.Error(), nil)
	}

	/// Echo Configs End

	config := new(gorm.Config)
	sqliteFilePath := "migrations/dev.db"

	nokocore.NoErr(nokocore.EnsureDirAndFile(sqliteFilePath))
	if db, err = gorm.Open(sqlite.Open(sqliteFilePath), config); err != nil {
		panic("failed to connect database")
	}

	tables := []interface{}{
		&models.User{},
		&models.Session{},
	}

	if err = db.AutoMigrate(tables...); err != nil {
		console.Fatal(fmt.Sprintf("failed to migrate database: %s\n", err.Error()))
	}

	/// dummy data

	users := []*models.User{
		{
			Model: models.Model{
				UUID: nokocore.NewUUID(),
			},
			Username: "admin",
			Password: "Admin@1234",
			Email:    sqlx.NewString("admin@example.com"),
			Phone:    sqlx.NewString("081234567890"),
			Admin:    true,
			Level:    1,
		},
		{
			Model: models.Model{
				UUID: nokocore.NewUUID(),
			},
			Username: "user",
			Password: "User@1234",
			Email:    sqlx.NewString("user@example.com"),
			Phone:    sqlx.NewString("081234567890"),
			Admin:    false,
			Level:    1,
		},
	}

	for i, user := range users {
		nokocore.KeepVoid(i)

		check := new(models.User)
		if err = db.First(check, "username = ?", user.Username).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				console.Fatal(fmt.Sprintf("failed to check user existence: %s\n", err.Error()))
			}

			// create new user
			if err = db.Create(user).Error; err != nil {
				console.Fatal(fmt.Sprintf("failed to migrate database: %s\n", err.Error()))
			}

			// log
			console.Log(fmt.Sprintf("created user: %s\n", user.Username))
		}

		// skipping user
		console.Log(fmt.Sprintf("skipping user: %s\n", user.Username))
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
		Origins:     []string{"*"},
		Methods:     []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"},
		Headers:     []string{"Accept", "Accept-Language", "Content-Language", "Content-Type"},
		Credentials: true,
		MaxAge:      86400,
	}))

	g := e.Group("/api/v1")

	/// Controllers Start

	controllers.AnonymousController(g, db)

	/// Controllers End

	h2s := &http2.Server{
		MaxConcurrentStreams: 100,
		MaxReadFrameSize:     16384,
		IdleTimeout:          10 * time.Second,
	}

	tasks := globals.GetTasksConfig()
	if taskConfig := tasks.GetTaskConfig("self"); taskConfig != nil {
		if taskConfig.Network != nil {
			host := taskConfig.Network.GetHost()
			nokocore.NoErr(e.StartH2CServer(host, h2s))
			return nokocore.ExitCodeSuccess
		}
	}

	console.Error("failed to start server.")
	return nokocore.ExitCodeFailure
}
