package rhtool_map

import (
	"fmt"
	"testing"
)

func TestSetUnit64(t *testing.T) {
	set := NewUint64Set()
	set.Add(12)
	set.AddList([]uint64{24, 53, 64})

	// [12 24 53 64]
	fmt.Println(set.List())

	set.DeleteList([]uint64{12, 24})
	// [53 64]
	fmt.Println(set.List())

	// true
	fmt.Println(set.In(64))
	// false
	fmt.Println(set.In(24))
	set.AddList([]uint64{888, 777, 666})

	// 888,777,53,64,666
	fmt.Println(set.Merge2String(","))
	// true
	fmt.Println(set.ListIn([]uint64{888, 777, 666}))
	// 5
	fmt.Println(set.Size())

}
