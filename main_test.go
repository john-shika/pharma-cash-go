package main

import (
	"fmt"
	"github.com/spf13/viper"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestConfig(t *testing.T) {
	func() {
		var ok bool
		var err error
		var nokoWebApiSelfRunEnv string
		nokocore.KeepVoid(ok, err, nokoWebApiSelfRunEnv)

		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigFile("nokowebapi.yaml")

		if err = viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}()

	func() {
		temp := &nokocore.ArrayStr{
			"cat",
			"dog",
		}

		nokocore.DeleteArrayItemReflect(temp, 0)

		fmt.Println(temp)

		globals.Del("jwt.audience.1")
		globals.Set("jwt.audience.0", "im-audience")
		console.Dir(globals.Get("jwt.audience"))
		console.Dir(globals.Get("jwt"))

		globals.Del("jwt.audience")
		console.Dir(globals.Get("jwt"))
	}()

	func() {
		temp := &[]string{
			"cat",
			"dog",
		}

		fmt.Println(nokocore.GetArrayItemReflect(temp, 0))
		nokocore.SetArrayItemReflect(temp, 0, "bird")
		console.Dir(temp)
		nokocore.DeleteArrayItemReflect(temp, 1)
		console.Dir(temp)
		fmt.Println(nokocore.GetArrayItem(*temp, 0))
		fmt.Println(nokocore.PassTypeIndirectReflect(temp))
	}()

	func() {
		type Person struct {
			Name     string    `mapstructure:"name"`
			Age      int       `mapstructure:"age"`
			Gender   string    `mapstructure:"gender"`
			Birthday time.Time `mapstructure:"birthday"`
		}

		type Family struct {
			Members []Person `mapstructure:"members"`
		}

		family := &Family{}

		nokocore.NoErr(nokocore.SetValueReflect(family, map[string]any{
			"members": []map[string]any{
				{
					"name":     "John Doe",
					"age":      30,
					"gender":   "male",
					"birthday": time.Now(),
				},
			},
		}))

		console.Dir(family)

		temp := &map[string]any{}

		nokocore.NoErr(nokocore.SetValueReflect(temp, map[string]any{
			"family": family,
		}))

		console.Dir(temp)
	}()
}
