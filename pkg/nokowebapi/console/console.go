package console

import (
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"nokowebapi/console/zapgorm"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"nokowebapi/xterm"
)

type LoggerImpl interface {
	GetZapStdout() *zap.Logger
	GetZapStderr() *zap.Logger
	GORMLogger() zapgorm.LoggerImpl
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
	Dir(obj any, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Sync() error
	Core() zapcore.Core
	Name() string
}

type Logger struct {
	stdout *zap.Logger
	stderr *zap.Logger
}

func NewLogger(stdout *zap.Logger, stderr *zap.Logger) LoggerImpl {
	return &Logger{
		stdout: stdout,
		stderr: stderr,
	}
}

func (w *Logger) GetZapStdout() *zap.Logger {
	return w.stdout
}

func (w *Logger) GetZapStderr() *zap.Logger {
	return w.stderr
}

func (w *Logger) GORMLogger() zapgorm.LoggerImpl {
	return zapgorm.New(w)
}

func (w *Logger) Sugar() *zap.SugaredLogger {
	// level error, panic, fatal may wrong directly
	return w.stdout.Sugar()
}

func (w *Logger) Named(s string) *zap.Logger {
	// level error, panic, fatal may wrong directly
	return w.stdout.Named(s)
}

func (w *Logger) WithOptions(opts ...zap.Option) *zap.Logger {
	// level error, panic, fatal may wrong directly
	return w.stdout.WithOptions(opts...)
}

func (w *Logger) With(fields ...zap.Field) *zap.Logger {
	// level error, panic, fatal may wrong directly
	return w.stdout.With(fields...)
}

func (w *Logger) WithLazy(fields ...zap.Field) *zap.Logger {
	// level error, panic, fatal may wrong directly
	return w.stdout.WithLazy(fields...)
}

func (w *Logger) Level() zapcore.Level {
	// level error, panic, fatal may wrong directly
	return w.stdout.Level()
}

func (w *Logger) Check(level zapcore.Level, msg string) *zapcore.CheckedEntry {
	// level error, panic, fatal may wrong directly
	return w.stdout.Check(level, msg)
}

func (w *Logger) Log(level zapcore.Level, msg string, fields ...zap.Field) {
	msg = color.New(color.FgGreen).Sprint(msg)
	w.stdout.Log(level, msg, fields...)
}

func (w *Logger) Debug(msg string, fields ...zap.Field) {
	msg = color.New(color.FgCyan).Sprint(msg)
	w.stdout.Debug(msg, fields...)
}

func (w *Logger) Info(msg string, fields ...zap.Field) {
	msg = color.New(color.FgCyan).Sprint(msg)
	w.stdout.Info(msg, fields...)
}

func (w *Logger) Dir(obj any, fields ...zap.Field) {
	temp := nokocore.ShikaYamlEncode(obj)
	w.Info(temp, fields...)
}

func (w *Logger) Warn(msg string, fields ...zap.Field) {
	msg = color.New(color.FgYellow).Sprint(msg)
	w.stdout.Warn(msg, fields...)
}

func (w *Logger) Error(msg string, fields ...zap.Field) {
	msg = color.New(color.FgRed).Sprint(msg)
	w.stderr.Error(msg, fields...)
}

func (w *Logger) Panic(msg string, fields ...zap.Field) {
	msg = color.New(color.FgRed).Sprint(msg)
	w.stderr.Panic(msg, fields...)
}

func (w *Logger) Fatal(msg string, fields ...zap.Field) {
	msg = color.New(color.FgRed).Sprint(msg)
	w.stderr.Fatal(msg, fields...)
}

func (w *Logger) Sync() error {
	// level error, panic, fatal may wrong directly
	return w.stdout.Sync()
}

func (w *Logger) Core() zapcore.Core {
	// level error, panic, fatal may wrong directly
	return w.stdout.Core()
}

func (w *Logger) Name() string {
	// level error, panic, fatal may wrong directly
	return w.stdout.Name()
}

func GetWriterSyncer(stdout any) zapcore.WriteSyncer {
	var ok bool
	var writer io.Writer
	var syncer zapcore.WriteSyncer
	nokocore.KeepVoid(ok, writer, syncer)

	if syncer, ok = stdout.(zapcore.WriteSyncer); !ok {
		if writer, ok = stdout.(io.Writer); !ok {
			panic("failed to convert stdout to either zapcore.WriteSyncer or io.Writer")
		}
		return zapcore.Lock(zapcore.AddSync(writer))
	}

	return syncer
}

func createZapLogger(writer io.Writer, config *nokocore.LoggerConfig, development bool) *zap.Logger {
	var zapLogger *zap.Logger
	nokocore.KeepVoid(zapLogger)

	level := config.GetLevel()
	syncer := GetWriterSyncer(writer)

	options := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	}

	if config.StackTraceEnabled {
		options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	options = append(options, zap.IncreaseLevel(level))

	if development {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoder := config.GetEncoder(encoderConfig)
		core := zapcore.NewCore(encoder, syncer, level)
		zapLogger = zap.New(core, options...)
	} else {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoder := config.GetEncoder(encoderConfig)
		core := zapcore.NewCore(encoder, syncer, level)
		zapLogger = zap.New(core, options...)
	}

	return zapLogger
}

func makeLogger(name string) LoggerImpl {
	development := globals.IsDevelopment()
	config := globals.GetLoggerConfig()

	// stdout logger
	stdout := createZapLogger(xterm.Stdout, config, development)
	stdout = stdout.Named(fmt.Sprintf("[%s]", name))
	zap.ReplaceGlobals(stdout)

	// stderr logger
	stderr := createZapLogger(xterm.Stderr, config, development)
	stderr = stderr.Named(fmt.Sprintf("[%s]", name))

	return NewLogger(stdout, stderr)
}

var cachesLogger = make(map[string]LoggerImpl)

func GetLogger(name string) LoggerImpl {
	var ok bool
	var logger LoggerImpl
	nokocore.KeepVoid(ok, logger)

	if logger, ok = cachesLogger[name]; !ok {
		logger = makeLogger(name)
		cachesLogger[name] = logger
		return logger
	}

	return logger
}

func Debug(msg string, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Info(msg, fields...)
}

func Dir(obj any, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Dir(obj, fields...)
}

func Log(level zapcore.Level, msg string, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Log(level, msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Fatal(msg, fields...)
}
