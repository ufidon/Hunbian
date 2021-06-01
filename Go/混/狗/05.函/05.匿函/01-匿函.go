/* 匿名函数与闭包
 */
package main

import "fmt"

// a_{n+1} = a_n + d
// a0
func 等差数列(a0, d int) func() int {
	an := a0 // 存活于函数之外
	return func() int {
		an += d
		return an
	}
}

// a_{n+1} = a_n * r
// a0
func 等比数列(a0, r int) func() int {
	an := a0 // 存活于函数之外

	return func() int {
		an *= r
		return an
	}
}

// a_{n+2} = a_{n+1} + a_n
// a0, a1
func 临和数列(a0, a1 int) func() int {
	前, 后 := a0, a1 // 存活于函数之外
	return func() int {
		前, 后 = 后, 后+前
		return 前
	}
}

func main() {
	fmt.Println("等差数列：")
	函 := 等差数列(0, 1)
	fmt.Println(函())
	fmt.Println(函())
	fmt.Println(函())
	fmt.Println(函())
	fmt.Println(函())

	fmt.Println("等比数列：")
	函 = 等比数列(1, 2)
	fmt.Println(函())
	fmt.Println(函())
	fmt.Println(函())
	fmt.Println(函())
	fmt.Println(函())

	fmt.Println("临和数列：")
	函 = 临和数列(0, 1)
	fmt.Println(函())
	fmt.Println(函())
	fmt.Println(函())
	fmt.Println(函())
	fmt.Println(函())
}
