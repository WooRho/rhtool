package test

import (
	"fmt"
	"testing"
)

var pc [256]byte

func TestPlus(t *testing.T) {
	var num uint64
	num = 100
	dm := int(pc[byte(num>>1)])
	fmt.Println(dm)
}
