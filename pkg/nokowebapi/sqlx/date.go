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

type DateOnlyImpl interface {
	MarshalText() (text []byte, err error)
	UnmarshalText(text []byte) (err error)
	MarshalJSON() (data []byte, err error)
	UnmarshalJSON(data []byte) (err error)
	Value() (driver.Value, error)
	Scan(value any) error
	String() string
	ToString() string
}

type NullDateOnly struct {
	DateOnly DateOnly
	Valid    bool
}

func NewDateOnly(value time.Time) NullDateOnly {
	return NullDateOnly{
		DateOnly: DateOnly{
			Time: value,
		},
		Valid: true,
	}
}

func (w *NullDateOnly) baseInit() {
	var dateOnly DateOnly
	if w.DateOnly != dateOnly {
		return
	}

	w.DateOnly = dateOnly
}

func (w NullDateOnly) MarshalText() (text []byte, err error) {
	if w.Valid {
		w.baseInit()
		return w.DateOnly.MarshalText()
	}

	return []byte("NULL"), nil
}

func (w *NullDateOnly) UnmarshalText(text []byte) (err error) {
	if !strings.EqualFold(string(text), "NULL") {
		w.baseInit()
		return w.DateOnly.UnmarshalText(text)
	}

	w.Valid = false
	return nil
}

func (w NullDateOnly) MarshalJSON() (data []byte, err error) {
	if w.Valid {
		w.baseInit()
		return w.DateOnly.MarshalJSON()
	}

	return []byte("null"), nil
}

func (w *NullDateOnly) UnmarshalJSON(data []byte) (err error) {
	if string(data) != "null" {
		w.baseInit()
		return w.DateOnly.UnmarshalJSON(data)
	}

	w.Valid = false
	return nil
}

func (w NullDateOnly) Value() (driver.Value, error) {
	if w.Valid {
		w.baseInit()
		return w.DateOnly.Value()
	}

	return nil, nil
}

func (w *NullDateOnly) Scan(value any) error {
	if value != nil {
		w.baseInit()
		return w.DateOnly.Scan(value)
	}

	w.Valid = false
	return nil
}

func (w NullDateOnly) String() string {
	w.baseInit()
	return w.DateOnly.String()
}

func (w *NullDateOnly) ToString() string {
	w.baseInit()
	return w.DateOnly.ToString()
}

// DateOnly custom type to store only time (hh:mm:ss)
type DateOnly struct {
	time.Time
}

func SafeParseDateOnly(value string) (NullDateOnly, error) {
	var err error
	var temp time.Time
	nokocore.KeepVoid(err, temp)

	if temp, err = time.Parse(nokocore.DateOnlyFormat, value); err != nil {
		return NullDateOnly{
			DateOnly: DateOnly{},
			Valid:    false,
		}, err
	}

	return NullDateOnly{
		DateOnly: DateOnly{
			Time: temp.UTC(),
		},
		Valid: true,
	}, nil
}

func ParseDateOnly(value string) NullDateOnly {
	return nokocore.Unwrap(SafeParseDateOnly(value))
}

func SafeParseDateOnlyNotNull(value string) (DateOnly, error) {
	var err error
	var dateOnlyNull NullDateOnly
	nokocore.KeepVoid(err, dateOnlyNull)

	if dateOnlyNull, err = SafeParseDateOnly(value); err != nil {
		return DateOnly{}, err
	}

	if !dateOnlyNull.Valid {
		return DateOnly{}, errors.New("invalid date")
	}

	return dateOnlyNull.DateOnly, nil
}

func ParseDateOnlyNotNull(value string) DateOnly {
	return nokocore.Unwrap(SafeParseDateOnly(value)).DateOnly
}

// MarshalText for text marshaling
func (w DateOnly) MarshalText() (text []byte, err error) {
	return []byte(w.Format(nokocore.DateOnlyFormat)), nil
}

// UnmarshalText for text unmarshalling
func (w *DateOnly) UnmarshalText(text []byte) (err error) {
	var temp time.Time
	nokocore.KeepVoid(temp)

	value := string(text)

	if temp, err = time.Parse(nokocore.DateOnlyFormat, value); err != nil {
		return err
	}

	w.Time = temp.UTC()
	return nil
}

// MarshalJSON for JSON marshaling
func (w DateOnly) MarshalJSON() (data []byte, err error) {
	return []byte(strconv.Quote(w.Format(nokocore.DateOnlyFormat))), nil
}

// UnmarshalJSON for JSON unmarshalling
func (w *DateOnly) UnmarshalJSON(data []byte) (err error) {
	var value string
	var temp time.Time
	nokocore.KeepVoid(value, temp)

	if value, err = strconv.Unquote(string(data)); err != nil {
		return err
	}

	if temp, err = time.Parse(nokocore.DateOnlyFormat, value); err != nil {
		return err
	}

	w.Time = temp.UTC()
	return nil
}

// Value implements the driver Valuer interface.
func (w DateOnly) Value() (driver.Value, error) {
	return w.Format(nokocore.DateOnlyFormat), nil
}

// Scan implements the Scanner interface.
func (w *DateOnly) Scan(value any) error {
	var err error
	var temp time.Time
	nokocore.KeepVoid(err, temp)

	if value != nil {
		switch val := value.(type) {
		case time.Time:
			*w = DateOnly{
				Time: val,
			}
			return nil

		case []byte:
			if temp, err = time.Parse(nokocore.DateOnlyFormat, string(val)); err != nil {
				return err
			}

			*w = DateOnly{
				Time: temp.UTC(),
			}
			return nil

		case string:
			if temp, err = time.Parse(nokocore.DateOnlyFormat, val); err != nil {
				return err
			}

			*w = DateOnly{
				Time: temp.UTC(),
			}
			return nil

		default:
			return errors.New(fmt.Sprintf("cannot convert %v to date only", value))
		}
	}

	*w = DateOnly{}
	return nil
}

func (w DateOnly) String() string {
	return w.Format(nokocore.DateOnlyFormat)
}

func (w *DateOnly) ToString() string {
	return w.String()
}
