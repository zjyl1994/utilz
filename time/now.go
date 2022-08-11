package time

import "time"

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
