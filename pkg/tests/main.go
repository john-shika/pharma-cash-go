package main

import (
	"fmt"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
)

func main() {

	timeOnlyPtr := nokocore.Unwrap(sqlx.SafeParseTimeOnly("07:00:00")).(*sqlx.TimeOnly)
	timeOnly := *timeOnlyPtr
	
	fmt.Println(string(nokocore.Unwrap(timeOnly.MarshalJSON())))
	fmt.Println(timeOnlyPtr)
}
