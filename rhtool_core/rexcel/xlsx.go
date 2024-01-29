package rexcel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"strconv"
)

// xlsx封装
type Xlsx struct {
	file       *excelize.File // 操作对象
	row        map[int]int    // sheet目前已使用的最大行
	sheet      map[string]int // sheet 名字映射
	workSheetI int
	workSheetS string
	name       string
}

// 新建一个文档对象
func NewXlsx(filename, sheetName string) Xlsx {
	x := Xlsx{
		file: excelize.NewFile(),
		name: filename,
	}
	x.workSheetI, _ = x.file.NewSheet(sheetName)
	x.row = map[int]int{x.workSheetI: 1}
	x.sheet = map[string]int{sheetName: x.workSheetI}
	x.workSheetS = sheetName
	x.file.DeleteSheet("Sheet1")
	return x
}

func (x *Xlsx) AddHeader(v ...interface{}) {
	row, ok := x.row[x.workSheetI]
	if !ok {
		row = 0
	}
	for i := range v {
		axis := getAxis(i, row)

		value := fmt.Sprintf("%v", v[i])
		x.file.SetColWidth(x.workSheetS, axis, axis, 1.5*float64(len(value)))
		x.file.SetCellValue(x.workSheetS, axis, value)

		//switch v[i].(type) {
		//case tools.Cent:
		//	x.SetCellStyle(i, x.NewChinaCurrencyStyle())
		//}
	}
	x.row[x.workSheetI] = row + 1
}

func createXlsxFile(w io.Writer, header []interface{}, sheetNames []string, contents [][]interface{}) (int64, error) {
	uid := "uuid"
	file := fmt.Sprintf("%s.xlsx", uid)
	// TODO 目前只取索引第0位 后期优化多表写入
	if len(sheetNames) > 0 {
		if sheetNames[0] == "" {
			sheetNames = append(sheetNames, "sheet")
		}
	}
	if len(sheetNames) == 0 {
		sheetNames = append(sheetNames, "sheet")
	}
	xlsx := NewXlsx(file, sheetNames[0])
	// xlsx.file.NewSheet("sheet2")
	xlsx.AddHeader(header...)
	for _, content := range contents {
		xlsx.AddRowEnd(content...)
	}

	return xlsx.WriteTo(w)
}

// 行列计算转换
func getAxis(col, row int) string {
	axis := ""
	base := col
	mod := 0
	for {
		mod = base % 26
		base = base / 26
		axis = string(byte('A')+byte(mod)) + axis
		if base <= 26 {
			if base != 0 {
				axis = string(byte('A')+byte(base-1)) + axis
			}
			break
		}
	}
	if row >= 0 {
		axis += strconv.Itoa(row)
	}
	return axis
}

// 设置当前单元格的格式
func (x *Xlsx) SetCellStyle(col int, style int) {
	row, ok := x.row[x.workSheetI]
	if !ok {
		row = 0
	}
	axis := getAxis(col, row)
	x.file.SetCellStyle(x.workSheetS, axis, axis, style)
}

// 向工作表添加一行数据
// 目前由于设计与使用场景原因，无法并发对一个文档对象的不同工作表进行写入
// 并发对不同工作表安全写入
func (x *Xlsx) AddRowEnd(v ...interface{}) {
	row, ok := x.row[x.workSheetI]
	if !ok {
		row = 0
	}
	for i := range v {
		axis := getAxis(i, row)

		//switch spv := v[i].(type) {
		//case tools.Cent:
		//	x.SetCellStyle(i, x.NewChinaCurrencyStyle())
		//	x.file.SetCellValue(x.workSheetS, axis, spv.ToYuan())
		//case tools.Gram:
		//	x.file.SetCellValue(x.workSheetS, axis, spv.ToKiloGram())
		//case tools.KiloGram:
		//	x.file.SetCellValue(x.workSheetS, axis, spv)
		//case tools.CentiMeterPerKiloGram:
		//	x.file.SetCellValue(x.workSheetS, axis, spv.ToMeterPerKiloGram())
		//default:
		x.file.SetCellValue(x.workSheetS, axis, v[i])
		//}
	}
	x.row[x.workSheetI] = row + 1
}

func (x *Xlsx) WriteTo(w io.Writer) (int64, error) {
	return x.file.WriteTo(w)
}
