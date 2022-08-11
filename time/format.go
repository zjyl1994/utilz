package time

import (
	"errors"
	"time"
)

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

var ErrFormatNotSupport = errors.New("format not support")

func Parse(s string) (time.Time, error) {
	t, err := time.Parse(TimeFormat, s)
	if err == nil {
		return t, nil
	}
	t, err = time.Parse(DateFormat, s)
	if err == nil {
		return t, nil
	}
	return time.Time{}, ErrFormatNotSupport
}

func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

func FormatTime(t time.Time) string {
	return t.Format(TimeFormat)
}
