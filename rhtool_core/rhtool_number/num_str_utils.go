package rhtool_number

import (
	"fmt"
	"strconv"
)

// number operate
func UInt642String(num uint64) string {
	r := strconv.FormatUint(num, 10)
	return r
}

func Int2String(num int) string {
	r := strconv.Itoa(num)
	return r
}

func Float2String(num float64) string {
	return fmt.Sprintf("%f", num)
}
