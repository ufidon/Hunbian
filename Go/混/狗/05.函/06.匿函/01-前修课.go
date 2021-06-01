/* 拓扑排序：列出每门课的前修课，形成有向无环图
 */

package main

import (
	"fmt"
	"sort"
)

var 前修课集 = map[string][]string{
	"算法":  {"数据结构", "微积分"},
	"微积分": {"线性代数"},
	"编译原理": {
		"数据结构",
		"计算理论",
		"计算机架构",
	},
	"数据结构": {"离散数学"},
	"数据库":  {"数据结构"},
	"离散数学": {"编程入门"},
	"计算理论": {"离散数学"},
	"网络":   {"操作系统", "编程入门"},
	"操作系统": {"数据结构", "计算机架构"},
	"编程语言": {"数据结构", "计算机架构", "编程入门"},
}

func main() {
	fmt.Println("计算机系修课顺序：")
	for 号, 课 := range 排课(前修课集) {
		fmt.Printf("%d:\t%s\n", 号+1, 课)
	}
}

// 从前修课集排出全部课程修课顺序
// 待善：列出并修课
func 排课(前修课集 map[string][]string) []string {
	var 全课序 []string
	排好 := make(map[string]bool)

	var 遍访 func(课集 []string)

	遍访 = func(课集 []string) { // 匿函修改外部变量
		for _, 课 := range 课集 {
			if !排好[课] {
				排好[课] = true
				遍访(前修课集[课])
				全课序 = append(全课序, 课)
			}
		}
	}

	//
	var 课集 []string
	for 课 := range 前修课集 {
		课集 = append(课集, 课)
	}
	sort.Strings(课集)
	遍访(课集)

	return 全课序
}
