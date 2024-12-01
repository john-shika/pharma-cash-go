package nokocore

import (
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	Encoding          string `mapstructure:"encoding" json:"encoding" yaml:"encoding"`
	Level             string `mapstructure:"level" json:"level" yaml:"level"`
	StackTraceEnabled bool   `mapstructure:"stack_trace_enabled" json:"stackTraceEnabled" yaml:"stack_trace_enabled"`
	Colorable         bool   `mapstructure:"colorable" json:"colorable" yaml:"colorable"`
}

func NewLoggerConfig() *LoggerConfig {
	return &LoggerConfig{}
}

func (LoggerConfig) GetNameType() string {
	return "Logger"
}

func (w *LoggerConfig) GetLevel() zapcore.Level {
	return GetLoggerConfigLevel(w)
}

func (w *LoggerConfig) GetEncoder(encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
	return GetLoggerConfigEncoder(w, encoderConfig)
}

func GetLoggerConfigLevel(loggerConfig *LoggerConfig) zapcore.Level {
	switch ToCamelCase(loggerConfig.Level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InvalidLevel
	}
}

func customEncoderConfig(encoderConfig zapcore.EncoderConfig, colorable bool) zapcore.EncoderConfig {
	if colorable {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return encoderConfig
}

func GetLoggerConfigEncoder(loggerConfig *LoggerConfig, encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
	switch ToSnakeCase(loggerConfig.Encoding) {
	case "console", "print", "text", "text/plain":
		return zapcore.NewConsoleEncoder(customEncoderConfig(encoderConfig, loggerConfig.Colorable))
	case "json", "application/json":
		return zapcore.NewJSONEncoder(customEncoderConfig(encoderConfig, loggerConfig.Colorable))
	default:
		return zapcore.NewConsoleEncoder(customEncoderConfig(encoderConfig, loggerConfig.Colorable))
	}
}
