package main

import (
	"fmt"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
)

func main() {
	config := globals.Globals().GetJwtConfig()
	fmt.Printf("%+v\n", config)

	temp := nokocore.MapAny{
		"test": &nokocore.MapAny{
			"value": 12,
		},
	}

	nokocore.SetValueReflect(temp.Get("test"), &nokocore.MapAny{
		"value": 24,
	})

	fmt.Println(nokocore.ShikaYamlEncode(temp))
}
