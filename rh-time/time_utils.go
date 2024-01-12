package rh_time

import (
	"time"
)

// 支持格式
// 2006-01-02
// 2006-01-02 15:04:05
func Str2Time(strTime string) time.Time {
	var (
		t    = time.Time{}
		sLen = len(strTime)
	)
	switch sLen {
	case 10:
		t, _ = time.ParseInLocation("2006-01-02", strTime, time.Local)
		return t
	case 19:
		t, _ = time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local)
		return t
	default:
		return t
	}
	return t
}
