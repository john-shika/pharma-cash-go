package extras

import (
	"github.com/labstack/echo/v4"
	"nokowebapi/nokocore"
	"strconv"
	"strings"
)

func ParseQueryToBool(ctx echo.Context, name string) bool {
	if value := ParseQueryToString(ctx, name); value != "" {
		switch strings.ToLower(value) {
		case "1", "y", "yes", "true":
			return true

		default:
			return false
		}
	}

	return false
}

func ParseQueryToInt(ctx echo.Context, name string) int {
	return int(ParseQueryToInt64(ctx, name))
}

func ParseQueryToInt32(ctx echo.Context, name string) int32 {
	var err error
	var temp int64
	nokocore.KeepVoid(err, temp)

	if value := ParseQueryToString(ctx, name); value != "" {
		if temp, err = strconv.ParseInt(value, 10, 32); err != nil {
			return 0
		}

		return int32(temp)
	}

	return 0
}

func ParseQueryToInt64(ctx echo.Context, name string) int64 {
	var err error
	var temp int64
	nokocore.KeepVoid(err, temp)

	if value := ParseQueryToString(ctx, name); value != "" {
		if temp, err = strconv.ParseInt(value, 10, 64); err != nil {
			return 0
		}

		return temp
	}

	return 0
}

func ParseQueryToUint(ctx echo.Context, name string) uint {
	return uint(ParseQueryToUint64(ctx, name))
}

func ParseQueryToUint32(ctx echo.Context, name string) uint32 {
	var err error
	var temp uint64
	nokocore.KeepVoid(err, temp)

	if value := ParseQueryToString(ctx, name); value != "" {
		if temp, err = strconv.ParseUint(value, 10, 32); err != nil {
			return 0
		}

		return uint32(temp)
	}

	return 0
}

func ParseQueryToUint64(ctx echo.Context, name string) uint64 {
	var err error
	var temp uint64
	nokocore.KeepVoid(err, temp)

	if value := ParseQueryToString(ctx, name); value != "" {
		if temp, err = strconv.ParseUint(value, 10, 64); err != nil {
			return 0
		}

		return temp
	}

	return 0
}

func ParseQueryToFloat32(ctx echo.Context, name string) float32 {
	var err error
	var temp float64
	nokocore.KeepVoid(err, temp)

	if value := ParseQueryToString(ctx, name); value != "" {
		if temp, err = strconv.ParseFloat(value, 32); err != nil {
			return 0
		}

		return float32(temp)
	}

	return 0
}

func ParseQueryToFloat64(ctx echo.Context, name string) float64 {
	var err error
	var temp float64
	nokocore.KeepVoid(err, temp)

	if value := ParseQueryToString(ctx, name); value != "" {
		if temp, err = strconv.ParseFloat(value, 64); err != nil {
			return 0
		}

		return temp
	}

	return 0
}

func ParseQueryToComplex64(ctx echo.Context, name string) complex64 {
	var err error
	var temp complex128
	nokocore.KeepVoid(err, temp)

	if value := ParseQueryToString(ctx, name); value != "" {
		if temp, err = strconv.ParseComplex(value, 64); err != nil {
			return 0
		}

		return complex64(temp)
	}

	return 0
}

func ParseQueryToComplex128(ctx echo.Context, name string) complex128 {
	var err error
	var temp complex128
	nokocore.KeepVoid(err, temp)

	if value := ParseQueryToString(ctx, name); value != "" {
		if temp, err = strconv.ParseComplex(value, 128); err != nil {
			return 0
		}

		return temp
	}

	return 0
}

func ParseQueryToString(ctx echo.Context, name string) string {
	return strings.TrimSpace(ctx.QueryParam(name))
}
