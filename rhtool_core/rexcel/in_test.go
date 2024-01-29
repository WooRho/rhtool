package rexcel

import (
	"fmt"
	"testing"
)

func TestIn(t *testing.T) {
	// 自定义文件
	data, err := LoadFromExcelFile(
		"狗狗.xlsx",
		Dog{},
		"狗狗",
	)

	list := make([]Dog, 0)

	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, item := range data {
		list = append(list, item.(Dog))
	}
	fmt.Println(list)
}
