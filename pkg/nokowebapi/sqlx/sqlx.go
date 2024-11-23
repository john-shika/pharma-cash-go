package sqlx

import (
	"database/sql"
	"time"
)

func NewValue[T any](value T) sql.Null[T] {
	return sql.Null[T]{V: value, Valid: true}
}

func NewString(value string) sql.NullString {
	if value != "" {
		return sql.NullString{String: value, Valid: true}
	}

	return sql.NullString{}
}

func NewTime(value time.Time) sql.NullTime {
	return sql.NullTime{Time: value, Valid: true}
}

func NewBool(value bool) sql.NullBool {
	return sql.NullBool{Bool: value, Valid: true}
}

func NewByte(value byte) sql.NullByte {
	return sql.NullByte{Byte: value, Valid: true}
}

func NewInt16(value int16) sql.NullInt16 {
	return sql.NullInt16{Int16: value, Valid: true}
}

func NewInt32(value int32) sql.NullInt32 {
	return sql.NullInt32{Int32: value, Valid: true}
}

func NewInt64(value int64) sql.NullInt64 {
	return sql.NullInt64{Int64: value, Valid: true}
}

func NewFloat64(value float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: value, Valid: true}
}
