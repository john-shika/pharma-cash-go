package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
)

// BaseControllerImpl interface for a defined router as extras.EchoRouterImpl and database with gorm.DB
type BaseControllerImpl interface {
	Router() extras.EchoRouterImpl
	DB() *gorm.DB
	Use(middleware ...echo.MiddlewareFunc)
	Connect(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc)
	Head(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc)
	Options(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc)
	Get(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc)
	Post(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc)
	Put(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc)
	Delete(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc)
	Patch(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc)
	Any(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc)
	Group(prefix string, middleware ...echo.MiddlewareFunc) *echo.Group
	RouteNotFound(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc)
}

type BaseController struct {
	router extras.EchoRouterImpl
	db     *gorm.DB
}

func NewBaseController(router extras.EchoRouterImpl, db *gorm.DB) BaseControllerImpl {
	return &BaseController{
		router: router,
		db:     db,
	}
}

func (c *BaseController) Router() extras.EchoRouterImpl {
	return c.router
}

func (c *BaseController) DB() *gorm.DB {
	return c.db
}

func (c *BaseController) Use(middleware ...echo.MiddlewareFunc) {
	c.router.Use(middleware...)
}

func (c *BaseController) Connect(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.router.CONNECT(path, handler, middlewares...)
}

func (c *BaseController) Head(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.router.HEAD(path, handler, middlewares...)
}

func (c *BaseController) Options(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.router.OPTIONS(path, handler, middlewares...)
}

func (c *BaseController) Get(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.router.GET(path, handler, middlewares...)
}

func (c *BaseController) Post(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.router.POST(path, handler, middlewares...)
}

func (c *BaseController) Put(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.router.PUT(path, handler, middlewares...)
}

func (c *BaseController) Delete(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.router.DELETE(path, handler, middlewares...)
}

func (c *BaseController) Patch(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.router.PATCH(path, handler, middlewares...)
}

func (c *BaseController) Any(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.router.Any(path, handler, middlewares...)
}

func (c *BaseController) Group(prefix string, middleware ...echo.MiddlewareFunc) *echo.Group {
	return c.router.Group(prefix, middleware...)
}

func (c *BaseController) RouteNotFound(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.router.RouteNotFound(path, handler, middlewares...)
}
