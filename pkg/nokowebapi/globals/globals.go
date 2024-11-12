package globals

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
	"nokowebapi/nokocore"
	"nokowebapi/task"
	"strings"
)

var defaultConfig = nokocore.MapAny{
	"nokowebapi": nokocore.MapAny{
		"name":        "nokowebapi",
		"description": "Nokotan Backend Web API",
		"version":     "v1.0.0-dev",
		"author":      "<ahmadasysyafiq@proton.me>",
		"url":         "https://github.com/john-shika/nokotan-backend-go",
		"license_url": "https://www.apache.org/licenses/LICENSE-2.0",
		"license":     "Apache-2.0",
		"production":  false,
	},
	"jwt": nokocore.MapAny{
		"algorithm":  nil,
		"audience":   "your-audience",
		"issuer":     "your-issuer",
		"secret_key": "your-super-secret-key-keep-it-mind-dont-tell-anyone",
		"expires_in": "1h",
	},
	"logger": nokocore.MapAny{
		"level":               "debug",
		"encoding":            "console",
		"stack_trace_enabled": true,
		"colorable":           true,
	},
	"tasks": nokocore.ArrayAny{},
}

func getNameKeyType[T any]() string {
	return getNameKey(new(T))
}

func getNameKey(obj any) string {
	var ok bool
	nokocore.KeepVoid(ok)

	name := nokocore.ToSnakeCase(nokocore.GetNameType(obj))
	name, ok = strings.CutSuffix(name, "_config")
	return name
}

type ConfigImpl interface {
	IsDevelopment() bool
	GetJwtConfig() *nokocore.JwtConfig
	GetLoggerConfig() *nokocore.LoggerConfig
	GetTasks() *task.Tasks
	Keys() []string
	Values() []any
	Get(key string) any
}

type Config struct {
	jwtConfig    *nokocore.JwtConfig
	loggerConfig *nokocore.LoggerConfig
	tasks        *task.Tasks
	locker       nokocore.LockerImpl
}

func NewConfig() ConfigImpl {
	return &Config{
		jwtConfig:    nil,
		loggerConfig: nil,
		tasks:        nil,
		locker:       nokocore.NewLocker(),
	}
}

func (c *Config) IsDevelopment() bool {
	return !c.Get("nokowebapi.production").(bool)
}

func (c *Config) GetJwtConfig() *nokocore.JwtConfig {
	c.locker.Lock(func() {
		if c.jwtConfig != nil {
			return
		}
		//c.jwtConfig = nokocore.Unwrap(ViperConfigUnmarshal[nokocore.JwtConfig]())
		c.jwtConfig = GetConfigGlobals[nokocore.JwtConfig]()
	})
	return c.jwtConfig
}

func (c *Config) GetLoggerConfig() *nokocore.LoggerConfig {
	c.locker.Lock(func() {
		if c.loggerConfig != nil {
			return
		}
		//c.loggerConfig = nokocore.Unwrap(ViperConfigUnmarshal[nokocore.LoggerConfig]())
		c.loggerConfig = GetConfigGlobals[nokocore.LoggerConfig]()
	})
	return c.loggerConfig
}

func (c *Config) GetTasks() *task.Tasks {
	c.locker.Lock(func() {
		if c.tasks != nil {
			return
		}
		c.tasks = GetConfigGlobals[task.Tasks]()
	})
	return c.tasks
}

func (c *Config) Keys() []string {
	return defaultConfig.Keys()
}

func (c *Config) Values() []any {
	return defaultConfig.Values()
}

func (c *Config) Get(key string) any {
	return defaultConfig.Get(key)
}

var globals = NewConfig()

func IsDevelopment() bool {
	return globals.IsDevelopment()
}

func GetJwtConfig() *nokocore.JwtConfig {
	return globals.GetJwtConfig()
}

func GetLoggerConfig() *nokocore.LoggerConfig {
	return globals.GetLoggerConfig()
}

func GetTasks() *task.Tasks {
	return globals.GetTasks()
}

func Keys() []string {
	return globals.Keys()
}

func Values() []any {
	return globals.Values()
}

func Get(key string) any {
	return globals.Get(key)
}

func ViperConfigUnmarshal[T any]() (*T, error) {
	var err error
	nokocore.KeepVoid(err)

	config := new(T)
	keyName := getNameKey(config)
	if err = viper.UnmarshalKey(keyName, config); err != nil {
		return nil, err
	}

	return config, nil
}

// GetConfigGlobals method, this method is used to get the default config or viper config with some problem issues.
func GetConfigGlobals[T any]() *T {
	var err error
	nokocore.KeepVoid(err)

	config := new(T)

	key := getNameKey(config)
	locals := defaultConfig.Get(key)

	if config, err = ViperConfigUnmarshal[T](); err != nil {
		panic("failed to unmarshal viper config")
	}

	// TODO: Implement support for setting values in nested maps within nested structs.
	//nokocore.ForEachStructFieldsReflect(config, false, func(name string, sFieldX nokocore.StructFieldExpandedImpl) {
	//	if sFieldX.IsZero() {
	//		nokocore.SetValueReflect(sFieldX.GetValue(), locals.Get(name))
	//		//val := nokocore.GetValueReflect(locals.Get(name))
	//		//sFieldX.Set(val)
	//		return
	//	}
	//	locals.Set(name, sFieldX.Interface())
	//})

	nokocore.NoErr(mapstructure.Decode(config, &locals))
	defaultConfig.Set(key, locals)

	return config
}
