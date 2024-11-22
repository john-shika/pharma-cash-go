package nokocore

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/url"
	"time"
)

const TimeFormatISO8601 = "2006-01-02T15:04:05.000Z07:00"

func GetTimeUtcNow() time.Time {
	return time.Now().UTC()
}

func GetTimeUtcNowTimestamp() int64 {
	return GetTimeUtcNow().UnixMilli()
}

func GetTimeUtcByTimestamp(timeStamp int64) time.Time {
	return time.UnixMilli(timeStamp).UTC()
}

func GetTimeUtcNowStringISO8601() string {
	return GetTimeUtcNow().Format(TimeFormatISO8601)
}

func ParseTimeUtcByStringISO8601(value string) (time.Time, error) {
	var err error
	var t time.Time
	if t, err = time.Parse(TimeFormatISO8601, value); err != nil {
		return Default[time.Time](), err
	}
	return t.UTC(), nil
}

type TimeAnyImpl interface {
	*jwt.NumericDate | time.Time | string | int64 | int32 | int
}

func GetTimeUtcISO8601(value any) (time.Time, error) {
	var err error
	var t time.Time
	KeepVoid(err, t)

	val := PassValueIndirectReflect(value)
	if !val.IsValid() {
		return Default[time.Time](), ErrDataTypeInvalid
	}
	value = val.Interface()
	switch value.(type) {
	case jwt.NumericDate:
		// why not use a pointer for jwt.NumericDate, because
		// has already pass value indirect reflection
		return value.(jwt.NumericDate).Time.UTC(), nil
	case time.Time:
		return value.(time.Time).UTC(), nil
	case string:
		v := Unwrap(Cast[string](value))
		return ParseTimeUtcByStringISO8601(v)
	case int64:
		v := Unwrap(Cast[int64](value))
		return GetTimeUtcByTimestamp(v), nil
	case int32:
		v := Unwrap(Cast[int32](value))
		return GetTimeUtcByTimestamp(int64(v)), nil
	case int:
		v := Unwrap(Cast[int](value))
		return GetTimeUtcByTimestamp(int64(v)), nil
	default:
		return Default[time.Time](), ErrDataTypeInvalid
	}
}

func GetTimeUtcAnyStrict[V TimeAnyImpl](value V) (time.Time, error) {
	return GetTimeUtcISO8601(value)
}

func IsTimeUtcISO8601(value any) bool {
	var err error
	var t time.Time
	KeepVoid(err, t)

	if t, err = GetTimeUtcISO8601(value); err != nil {
		return false
	}
	return true
}

func ToTimeUtcStringISO8601(value any) string {
	var err error
	var t time.Time
	KeepVoid(err, t)

	if t, err = GetTimeUtcISO8601(value); err != nil {
		return Default[time.Time]().Format(TimeFormatISO8601)
	}
	return t.Format(TimeFormatISO8601)
}

func GetURL(value any) (url.URL, error) {
	// TODO: implement it
	return url.URL{}, errors.New("not implemented yet")
}

func IsURL(value any) bool {
	// TODO: implement it
	return false
}

func ToURLString(value any) string {
	// TODO: implement it
	panic("not implemented yet")
}

func GetUUID(value any) (uuid.UUID, error) {
	// TODO: implement it
	return uuid.Nil, errors.New("not implemented yet")
}

func IsUUID(value any) bool {
	// TODO: implement it
	return false
}

func ToUUIDString(value any) string {
	// TODO: implement it
	panic("not implemented yet")
}
