package rtime

import (
	"strconv"
	"time"
)

const (
	BackTimeFormatYMD      = 1 // 只获取时间
	BackTimeFormatYMDHMS   = 2 // 获取年月日
	BackTimeFormatLastTime = 3 // 获取当天到最后一秒
)

// 字符串转go time ----- string to time
// 支持格式 support format
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
}

// 获取最后一秒 get last time
// 支持格式 support format
// 2006-01-02
// 2006-01-02 15:04:05
func Str2StrLastTime(strTime string) string {
	return strTimeTransfer(strTime, BackTimeFormatLastTime)
}

// 只截取日期
// get yyyy-dd-mm
func Str2GetYMD(strTime string) string {
	return strTimeTransfer(strTime, BackTimeFormatYMD)
}

// 是否是0时
// is zero time
func IsStrTimeZero(strTime string) bool {
	slen := len(strTime)
	if slen == 0 {
		return true
	}
	switch slen {
	case 10:
		_, err := time.ParseInLocation("2006-01-02", strTime, time.Local)
		return err != nil
	case 19:
		_, err := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local)
		return err != nil
	default:
		return false
	}
}

// to unix
func ToUnix(strTime string) int64 {
	return Str2Time(strTime).Unix()
}

// unixString 2 time
func UnixString2Time(strTime string) (time.Time, error) {
	unix, err := strconv.ParseInt(strTime, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(unix, 0), nil
}

func String2Duration(strTime string) (t time.Duration, err error) {
	r, err := strconv.Atoi(strTime)
	if err != nil {
		return 0, err
	}
	t = time.Duration(r) * time.Second

	return t, nil
}

func strTimeTransfer(strTime string, backType int) string {
	if IsStrTimeZero(strTime) {
		return ""
	}
	switch backType {
	// 获取年月日
	case BackTimeFormatYMD:
		t, _ := time.ParseInLocation("2006-01-02", strTime, time.Local)
		return t.Format("2006-01-02")
	// 获取年月日时分秒
	case BackTimeFormatYMDHMS:
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local)
		return t.Format("2006-01-02 15:04:05")
	// 获取当天最后一秒-年月日时分秒
	case BackTimeFormatLastTime:
		if len(strTime) == 10 {
			t, _ := time.ParseInLocation("2006-01-02", strTime, time.Local)
			return t.Format("2006-01-02") + " 23:59:59"
		}
		if len(strTime) == 19 {
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local)
			return t.Format("2006-01-02") + " 23:59:59"
		}
		return ""
	}
	return ""
}
