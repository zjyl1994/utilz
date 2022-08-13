package utilz

import (
	"errors"
	"time"
)

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

var ErrFormatNotSupport = errors.New("format not support")

func Now() string {
	return time.Now().Format(TimeFormat)
}

func Today() string {
	return time.Now().Format(DateFormat)
}

func Yesterday() string {
	return time.Now().AddDate(0, 0, -1).Format(DateFormat)
}

func Time() int64 {
	return time.Now().Unix()
}

func ParseTime(s string) (time.Time, error) {
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
