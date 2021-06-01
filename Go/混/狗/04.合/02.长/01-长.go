/* 切片的增长
 */
package main

import (
	"fmt"
)

func 倍增(片 []rune, 符 rune) []rune {
	var 新片 []rune
	片长 := len(片)
	需长 := 片长 + 1

	if 需长 <= cap(片) { // 片还有空余格
		新片 = 片[:需长]
	} else { // 无空余格，造两倍容量的切片
		新片容 := 需长
		if 新片容 < 2*片长 {
			新片容 = 2 * 片长
		}
		新片 = make([]rune, 需长, 新片容)
		copy(新片, 片)
	}
	新片[片长] = 符
	return 新片
}

// 多符可以是多个字或单片
func 翻增(片 []rune, 多符 ...rune) []rune {
	var 新片 []rune
	片长 := len(片)
	需长 := 片长 + len(多符)

	if 需长 <= cap(片) { // 片还有空余格
		新片 = 片[:需长]
	} else { // 无空余格，造两倍容量的切片
		新片容 := 需长
		if 新片容 < 2*片长 {
			新片容 = 2 * 片长
		}
		新片 = make([]rune, 需长, 新片容)
		copy(新片, 片)
	}
	copy(新片[片长:], 多符)
	return 新片
}

func main() {
	var 甲, 乙 []rune

	fmt.Printf("新定义切片 容量=%d\t内容=%X\n", cap(甲), 甲)
	fmt.Printf("下面演示其倍增过程:\n\n")

	字列 := []rune("世界人民大团结！")

	for 号, 字 := range 字列 {
		甲 = 倍增(乙, 字)
		fmt.Printf("%02d 容量=%d\t切片=%X\n", 号, cap(甲), 甲)
		乙 = 甲
	}

	// 翻增：多参数倍增
	甲 = 翻增(甲, 字列...)
	fmt.Printf("增字列：\n   容量=%d\t切片=%X\n", cap(甲), 甲)

	甲 = 翻增(甲, 字列[0], 字列[2])
	fmt.Printf("增多字：\n   容量=%d\t切片=%X\n", cap(甲), 甲)
}
