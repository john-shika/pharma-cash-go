package console

import (
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"moul.io/zapgorm2"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"nokowebapi/xterm"
)

type LoggerImpl interface {
	ZapLogger() *zap.Logger
	GORMLogger() zapgorm2.Logger
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Dir(obj any, fields ...zap.Field)
	Log(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

type Logger struct {
	*zap.Logger
}

func NewLogger(logger *zap.Logger) LoggerImpl {
	return &Logger{
		Logger: logger,
	}
}

func (l *Logger) ZapLogger() *zap.Logger {
	return l.Logger
}

func (l *Logger) GORMLogger() zapgorm2.Logger {
	return zapgorm2.New(l.Logger)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	msg = color.New(color.FgCyan).Sprint(msg)
	l.Logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	msg = color.New(color.FgCyan).Sprint(msg)
	l.Logger.Info(msg, fields...)
}

func (l *Logger) Dir(obj any, fields ...zap.Field) {
	temp := nokocore.ShikaYamlEncode(obj)
	temp = color.New(color.FgYellow).Sprint(temp)
	l.Logger.Info(temp, fields...)
}

func (l *Logger) Log(msg string, fields ...zap.Field) {
	msg = color.New(color.FgGreen).Sprint(msg)
	l.Logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	msg = color.New(color.FgYellow).Sprint(msg)
	l.Logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	msg = color.New(color.FgRed).Sprint(msg)
	l.Logger.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	msg = color.New(color.FgRed).Sprint(msg)
	l.Logger.Fatal(msg, fields...)
}

func GetWriterSyncer(stdout any) zapcore.WriteSyncer {
	var ok bool
	var writer io.Writer
	var writeSyncer zapcore.WriteSyncer
	nokocore.KeepVoid(ok, writer, writeSyncer)

	if writeSyncer, ok = stdout.(zapcore.WriteSyncer); !ok {
		if writer, ok = stdout.(io.Writer); !ok {
			panic("failed to convert stdout to either zapcore.WriteSyncer or io.Writer")
		}
		return zapcore.Lock(zapcore.AddSync(writer))
	}

	return writeSyncer
}

func makeLogger(name string) LoggerImpl {
	var zapLogger *zap.Logger
	nokocore.KeepVoid(zapLogger)

	isDevelopment := globals.IsDevelopment()
	loggerConfig := globals.GetLoggerConfig()

	writerSyncer := GetWriterSyncer(xterm.Stdout)
	level := loggerConfig.GetLevel()

	options := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	}

	if loggerConfig.StackTraceEnabled {
		options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	options = append(options, zap.IncreaseLevel(level))

	if isDevelopment {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoder := loggerConfig.GetEncoder(encoderConfig)
		core := zapcore.NewCore(encoder, writerSyncer, level)
		zapLogger = zap.New(core, options...)
	} else {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoder := loggerConfig.GetEncoder(encoderConfig)
		core := zapcore.NewCore(encoder, writerSyncer, level)
		zapLogger = zap.New(core, options...)
	}

	zapLogger = zapLogger.Named(fmt.Sprintf("[%s]", name))
	zap.ReplaceGlobals(zapLogger)
	return NewLogger(zapLogger)
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

func Log(msg string, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Log(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger := GetLogger("NokoWebApi.Console")
	logger.Fatal(msg, fields...)
}
