/* 汇率
* 甲币对乙币汇率 = 多少甲币每乙币
 */

package main

import (
	"fmt"
	"os"
)

const 中美汇率 = 7.2

func main() {
	var 入 string
	刀, 块 := 0.0, 0.0
	var 障 error = nil

	for {
		fmt.Println("中美换汇小程序，按t退出， 按j继续")
		_, 障 = fmt.Fscanf(os.Stdin, "%s\n", &入)
		if 障 != nil {
			fmt.Println("输入错误, 请继续")
			continue
		}
		if 入 == "t" {
			os.Exit(1)
		}

		fmt.Println("您要换多少块？")
		_, 障 = fmt.Fscanf(os.Stdin, "%f\n", &块)
		if 障 != nil {
			fmt.Println("输入错误, 请继续")
			continue
		}
		刀 = 块 / 中美汇率
		fmt.Printf("%f块可以换%f刀\n", 块, 刀)
	}
}
