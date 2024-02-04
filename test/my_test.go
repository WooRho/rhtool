package test

import (
	"fmt"
	"testing"
)

type Currency int

const (
	USD Currency = iota + 1
	EUR
	GBP
	RMB
)

func TestPlus(t *testing.T) {
	sysbom := [...]string{USD: "$", EUR: "jj", GBP: "dd", RMB: "dfe"}
	fmt.Println(sysbom)
}
