/* 一次赋值多个变量
 */
package main

import (
	"fmt"
)

func 交换(甲, 乙 *int) {
	*甲, *乙 = *乙, *甲
}

func 最大公约数(甲, 乙 int) int {
	for 乙 != 0 {
		甲, 乙 = 乙, 甲%乙
	}
	return 甲
}

/*a_{n} = a_{n-1} + a_{n-2}
* a_0 = 0, a_1 = 1
 */

func 临加列(号 int) int {
	前, 后 := 0, 1
	for 氵 := 0; 氵 < 号; 氵++ {
		前, 后 = 后, 前+后
	}
	return 前
}

func main() {
	甲, 乙 := 100, 75

	fmt.Printf("甲乙交换前：甲=%d 乙=%d\n", 甲, 乙)
	交换(&甲, &乙)
	fmt.Printf("甲乙交换后：甲=%d 乙=%d\n", 甲, 乙)
	公 := 最大公约数(甲, 乙)
	fmt.Printf("甲乙最大公约数：%d\n", 公)
	fmt.Printf("临加列第%d号是%d\n", 5, 临加列(5))
	fmt.Printf("临加列第%d号是%d\n", 25, 临加列(25))
}
