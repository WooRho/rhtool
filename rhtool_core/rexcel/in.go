package rexcel

import (
	"errors"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
)

// 从Excel文件中读取数据，并将其映射到指定的结构体
func LoadFromExcelFile(filepath string, v interface{}, sheet string) ([]interface{}, error) {
	typ := reflect.TypeOf(v)
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("the type must be a struct")
	}

	file, err := excelize.OpenFile(filepath)
	if err != nil {
		return nil, err
	}

	var data []interface{}

	rows, err := file.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	data = ReflectExcel(rows, v)

	return data, nil
}

func ReflectExcel(rows [][]string, v interface{}) (data []interface{}) {
	typ := reflect.TypeOf(v)
	fileNames := make(map[string]string)
	ForeachField(v, func(field reflect.StructField, value interface{}) bool {
		excelName := GetTagKey(field, "excel")
		jsonName := field.Name
		if excelName != "" {
			fileNames[excelName] = jsonName
		}
		return true
	})

	colNames := make(map[int]string)
	for i, colName := range rows[0] {
		if colName, ok := fileNames[colName]; ok {
			colNames[i] = colName
		}
	}

	for _, row := range rows[1:] {
		obj := reflect.New(typ).Elem()
		for i, cell := range row {
			fieldName, ok := colNames[i]
			if !ok {
				continue
			}

			field := obj.FieldByName(fieldName)
			if !field.IsValid() {
				continue
			}

			switch field.Type().Kind() {
			case reflect.Bool:
				fieldValue, _ := strconv.ParseBool(cell)
				field.SetBool(fieldValue)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fieldValue, _ := strconv.ParseInt(cell, 10, 64)
				field.SetInt(fieldValue)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				fieldValue, _ := strconv.ParseUint(cell, 10, 64)
				field.SetUint(fieldValue)
			case reflect.Float32, reflect.Float64:
				fieldValue, _ := strconv.ParseFloat(cell, 64)
				field.SetFloat(fieldValue)
			case reflect.String:
				field.SetString(cell)
			}
		}

		data = append(data, obj.Interface())
	}
	return data
}

func ForeachField(o interface{}, f func(field reflect.StructField, value interface{}) bool) {
	if o == nil {
		return
	}

	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	if t.Kind() == reflect.Struct {

		for i := 0; i < v.NumField(); i++ {
			tf := t.Field(i)
			vf := v.Field(i)

			// 不包括私有字段（私有字段不在递归）
			if !IsCapitalHeader(tf.Name) {
				continue
			}
			if tf.Type.Kind() == reflect.Slice {
				if tf.Type.Elem().Kind() == reflect.Ptr {
					if tf.Type.Elem().Elem().Kind() == reflect.Struct {
						vok := reflect.New(tf.Type.Elem().Elem()).Interface()
						ForeachField(vok, f)
					}
				} else if tf.Type.Kind() == reflect.Struct {
					vok := reflect.New(tf.Type.Elem()).Interface()
					ForeachField(vok, f)
				}
			} else if tf.Type.Kind() == reflect.Ptr {
				if tf.Type.Elem().Kind() == reflect.Struct {
					vok := reflect.New(tf.Type.Elem()).Interface()
					ForeachField(vok, f)
				}
			} else if tf.Type.Kind() == reflect.Struct {
				ForeachField(vf.Interface(), f)
			}

			success := f(tf, vf.Interface())
			if !success {
				return
			}
		}
	}
}

// 是否为大写开头
func IsCapitalHeader(s string) bool {
	if len(s) == 0 {
		return false
	}
	head := s[:1]
	t := []rune(head)
	if t[0] >= 65 && t[0] < 91 {
		return true
	} else {
		return false
	}
}
