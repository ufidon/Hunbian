/* 在位更新
 */
package main

import (
	"fmt"
)

// 非空返回仅含非空串的切片，切片本身被更改
func 非空(串列 []string) []string {
	非空数 := 0
	for _, 串 := range 串列 {
		if 串 != "" {
			串列[非空数] = 串
			非空数++
		}
	}
	return 串列[:非空数]
}

func 不空(串列 []string) []string {
	不空串 := 串列[:0]
	for _, 串 := range 串列 {
		if 串 != "" {
			不空串 = append(不空串, 串)
		}
	}
	return 不空串
}

func main() {
	列 := []string{"香港", "", "台湾宝岛"}
	fmt.Printf("原切片：%q\n", 列)
	fmt.Printf("仅含非空串之切片：%q\n", 非空(列))
	fmt.Printf("处理后原切片：%q\n", 列)

	列 = []string{"", "澳门", "台湾宝岛"}
	fmt.Printf("原切片：%q\n", 列)
	fmt.Printf("仅含非空串之切片：%q\n", 不空(列))
	fmt.Printf("处理后原切片：%q\n", 列)
}
