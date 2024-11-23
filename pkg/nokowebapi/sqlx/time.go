package sqlx

import (
	"database/sql/driver"
	"fmt"
	"nokowebapi/nokocore"
	"strconv"
	"strings"
	"time"
)

const TimeOnlyFormat = "15:04:05"

type TimeOnlyImpl interface {
	MarshalText() (text []byte, err error)
	UnmarshalText(text []byte) (err error)
	MarshalJSON() (data []byte, err error)
	UnmarshalJSON(data []byte) (err error)
	Value() (driver.Value, error)
	Scan(value any) error
	ToTimeDuration() time.Duration
}

type NullTimeOnly struct {
	TimeOnly TimeOnlyImpl
	Valid    bool
}

func NewTimeOnly(value time.Time) NullTimeOnly {
	return NullTimeOnly{
		TimeOnly: &TimeOnly{
			Time: value,
		},
		Valid: true,
	}
}

func (w *NullTimeOnly) baseInit() {
	if w.TimeOnly != nil {
		return
	}

	w.TimeOnly = &TimeOnly{
		Time: time.Time{},
	}
}

func (w *NullTimeOnly) MarshalText() (text []byte, err error) {
	if w.Valid {
		w.baseInit()
		return w.TimeOnly.MarshalText()
	}

	return []byte("NULL"), nil
}

func (w *NullTimeOnly) UnmarshalText(text []byte) (err error) {
	if !strings.EqualFold(string(text), "NULL") {
		w.baseInit()
		return w.TimeOnly.UnmarshalText(text)
	}

	w.Valid = false
	return nil
}

func (w *NullTimeOnly) MarshalJSON() (data []byte, err error) {
	if w.Valid {
		w.baseInit()
		return w.TimeOnly.MarshalJSON()
	}

	return []byte("null"), nil
}

func (w *NullTimeOnly) UnmarshalJSON(data []byte) (err error) {
	if string(data) != "null" {
		w.baseInit()
		return w.TimeOnly.UnmarshalJSON(data)
	}

	w.Valid = false
	return nil
}

func (w NullTimeOnly) Value() (driver.Value, error) {
	if w.Valid {
		w.baseInit()
		return w.TimeOnly.Value()
	}

	return nil, nil
}

func (w *NullTimeOnly) Scan(value any) error {
	if value != nil {
		w.baseInit()
		return w.TimeOnly.Scan(value)
	}

	w.Valid = false
	return nil
}

func (w NullTimeOnly) ToTimeDuration() time.Duration {
	w.baseInit()
	return w.TimeOnly.ToTimeDuration()
}

// TimeOnly custom type to store only time (hh:mm:ss)
type TimeOnly struct {
	time.Time
}

func SafeParseTimeOnly(value string) (NullTimeOnly, error) {
	var err error
	var temp time.Time
	nokocore.KeepVoid(err, temp)

	if temp, err = time.Parse(TimeOnlyFormat, value); err != nil {
		return NullTimeOnly{
			TimeOnly: &TimeOnly{
				Time: time.Time{},
			},
			Valid: false,
		}, err
	}

	return NullTimeOnly{
		TimeOnly: &TimeOnly{
			Time: temp,
		},
		Valid: true,
	}, nil
}

func ParseTimeOnly(value string) NullTimeOnly {
	return nokocore.Unwrap(SafeParseTimeOnly(value))
}

// MarshalText for text marshaling
func (w *TimeOnly) MarshalText() (text []byte, err error) {
	return []byte(w.Format(TimeOnlyFormat)), nil
}

// UnmarshalText for text unmarshalling
func (w *TimeOnly) UnmarshalText(text []byte) (err error) {
	var temp time.Time
	nokocore.KeepVoid(temp)

	value := string(text)

	if temp, err = time.Parse(TimeOnlyFormat, value); err != nil {
		return err
	}

	w.Time = temp
	return nil
}

// MarshalJSON for JSON marshaling
func (w *TimeOnly) MarshalJSON() (data []byte, err error) {
	return []byte(strconv.Quote(w.Format(TimeOnlyFormat))), nil
}

// UnmarshalJSON for JSON unmarshalling
func (w *TimeOnly) UnmarshalJSON(data []byte) (err error) {
	var value string
	var temp time.Time
	nokocore.KeepVoid(value, temp)

	if value, err = strconv.Unquote(string(data)); err != nil {
		return err
	}

	if temp, err = time.Parse(TimeOnlyFormat, value); err != nil {
		return err
	}

	w.Time = temp
	return nil
}

// Value implements the driver Valuer interface.
func (w TimeOnly) Value() (driver.Value, error) {
	return w.Format(TimeOnlyFormat), nil
}

// Scan implements the Scanner interface.
func (w *TimeOnly) Scan(value any) error {
	var err error
	var temp time.Time
	nokocore.KeepVoid(err, temp)

	if value != nil {
		switch val := value.(type) {
		case time.Time:
			*w = TimeOnly{
				Time: val,
			}
			return nil

		case []byte:
			if temp, err = time.Parse(TimeOnlyFormat, string(val)); err != nil {
				return err
			}

			*w = TimeOnly{
				Time: temp,
			}
			return nil

		case string:
			if temp, err = time.Parse(TimeOnlyFormat, val); err != nil {
				return err
			}

			*w = TimeOnly{
				Time: temp,
			}
			return nil

		default:
			return fmt.Errorf("cannot convert %v to TimeOnly", value)
		}
	}

	*w = TimeOnly{
		Time: time.Time{},
	}
	return nil
}

func (w *TimeOnly) ToTimeDuration() time.Duration {
	hours := time.Duration(w.Hour()) * time.Hour
	minutes := time.Duration(w.Minute()) * time.Minute
	seconds := time.Duration(w.Second()) * time.Second
	return hours + minutes + seconds
}
