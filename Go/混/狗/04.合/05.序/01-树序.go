/* 用二叉树实现插入排序
 */

package main

import (
	"fmt"
)

type 二叉树 struct {
	值    int
	左, 右 *二叉树
}

// 用二叉树对数列做在位排序
func 排序(数列 []int) {
	var 根 *二叉树
	for _, 数 := range 数列 {
		根 = 添(根, 数)
	}
	附值(数列[:0], 根)
}

// 附值依序推值入列并返回序列
func 附值(数列 []int, 树 *二叉树) []int {
	if 树 != nil {
		数列 = 附值(数列, 树.左)
		数列 = append(数列, 树.值)
		数列 = 附值(数列, 树.右)
	}
	return 数列
}

func 添(树 *二叉树, 值 int) *二叉树 {
	if 树 == nil {
		树 = new(二叉树)
		树.值 = 值
		return 树
	}
	if 值 < 树.值 {
		树.左 = 添(树.左, 值)
	} else {
		树.右 = 添(树.右, 值)
	}
	return 树
}

func main() {
	数列 := []int{99, 44, 32, 76, 54, 1, 29, 83, 25}

	fmt.Println("二叉树排序：")
	fmt.Printf("排序前：%v\n", 数列)
	排序(数列)
	fmt.Printf("排序后：%v\n", 数列)
}
