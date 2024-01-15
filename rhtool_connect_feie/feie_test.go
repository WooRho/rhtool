package printer

import (
	"fmt"
	"testing"
)

func TestFeie(t *testing.T) {
	printer := NewFeiePrinter("jjj", "fff", "https://api.feieyun.cn/Api/Open/")
	fmt.Println(printer.AddPrinter("dasffasdf"))
	fmt.Println(printer.Delete("dsfa"))
	fmt.Println(printer.PrinterStatus("33"))
}
