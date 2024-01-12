package rhtool_strings

import (
	"strconv"
	"strings"
)

func ToInt(str string) []int {
	return toIntCustomSplit(str, ",")
}

func ToIntCustomSplit(str, split string) []int {
	return toIntCustomSplit(str, split)
}

func ToUint64(str string) []uint64 {
	return toUint64CustomSplit(str, ",")
}

func ToUint64CustomSplit(str, split string) []uint64 {
	return toUint64CustomSplit(str, split)
}

func IsNil(str string) bool {
	r := ToUint64CustomSplit(str, ",")
	return len(r) == 0
}

func IsNilCustomSplit(str, split string) bool {
	r := ToUint64CustomSplit(str, split)
	return len(r) == 0
}

func String2Float(str string) (float64, error) {
	r, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return r, nil
}

// 范围替换 rand replace
func MaskByPhoneString(s string, maskStr string, start int, end int) (result string) {
	if maskStr == "" {
		maskStr = "*"
	}

	length := len(s)
	if length > 0 && length > start {
		if end >= length {
			end = length
		}
		oldStr := s[start:end]
		str := ""
		for i := 0; i < len(oldStr); i++ {
			str += maskStr
		}
		result = strings.ReplaceAll(s, oldStr, str)
	}

	return
}

func toIntCustomSplit(str, split string) []int {
	slice := make([]int, 0)
	if strings.Trim(str, " ") == "" {
		return slice
	}
	list := strings.Split(str, split)
	for _, t := range list {

		i, err := strconv.Atoi(t)
		if err != nil {
			continue
		}

		slice = append(slice, i)
	}
	return slice
}

func toUint64CustomSplit(str, split string) []uint64 {
	slice := make([]uint64, 0)
	if strings.Trim(str, " ") == "" {
		return slice
	}
	list := strings.Split(str, split)
	for _, t := range list {

		i, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			continue
		}

		slice = append(slice, i)
	}
	return slice
}
