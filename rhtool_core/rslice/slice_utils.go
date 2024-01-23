package rslice

import "strconv"

// 检测集合
// union 并集
// intersection 交集
func CheckUint64Set(left, right []uint64) (justLeft, justRight, union, intersection []uint64) {
	leftSet := make(map[uint64]interface{})
	rightSet := make(map[uint64]interface{})
	unionSet := make(map[uint64]interface{})

	for i := range left {
		leftSet[left[i]] = nil
		unionSet[left[i]] = nil
	}

	for i := range right {
		rightSet[right[i]] = nil
		unionSet[right[i]] = nil
	}

	for i := range unionSet {
		_, inLeft := leftSet[i]
		_, inRight := rightSet[i]
		if inLeft && inRight {
			intersection = append(intersection, i)
		} else if inLeft {
			justLeft = append(justLeft, i)
		} else if inRight {
			justRight = append(justRight, i)
		}
		union = append(union, i)
	}
	return
}

// 检测集合
// union 并集
// intersection 交集
func CheckStringSet(left, right []string) (justLeft, justRight, union, intersection []string) {
	leftSet := make(map[string]interface{})
	rightSet := make(map[string]interface{})
	unionSet := make(map[string]interface{})

	for i := range left {
		leftSet[left[i]] = nil
		unionSet[left[i]] = nil
	}

	for i := range right {
		rightSet[right[i]] = nil
		unionSet[right[i]] = nil
	}

	for i := range unionSet {
		_, inLeft := leftSet[i]
		_, inRight := rightSet[i]
		if inLeft && inRight {
			intersection = append(intersection, i)
		} else if inLeft {
			justLeft = append(justLeft, i)
		} else if inRight {
			justRight = append(justRight, i)
		}
		union = append(union, i)
	}
	return
}

// []uint64 -> string
func UInt64s2String(value []uint64) (str string) {
	for _, v := range value {
		if len(str) == 0 {
			str = strconv.FormatUint(v, 10)
		} else {
			str = str + "," + strconv.FormatUint(v, 10)
		}
	}

	return str
}

// []string 2 []int
func StringArr2IntArr(strs []string) (res []int, err error) {
	res = make([]int, len(strs))
	for index, val := range strs {
		res[index], err = strconv.Atoi(val)
		if err != nil {
			return
		}
	}
	return
}
