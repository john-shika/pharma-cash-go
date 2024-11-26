package nokocore

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"net/url"
	"time"
)

const TimeOnlyFormat = "15:04:05"
const DateOnlyFormat = "2006-01-02"
const DateTimeFormat = "2006-01-02 15:04:05"
const DateTimeFormatISO8601 = "2006-01-02T15:04:05.000Z07:00"

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
	return GetTimeUtcNow().Format(DateTimeFormatISO8601)
}

func ParseTimeUtcByStringISO8601(value string) (time.Time, error) {
	var err error
	var t time.Time
	if t, err = time.Parse(DateTimeFormatISO8601, value); err != nil {
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
		return ParseTimeUtcByStringISO8601(value.(string))
	case int64:
		return GetTimeUtcByTimestamp(value.(int64)), nil
	case int32:
		return GetTimeUtcByTimestamp(int64(value.(int32))), nil
	case int:
		return GetTimeUtcByTimestamp(int64(value.(int))), nil
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
		return Default[time.Time]().Format(DateTimeFormatISO8601)
	}
	return t.Format(DateTimeFormatISO8601)
}

func GetURL(value any) (*url.URL, error) {
	switch value.(type) {
	case *url.URL:
		return value.(*url.URL), nil
	case string:
		return url.Parse(value.(string))
	default:
		return nil, errors.New("invalid data type")
	}
}

func IsURL(value any) bool {
	var err error
	var check *url.URL
	KeepVoid(err, check)

	if check, err = GetURL(value); err != nil {
		return false
	}

	return true
}

func ToURLString(value any) string {
	return Unwrap(GetURL(value)).String()
}

func GetUUID(value any) (uuid.UUID, error) {
	switch value.(type) {
	case uuid.UUID:
		return value.(uuid.UUID), nil
	case string:
		return uuid.Parse(value.(string))
	default:
		return uuid.Nil, errors.New("invalid data type")
	}
}

func IsUUID(value any) bool {
	var err error
	var check uuid.UUID
	KeepVoid(err, check)

	if check, err = GetUUID(value); err != nil {
		return false
	}

	return true
}

func ToUUIDString(value any) string {
	return Unwrap(GetUUID(value)).String()
}

func GetDecimal(value any) (decimal.Decimal, error) {
	switch value.(type) {
	case decimal.Decimal:
		return value.(decimal.Decimal), nil
	case string:
		return decimal.NewFromString(value.(string))
	default:
		return decimal.NewFromInt(0), errors.New("invalid data type")
	}
}

func IsDecimal(value any) bool {
	var err error
	var check decimal.Decimal
	KeepVoid(err, check)

	if check, err = GetDecimal(value); err != nil {
		return false
	}

	return true
}

func ToDecimalString(value any) string {
	return Unwrap(GetDecimal(value)).String()
}
