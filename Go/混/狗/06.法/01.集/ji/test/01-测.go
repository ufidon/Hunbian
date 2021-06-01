/* 测试位图实现的小非负整数集
 */

package main

import (
	"fmt"
	"xj/xiaoji"
)

func main() {
	var 甲, 乙 xiaoji.Z非负整集
	甲.Z添(2)
	甲.Z添(11)
	甲.Z添(131)

	fmt.Println("甲：", 甲.String())

	乙.Z添(127)
	乙.Z添(91)
	fmt.Println("乙：", 乙.String())

	甲.Z并(&乙)
	fmt.Println("甲并乙：", 甲.String())

	fmt.Println("10 ∈ 甲：", 甲.Z有(10), "；", "91 ∈ 乙：", 乙.Z有(91))
	fmt.Println("&甲：", &甲)
	fmt.Printf("甲内部字列(十六进制)：\n%X\n", 甲)
	fmt.Printf("甲内部字列(二进制)：\n%b\n", 甲)
}
