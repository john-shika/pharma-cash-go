package console

import (
	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"nokowebapi/xterm"
)

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

func makeLogger() *zap.Logger {
	var logger *zap.Logger
	nokocore.KeepVoid(logger)

	isDevelopment := globals.IsDevelopment()
	loggerConfig := globals.GetLoggerConfig()

	writerSyncer := GetWriterSyncer(xterm.Stdout)
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

	logger = logger.Named("<NokoWebApi>")
	zap.ReplaceGlobals(logger)
	return logger
}

var Logger = makeLogger()

func Debug(msg string, fields ...zap.Field) {
	msg = color.New(color.FgCyan).Sprint(msg)
	Logger.Debug(msg, fields...)
}

func Dir(obj any, fields ...zap.Field) {
	temp := nokocore.ShikaYamlEncode(obj)
	temp = color.New(color.FgYellow).Sprint(temp)
	Logger.Info(temp, fields...)
}

func Log(msg string, fields ...zap.Field) {
	msg = color.New(color.FgGreen).Sprint(msg)
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	msg = color.New(color.FgYellow).Sprint(msg)
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	msg = color.New(color.FgRed).Sprint(msg)
	Logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	msg = color.New(color.FgRed).Sprint(msg)
	Logger.Fatal(msg, fields...)
}
