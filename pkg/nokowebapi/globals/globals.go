package globals

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
	"nokowebapi/nokocore"
	"nokowebapi/task"
	"reflect"
	"strings"
)

// WARN: except 'nokowebapi' will be overwritten.
var defaultConfig = &nokocore.MapAny{
	"nokowebapi": &nokocore.MapAny{
		"name":        "nokowebapi",
		"description": "Nokotan Backend Web API",
		"version":     "v1.0.0-dev",
		"author":      "ahmadasysyafiq@proton.me",
		"url":         "https://github.com/john-shika/nokotan-backend-go",
		"license_url": "https://www.apache.org/licenses/LICENSE-2.0",
		"license":     "Apache-2.0",
		"production":  false,
	},
	"jwt": &nokocore.MapAny{
		"algorithm": "HS256",
		"audience": &nokocore.ArrayStr{
			"your-audience-1",
			"your-audience-2",
			"your-audience-3",
		},
		"issuer":     "your-issuer",
		"secret_key": "your-super-secret-key-keep-it-mind-dont-tell-anyone",
		"expires_in": "1h",
	},
	"logger": &nokocore.MapAny{
		"level":               "debug",
		"encoding":            "console",
		"stack_trace_enabled": true,
		"colorable":           true,
	},
	"tasks": &nokocore.ArrayAny{},
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
	Set(key string, value any) bool
	Del(key string) bool
}

type Config struct {
	jwt    *nokocore.JwtConfig
	logger *nokocore.LoggerConfig
	tasks  *task.Tasks
	locker nokocore.LockerImpl
}

func NewConfig() ConfigImpl {
	return &Config{
		jwt:    nil,
		logger: nil,
		tasks:  nil,
		locker: nokocore.NewLocker(),
	}
}

func (c *Config) IsDevelopment() bool {
	return !c.Get("nokowebapi.production").(bool)
}

func (c *Config) GetJwtConfig() *nokocore.JwtConfig {
	c.locker.Lock(func() {
		if c.jwt != nil {
			return
		}
		//c.jwt = nokocore.Unwrap(ViperConfigUnmarshal[nokocore.JwtConfig]())
		c.jwt = GetConfigGlobals[nokocore.JwtConfig]()
	})
	return c.jwt
}

func (c *Config) GetLoggerConfig() *nokocore.LoggerConfig {
	c.locker.Lock(func() {
		if c.logger != nil {
			return
		}
		//c.logger = nokocore.Unwrap(ViperConfigUnmarshal[nokocore.LoggerConfig]())
		c.logger = GetConfigGlobals[nokocore.LoggerConfig]()
	})
	return c.logger
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
	return defaultConfig.QueryGet(key)
}

func (c *Config) Set(key string, value any) bool {
	return defaultConfig.QuerySet(key, value)
}

func (c *Config) Del(key string) bool {
	return defaultConfig.QueryDel(key)
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

func GetTaskConfig(name string) task.ConfigImpl {
	return globals.GetTasks().GetTaskConfig(name)
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

func Set(key string, value any) bool {
	return globals.Set(key, value)
}

func Del(key string) bool {
	return globals.Del(key)
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
	keyName := getNameKey(config)

	var locals any

	localized := false
	if defaultConfig.HasKey(keyName) {
		locals = defaultConfig.QueryGet(keyName)
		localized = true
	}

	if config, err = ViperConfigUnmarshal[T](); err != nil {
		panic("failed to unmarshal viper config")
	}

	vConfig := nokocore.PassValueIndirectReflect(config)
	switch vConfig.Kind() {
	case reflect.Struct:

		// no locals config found
		if !localized {
			break
		}

		// read config and set to locals if is necessary
		options := nokocore.NewForEachStructFieldsOptions()
		nokocore.NoErr(nokocore.ForEachStructFieldsReflect(config, options, func(name string, sFieldX nokocore.StructFieldExImpl) error {

			// some config has zero value, set from locals
			if sFieldX.IsZero() && sFieldX.Kind() != reflect.Bool {
				val := nokocore.PassValueIndirectReflect(nokocore.GetMapValueReflect(locals, name))
				//nokocore.SetValueReflect(sFieldX.GetValue(), val)
				sFieldX.Set(val)
				return nil
			}

			// TODO: check the type
			// maybe the value strict with interface or pointer
			nokocore.SetMapValueReflect(locals, name, sFieldX.GetValue())
			return nil
		}))

	default:
		nokocore.KeepVoid(mapstructure.Decode(config, locals))
	}

	// set back to locals config
	if localized {
		defaultConfig.QuerySet(keyName, locals)
	}

	return config
}
