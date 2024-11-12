package console

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"nokowebapi/xterm"
)

var WriterSyncer zapcore.WriteSyncer

func updateWriterSyncer(stdout io.Writer) zapcore.WriteSyncer {
	return zapcore.Lock(zapcore.AddSync(stdout))
}

func NewWriterSyncer(stdout io.Writer) zapcore.WriteSyncer {
	if WriterSyncer != nil {
		return WriterSyncer
	}

	return updateWriterSyncer(stdout)
}

func makeLogger() *zap.Logger {
	var logger *zap.Logger
	nokocore.KeepVoid(logger)

	isDevelopment := globals.IsDevelopment()
	loggerConfig := globals.GetLoggerConfig()
	writerSyncer := NewWriterSyncer(xterm.Stdout)
	level := loggerConfig.GetLevel()

	options := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(4),
	}

	if loggerConfig.StackTraceEnabled {
		options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	options = append(options, zap.IncreaseLevel(level))

	if isDevelopment {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoder := loggerConfig.GetEncoder(encoderConfig)
		core := zapcore.NewCore(encoder, writerSyncer, level)
		logger = zap.New(core, options...)
	} else {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoder := loggerConfig.GetEncoder(encoderConfig)
		core := zapcore.NewCore(encoder, writerSyncer, level)
		logger = zap.New(core, options...)
	}

	logger = logger.Named("[NokoWebApi]")
	zap.ReplaceGlobals(logger)
	return logger
}

var Logger = makeLogger()

func Dir(obj any, fields ...zap.Field) {
	Logger.Info(nokocore.ShikaYamlEncode(obj), fields...)
}

func Log(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Logf(format string, args ...any) {
	Logger.Info(fmt.Sprintf(format, args...))
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	defer updateWriterSyncer(xterm.Stdout)
	updateWriterSyncer(xterm.Stderr)

	Logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	defer updateWriterSyncer(xterm.Stdout)
	updateWriterSyncer(xterm.Stderr)

	Logger.Fatal(msg, fields...)
}
