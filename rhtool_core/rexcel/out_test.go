package rexcel

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

type Dog struct {
	Code     string `excel:"编号"`
	Name     string `excel:"名称"`
	Color    string `excel:"颜色"`
	Birthday string `excel:"生日"`
}

func TestOut(t *testing.T) {
	list := make([]Dog, 0)
	buffer := &bytes.Buffer{}
	for i := 0; i < 10; i++ {
		list = append(list, Dog{
			Code:     "xg-" + strconv.Itoa(i+1),
			Name:     "旺财-" + strconv.Itoa(i+1),
			Color:    "白色-" + strconv.Itoa(i+1),
			Birthday: "秘密-" + strconv.Itoa(i+1),
		})
	}

	err := BufferToExcel(list, buffer, "狗狗", "狗狗")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Excel file saved successfully.")
}
