package parse

import (
	"strings"
	"time"
)

func NowUnix() int64 {
	return time.Now().Unix()
}

func NowDate() string {
	return time.Now().String()
}

func Date2Unix(date, format string) int64 {
	t, _ := time.Parse(format, date)
	return t.Unix()
}

func Unix2Date(t int64, format string) string {
	return time.Unix(t, 0).Format(format)
}

func Weekday(date string, format string) string {
	t, _ := time.Parse(format, date)
	return t.Weekday().String()
}

func DateTransferFormat(date, preformat, format string) string {
	t, _ := time.Parse(preformat, date)
	return t.Format(format)
}

func UTCDate2LocalDate(date, format string) string {
	t, _ := time.ParseInLocation(format, date, time.UTC)
	return time.Unix(t.Unix(), 0).Format(format)
}
