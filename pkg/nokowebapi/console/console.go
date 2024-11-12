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

func updateWriterSyncer(stdout any) zapcore.WriteSyncer {
	var ok bool
	var writer io.Writer
	var writeSyncer zapcore.WriteSyncer
	nokocore.KeepVoid(ok, writer, writeSyncer)

	if writeSyncer, ok = stdout.(zapcore.WriteSyncer); !ok {
		if writer, ok = stdout.(io.Writer); !ok {
			panic("failed to convert stdout to either zapcore.WriteSyncer or io.Writer")
		}
		WriterSyncer = zapcore.Lock(zapcore.AddSync(writer))
		return WriterSyncer
	}

	WriterSyncer = writeSyncer
	return WriterSyncer
}

func makeLogger() *zap.Logger {
	var logger *zap.Logger
	nokocore.KeepVoid(logger)

	isDevelopment := globals.IsDevelopment()
	loggerConfig := globals.GetLoggerConfig()
	writerSyncer := updateWriterSyncer(xterm.Stdout)
	level := loggerConfig.GetLevel()

	options := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(1),
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
