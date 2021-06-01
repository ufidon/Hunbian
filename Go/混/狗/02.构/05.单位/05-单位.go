/* 用域中，物理量类型用单位标识
* 诚然，一物理量可有多个单位
 */
package main

import (
	"fmt"
)

type 米 float64

func (长度 米) String() string {
	return fmt.Sprintf("%f米", 长度)
}

type 公斤 float64

func (重量 公斤) String() string {
	return fmt.Sprintf("%f公斤", 重量)
}

type 高重指 float64

func (指 高重指) String() string {
	return fmt.Sprintf("%f公斤每平方米", 指)
}

// 高重指数 = 重/高^2
func 高重指数(身高 米, 体重 公斤) 高重指 {
	return 高重指(float64(体重) / float64(身高) / float64(身高))
	//return 体重 / 身高 / 身高
}

func main() {
	var 身高 米 = 1.72
	var 体重 公斤 = 80.0

	var 指 = 高重指数(身高, 体重)

	fmt.Println(指.String())
	fmt.Printf("%v\n", 指)
	fmt.Printf("%s\n", 指)
	fmt.Println(指)
	fmt.Printf("%f\n", 指)
	fmt.Println(float64(指))
}
