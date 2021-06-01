/* 延迟函数访问父函数中的变量
* 使用延迟函数需注意父函数退出操作实际执行时刻
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	角 := math.Pi / 4
	正弦(角)

	结果 := 半角正弦(角)
	fmt.Printf("外：sin(%g)=%g\n", 角, 结果)
}

// 延迟函数访问父函数中的变量
func 正弦(角度 float64) (结果 float64) {
	defer func() {
		fmt.Printf("sin(%g)=%g\n", 角度, 结果)
	}() // 注意勿漏这对圆括号
	return math.Sin(角度)
}

func 半角正弦(角度 float64) (结果 float64) {
	defer func() { // 延迟函数在父函数返回后才执行
		角度 /= 2
		内结果 := math.Sin(角度)
		fmt.Printf("内：sin(%g)=%g\n", 角度, 结果)
		fmt.Printf("内：sin(%g)=%g\n", 角度, 内结果)
	}()
	return math.Sin(角度)
}
