package thtool_time

import (
	"strconv"
	"time"
)

// is zero
func IsTimeZero(t time.Time) bool {
	if t.Unix() <= 0 {
		return true
	}
	return false
}

// time to string
func Date(t time.Time) string {
	if IsTimeZero(t) {
		return ""
	}
	return t.Format("2006-01-02")
}
func Time(t time.Time) string {
	if IsTimeZero(t) {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func Time2UnixString(t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}
