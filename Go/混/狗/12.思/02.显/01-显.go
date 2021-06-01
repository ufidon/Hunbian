/* 显示结构化数据
 */
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func X显(名 string, 量 interface{}) {
	fmt.Printf("显示 %s (%T):\n", 名, 量)
	显(名, reflect.ValueOf(量))
}

// 把值格式化成串，不查其内部结构
func 格元(值 reflect.Value) string {
	switch 值.Kind() {
	case reflect.Invalid:
		return "无效值"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(值.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(值.Uint(), 10)
	// ...其余数型略...
	case reflect.Bool:
		if 值.Bool() {
			return "真"
		}
		return "假"
	case reflect.String:
		return strconv.Quote(值.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return 值.Type().String() + " 0x" +
			strconv.FormatUint(uint64(值.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return 值.Type().String() + " 值"
	}
}

// 显示结构数
func 显(径 string, 值 reflect.Value) {
	switch 值.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = 无效径\n", 径)
	case reflect.Slice, reflect.Array:
		for i := 0; i < 值.Len(); i++ {
			显(fmt.Sprintf("%s[%d]", 径, i), 值.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < 值.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", 径, 值.Type().Field(i).Name)
			显(fieldPath, 值.Field(i))
		}
	case reflect.Map:
		for _, key := range 值.MapKeys() {
			显(fmt.Sprintf("%s[%s]", 径,
				格元(key)), 值.MapIndex(key))
		}
	case reflect.Ptr:
		if 值.IsNil() {
			fmt.Printf("%s = 空\n", 径)
		} else {
			显(fmt.Sprintf("(*%s)", 径), 值.Elem())
		}
	case reflect.Interface:
		if 值.IsNil() {
			fmt.Printf("%s = 空\n", 径)
		} else {
			fmt.Printf("%s.type = %s\n", 径, 值.Elem().Type())
			显(径+".value", 值.Elem())
		}
	default: // 基型，管道，函
		fmt.Printf("%s = %s\n", 径, 格元(值))
	}
}

func X法(量 interface{}) {
	值 := reflect.ValueOf(量)
	型 := 值.Type()
	fmt.Printf("型: %s\n", 型)

	for i := 0; i < 值.NumMethod(); i++ {
		法型 := 值.Method(i).Type()
		fmt.Printf("法 (%s) %s%s\n", 型, 型.Method(i).Name,
			strings.TrimPrefix(法型.String(), "函"))
	}
}

func main() {
	var i interface{} = 3
	type 圈 struct {
		头 int
		尾 *圈
	}

	var 圈1 圈
	圈1 = 圈{12, &圈1}
	时 := time.Duration(1 * time.Second)

	X显("i", i)
	X显("&i", &i)
	//X显("圈1", 圈1)
	X法(i)
	X法(时)
}
