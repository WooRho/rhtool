package test

import (
	"fmt"
	"testing"
)

type Val struct {
	Name string
}

func TestMapObj(t *testing.T) {
	var (
		m = make(map[string]*Val)
	)

	// 假设已经存在
	m["one"] = &Val{}

	val := m["one"]
	m["one"].Name = "one"
	delete(m, "one")
	fmt.Println(*val)

}
