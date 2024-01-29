package rexcel

import (
	"bytes"
	"fmt"
	"github.com/WooRho/rhtool/rhtool_core/rfile"
	"io/ioutil"
	"os"
	"reflect"
)

// 递归获取嵌套结构体字段值
func getNestedFields(vItem reflect.Value, t reflect.Type, content []interface{}) []interface{} {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := GetTagKey(field, "excel")
		if tag != "" {
			content = append(content, vItem.Field(i).Interface())
		}

		// 检查是否是嵌套结构体
		if field.Type.Kind() == reflect.Struct {
			nestedContent := getNestedFields(vItem.Field(i), field.Type, make([]interface{}, 0))
			content = append(content, nestedContent...)
		}
	}
	return content
}

func GetTagKey(f reflect.StructField, key string) string {
	tag := f.Tag.Get(key)
	if tag == "-" {
		return ""
	}
	return tag
}

func GetXlsxItemListByList(list interface{}) [][]interface{} {
	vList := reflect.ValueOf(list)
	contents := make([][]interface{}, 0)
	for i := 0; i < vList.Len(); i++ {
		vItem := vList.Index(i)
		if vItem.Kind() == reflect.Ptr {
			vItem = vItem.Elem()
		}
		content := make([]interface{}, 0)
		tList := vItem.Type()
		for i := 0; i < tList.NumField(); i++ {
			tag := GetTagKey(tList.Field(i), "excel")
			if tag != "" {
				content = append(content, vItem.Field(i).Interface())
			}
			// 检查是否是嵌套结构体
			if tList.Field(i).Type.Kind() == reflect.Struct {
				nestedContent := getNestedFields(vItem.Field(i), tList.Field(i).Type, make([]interface{}, 0))
				content = append(content, nestedContent...)
			}
		}
		contents = append(contents, content)
	}
	return contents
}

func GetXlsxHeader(list interface{}) []interface{} {
	r := make([]interface{}, 0)
	vList := reflect.ValueOf(list)
	if vList.Kind() == reflect.Ptr {
		vList = vList.Elem()
	}
	tList := vList.Type()
	if tList.Kind() == reflect.Slice {
		item := tList.Elem()
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}
		for i := 0; i < item.NumField(); i++ {
			field := item.Field(i)
			tag := GetTagKey(field, "excel")

			if field.Type.Kind() == reflect.Struct {
				headers := getNestedHeaders(field.Type, make([]interface{}, 0))
				r = append(r, headers...)
			}

			if tag != "" {
				r = append(r, tag)
			}
		}
	}

	return r
}

// 递归获取嵌套结构体字段值
func getNestedHeaders(t reflect.Type, content []interface{}) []interface{} {

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := GetTagKey(field, "excel")

		if tag != "" {
			content = append(content, tag)
		}
		//// 检查是否是嵌套结构体
		if t.Field(i).Type.Kind() == reflect.Struct {
			nestedContent := getNestedHeaders(t.Field(i).Type, make([]interface{}, 0))
			content = append(content, nestedContent...)
		}
	}
	return content
}

func BufferToExcel(list interface{}, buffer *bytes.Buffer, fileName, sheetName string) (err error) {

	title := GetXlsxHeader(list)
	contents := GetXlsxItemListByList(list)
	_, err = createXlsxFile(buffer, title, []string{sheetName}, contents)
	if err != nil {
		return
	}

	// 假设 buffer 中已经包含了完整的 Excel 文件的二进制数据
	excelData := buffer.Bytes()

	// 创建一个新的内存临时文件，以便于读取缓冲区数据到 Excel 库
	tmpFile, err := ioutil.TempFile("", "*.xlsx")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // 清理临时文件

	// 将缓冲区中的数据写入临时文件
	_, err = tmpFile.Write(excelData)
	if err != nil {
		return fmt.Errorf("failed to write data to temp file: %v", err)
	}

	// 确保数据已刷新到磁盘并关闭文件
	err = tmpFile.Close()
	if err != nil {
		return fmt.Errorf("failed to close temp file: %v", err)
	}

	// 使用excelize等库打开临时文件进行验证或直接移动文件至目标位置
	// 移动临时文件到最终的目标路径

	err = rfile.MoveFile(tmpFile.Name(), fileName)

	//err = os.Rename(tmpFile.Name(), fileName)
	if err != nil {
		return fmt.Errorf("failed to rename temp file to final location: %v", err)
	}

	return nil
}
