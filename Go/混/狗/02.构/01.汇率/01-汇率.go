/* 汇率
* 甲币对乙币汇率 = 多少甲币每乙币
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"unsafe"
)

const 中美汇率 = 7.2

func main() {
	var 入 string
	刀, 块 := 0.0, 0.0
	var 障 error = nil

	for {
		_, 障 = 输入("\n中美换汇小程序，按t退出， 按j继续", unsafe.Pointer(&入), "串")
		if 障 != nil {
			fmt.Println("输入错误, 请继续")
			continue
		}
		if 入 == "t" {
			os.Exit(1)
		}

		_, 障 = 输入("您要换多少块？", unsafe.Pointer(&块), "数")
		if 障 != nil {
			fmt.Println("输入错误, 请继续")
			continue
		}
		刀 = 块换刀(块, 中美汇率)
		fmt.Printf("%f块可以换%f刀\n", 块, 刀)
	}
}

func 输入(提示 string, 签 unsafe.Pointer, 签型 string) (节数 int, 障 error) {

	fmt.Println(提示)

	switch 签型 {
	case "串":
		return fmt.Fscanf(os.Stdin, "%v\n", (*string)(签))
	case "数":
		return fmt.Fscanf(os.Stdin, "%v\n", (*float64)(签))
	default:
		return 0, errors.New("类型错误")
	}
}

func 块换刀(块 float64, 汇率 float64) float64 {
	return 块 / 汇率
}
