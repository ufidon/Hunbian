/* 格式化输出任何类型
 */
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// 格式化任何值成串
func G格值(值 interface{}) string {
	return 格元(reflect.ValueOf(值))
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
	case reflect.Bool:
		return strconv.FormatBool(值.Bool())
	case reflect.String:
		return strconv.Quote(值.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return 值.Type().String() + " 0x" +
			strconv.FormatUint(uint64(值.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return 值.Type().String() + " 值"
	}
}

func main() {
	var 数 int32 = 32
	var 时 time.Duration = 2 * time.Microsecond
	var 值集 []interface{} = []interface{}{
		数,
		时,
		[]int32{数},
		[]time.Duration{时},
		格元,
	}

	for _, 值 := range 值集 {
		fmt.Println(G格值(值))
	}
}
