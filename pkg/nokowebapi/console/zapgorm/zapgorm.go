package zapgorm

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"nokowebapi/nokocore"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	gormPackage    = filepath.Join("gorm.io", "gorm")
	zapgormPackage = filepath.Join("nokowebapi", "nokocore/console/zapgorm")
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

type FC func() (sql string, rowsAffected int64)
type Handler func(ctx context.Context) []zapcore.Field

type LoggerImpl interface {
	SetAsDefault()
	LogMode(level logger.LogLevel) logger.Interface
	Info(ctx context.Context, str string, args ...any)
	Warn(ctx context.Context, str string, args ...any)
	Error(ctx context.Context, str string, args ...any)
	Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
}

type Logger struct {
	ZapLogger                 ZapLoggerImpl
	LogLevel                  logger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
	Handler                   Handler
}

func New(zapLogger ZapLoggerImpl) LoggerImpl {
	return &Logger{
		ZapLogger:                 zapLogger,
		LogLevel:                  logger.Warn,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
		Handler:                   nil,
	}
}

func (w *Logger) SetAsDefault() {
	logger.Default = w
}

func (w *Logger) LogMode(level logger.LogLevel) logger.Interface {
	return &Logger{
		ZapLogger:                 w.ZapLogger,
		SlowThreshold:             w.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          w.SkipCallerLookup,
		IgnoreRecordNotFoundError: w.IgnoreRecordNotFoundError,
		Handler:                   w.Handler,
	}
}

func (w *Logger) Info(ctx context.Context, str string, args ...any) {
	if w.LogLevel >= logger.Info {
		w.zapLogger(ctx).Debug(fmt.Sprintf(str, args...))
	}
}

func (w *Logger) Warn(ctx context.Context, str string, args ...any) {
	if w.LogLevel >= logger.Warn {
		w.zapLogger(ctx).Warn(fmt.Sprintf(str, args...))
	}
}

func (w *Logger) Error(ctx context.Context, str string, args ...any) {
	if w.LogLevel >= logger.Error {
		w.zapLogger(ctx).Error(fmt.Sprintf(str, args...))
	}
}

func (w *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if w.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	zapLogger := w.zapLogger(ctx)

	if w.LogLevel >= logger.Info || w.LogLevel >= logger.Warn || w.LogLevel >= logger.Error {
		sql, rows := fc()
		message := fmt.Sprintf("elapsed= %s, rows= %d, sql= %s\n", elapsed.String(), rows, sql)
		switch {
		case err != nil && w.LogLevel >= logger.Error && (!w.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
			zapLogger.Error(message, zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
			break

		case w.SlowThreshold != 0 && elapsed > w.SlowThreshold && w.LogLevel >= logger.Warn:
			zapLogger.Warn(message, zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
			break

		case w.LogLevel >= logger.Info:
			zapLogger.Debug(message, zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
			break
		}
	}
}

func (w *Logger) zapLogger(ctx context.Context) ZapLoggerImpl {
	zapLogger := w.ZapLogger

	// options
	skip := 12
	start := 2

	if w.Handler != nil {
		fields := w.Handler(ctx)
		zapLogger = zapLogger.With(fields...)
	}

	if w.SkipCallerLookup {
		return zapLogger
	}

	skip += start
	for i := start; i <= skip; i++ {
		pc, file, line, ok := runtime.Caller(i)
		nokocore.KeepVoid(pc, file, line, ok)

		switch {
		case !ok:
			break

		case strings.HasSuffix(file, "_test.go"):
			break

		case strings.Contains(file, gormPackage):
			break

		case strings.Contains(file, zapgormPackage):
			break

		default:
			return zapLogger.WithOptions(zap.AddCallerSkip(i))
		}
	}

	return zapLogger
}
