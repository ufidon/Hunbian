/* 数字分组与显示
 */
package main

import (
	"bytes"
	"fmt"
)

// 按英文习惯分组，3个一组
func 数字分组(串, 隔 string) string {
	长 := len(串)
	if 长 <= 3 {
		return 串
	}

	return 数字分组(串[:长-3], 隔) + 隔 + 串[长-3:]
}

func 数列转串(数列 []int, 隔 string) string {
	var 串 bytes.Buffer

	串.WriteRune('【')
	for 号, 数 := range 数列 {
		if 号 > 0 {
			串.WriteString(隔)
		}
		fmt.Fprintf(&串, "%d", 数)
	}
	串.WriteRune('】')

	return 串.String()
}

func main() {
	数串集 := []string{
		"123",
		"1234567",
		"12",
		"123456789",
	}

	for _, 数 := range 数串集 {
		fmt.Printf("%s -> %s\n", 数, 数字分组(数, "_"))
	}

	数列 := []int{
		123,
		1234567,
		12,
		123456789,
	}
	fmt.Printf("\n\n%s\n", 数列转串(数列, "，"))
}
