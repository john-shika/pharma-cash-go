package main

import (
	"encoding/json"
	"fmt"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
)

func main() {

	timeOnly := sqlx.ParseTimeOnly("07:00:00")

	fmt.Println(string(nokocore.Unwrap(json.Marshal(timeOnly))))
	fmt.Println(timeOnly)
	
	fmt.Println(nokocore.ToPascalCase("ttk"))
}
