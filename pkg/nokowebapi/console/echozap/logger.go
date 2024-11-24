package echozap

import (
	"fmt"
	"nokowebapi/nokocore"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLoggerImpl interface {
	Sugar() *zap.SugaredLogger
	Named(s string) *zap.Logger
	WithOptions(opts ...zap.Option) *zap.Logger
	With(fields ...zap.Field) *zap.Logger
	WithLazy(fields ...zap.Field) *zap.Logger
	Level() zapcore.Level
	Check(level zapcore.Level, msg string) *zapcore.CheckedEntry
	Log(level zapcore.Level, msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Sync() error
	Core() zapcore.Core
	Name() string
}

// New is a middleware and zap to provide an "access log" like logging for each request.
func New(log ZapLoggerImpl) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error
			var id string
			nokocore.KeepVoid(err, id)

			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			host := req.Host
			latency := time.Since(start).String()
			httpStatusCode := nokocore.HttpStatusCode(res.Status)
			statusCode := httpStatusCode.ToInt()
			status := httpStatusCode.ToString()
			request := fmt.Sprintf("%s %s", req.Method, req.RequestURI)
			size := nokocore.ToFileSizeFormat(res.Size)
			ipAddr := c.RealIP()
			userAgent := req.UserAgent()

			message := fmt.Sprintf("host = %s, latency = %s, statusCode = %d, status = %s, request = %s, size = %s, ip_addr = %s, user_agent = %s", host, latency, statusCode, status, request, size, ipAddr, userAgent)

			fields := []zapcore.Field{
				zap.String("host", host),
				zap.String("latency", latency),
				zap.Int("statusCode", statusCode),
				zap.String("status", status),
				zap.String("request", request),
				zap.String("size", size),
				zap.String("ip_addr", ipAddr),
				zap.String("user_agent", userAgent),
			}

			if id = req.Header.Get(echo.HeaderXRequestID); id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			if id != "" {
				fields = append(fields, zap.String("request_id", id))
			}

			if statusCode >= 400 && err != nil {
				fields = append(fields, zap.Error(err))
			}

			switch {
			case statusCode >= 500:
				log.Error(fmt.Sprintf("Server Error, %s", message), fields...)
				break

			case statusCode >= 400:
				log.Warn(fmt.Sprintf("Client Error, %s", message), fields...)
				break

			case statusCode >= 300:
				log.Info(fmt.Sprintf("Redirection, %s", message), fields...)
				break

			case statusCode >= 200:
				log.Info(fmt.Sprintf("Success, %s", message), fields...)
				break

			default:
				log.Info(fmt.Sprintf("Information, %s", message), fields...)
			}

			return nil
		}
	}
}
