package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
)

// BaseControllerImpl interface for a defined group as extras.EchoGroupImpl and database with gorm.DB
type BaseControllerImpl interface {
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
	group extras.EchoGroupImpl
	db    *gorm.DB
}

func NewBaseController(group extras.EchoGroupImpl, db *gorm.DB) BaseControllerImpl {
	return &BaseController{
		group: group,
		db:    db,
	}
}

func (c *BaseController) DB() *gorm.DB {
	return c.db
}

func (c *BaseController) Use(middleware ...echo.MiddlewareFunc) {
	c.group.Use(middleware...)
}

func (c *BaseController) Connect(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.group.CONNECT(path, handler, middlewares...)
}

func (c *BaseController) Head(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.group.HEAD(path, handler, middlewares...)
}

func (c *BaseController) Options(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.group.OPTIONS(path, handler, middlewares...)
}

func (c *BaseController) Get(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.group.GET(path, handler, middlewares...)
}

func (c *BaseController) Post(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.group.POST(path, handler, middlewares...)
}

func (c *BaseController) Put(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.group.PUT(path, handler, middlewares...)
}

func (c *BaseController) Delete(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.group.DELETE(path, handler, middlewares...)
}

func (c *BaseController) Patch(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.group.PATCH(path, handler, middlewares...)
}

func (c *BaseController) Any(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.group.Any(path, handler, middlewares...)
}

func (c *BaseController) Group(prefix string, middleware ...echo.MiddlewareFunc) *echo.Group {
	return c.group.Group(prefix, middleware...)
}

func (c *BaseController) RouteNotFound(path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	c.group.RouteNotFound(path, handler, middlewares...)
}
