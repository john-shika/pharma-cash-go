package sqlx

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"nokowebapi/nokocore"
	"strconv"
	"strings"
	"time"
)

type TimeOnlyImpl interface {
	MarshalText() (text []byte, err error)
	UnmarshalText(text []byte) (err error)
	MarshalJSON() (data []byte, err error)
	UnmarshalJSON(data []byte) (err error)
	Value() (driver.Value, error)
	Scan(value any) error
	String() string
	ToTimeDuration() time.Duration
	ToString() string
}

type NullTimeOnly struct {
	TimeOnly TimeOnly
	Valid    bool
}

func NewTimeOnly(value time.Time) NullTimeOnly {
	return NullTimeOnly{
		TimeOnly: TimeOnly{
			Time: value,
		},
		Valid: true,
	}
}

func (w *NullTimeOnly) baseInit() {
	var timeOnly TimeOnly
	if w.TimeOnly != timeOnly {
		return
	}

	w.TimeOnly = timeOnly
}

func (w NullTimeOnly) MarshalText() (text []byte, err error) {
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

func (w NullTimeOnly) MarshalJSON() (data []byte, err error) {
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

func (w NullTimeOnly) String() string {
	w.baseInit()
	return w.TimeOnly.String()
}

func (w *NullTimeOnly) ToTimeDuration() time.Duration {
	w.baseInit()
	return w.TimeOnly.ToTimeDuration()
}

func (w *NullTimeOnly) ToString() string {
	w.baseInit()
	return w.TimeOnly.ToString()
}

// TimeOnly custom type to store only time (hh:mm:ss)
type TimeOnly struct {
	time.Time
}

func SafeParseTimeOnly(value string) (NullTimeOnly, error) {
	var err error
	var temp time.Time
	nokocore.KeepVoid(err, temp)

	if temp, err = time.Parse(nokocore.TimeOnlyFormat, value); err != nil {
		return NullTimeOnly{
			TimeOnly: TimeOnly{},
			Valid:    false,
		}, err
	}

	return NullTimeOnly{
		TimeOnly: TimeOnly{
			Time: temp.UTC(),
		},
		Valid: true,
	}, nil
}

func ParseTimeOnly(value string) NullTimeOnly {
	return nokocore.Unwrap(SafeParseTimeOnly(value))
}

func SafeParseTimeOnlyNotNull(value string) (TimeOnly, error) {
	var err error
	var timeOnlyNull NullTimeOnly
	nokocore.KeepVoid(err, timeOnlyNull)

	if timeOnlyNull, err = SafeParseTimeOnly(value); err != nil {
		return TimeOnly{}, err
	}

	if !timeOnlyNull.Valid {
		return TimeOnly{}, errors.New("invalid date")
	}

	return timeOnlyNull.TimeOnly, nil
}

func ParseTimeOnlyNotNull(value string) TimeOnly {
	return nokocore.Unwrap(SafeParseTimeOnly(value)).TimeOnly
}

// MarshalText for text marshaling
func (w TimeOnly) MarshalText() (text []byte, err error) {
	return []byte(w.Format(nokocore.TimeOnlyFormat)), nil
}

// UnmarshalText for text unmarshalling
func (w *TimeOnly) UnmarshalText(text []byte) (err error) {
	var temp time.Time
	nokocore.KeepVoid(temp)

	value := string(text)

	if temp, err = time.Parse(nokocore.TimeOnlyFormat, value); err != nil {
		return err
	}

	w.Time = temp.UTC()
	return nil
}

// MarshalJSON for JSON marshaling
func (w TimeOnly) MarshalJSON() (data []byte, err error) {
	return []byte(strconv.Quote(w.Format(nokocore.TimeOnlyFormat))), nil
}

// UnmarshalJSON for JSON unmarshalling
func (w *TimeOnly) UnmarshalJSON(data []byte) (err error) {
	var value string
	var temp time.Time
	nokocore.KeepVoid(value, temp)

	if value, err = strconv.Unquote(string(data)); err != nil {
		return err
	}

	if temp, err = time.Parse(nokocore.TimeOnlyFormat, value); err != nil {
		return err
	}

	w.Time = temp.UTC()
	return nil
}

// Value implements the driver Valuer interface.
func (w TimeOnly) Value() (driver.Value, error) {
	return w.Format(nokocore.TimeOnlyFormat), nil
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
			if temp, err = time.Parse(nokocore.TimeOnlyFormat, string(val)); err != nil {
				return err
			}

			*w = TimeOnly{
				Time: temp.UTC(),
			}
			return nil

		case string:
			if temp, err = time.Parse(nokocore.TimeOnlyFormat, val); err != nil {
				return err
			}

			*w = TimeOnly{
				Time: temp.UTC(),
			}
			return nil

		default:
			return errors.New(fmt.Sprintf("cannot convert %v to time only", value))
		}
	}

	*w = TimeOnly{}
	return nil
}

func (w TimeOnly) String() string {
	return w.Format(nokocore.TimeOnlyFormat)
}

func (w *TimeOnly) ToTimeDuration() time.Duration {
	hours := time.Duration(w.Hour()) * time.Hour
	minutes := time.Duration(w.Minute()) * time.Minute
	seconds := time.Duration(w.Second()) * time.Second
	return hours + minutes + seconds
}

func (w *TimeOnly) ToString() string {
	return w.String()
}
