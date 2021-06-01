// 独立程序
package main

// 导入包
import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("用法：本程序名 客名")
		return
	}

	fmt.Println(os.Args[1], "您好， 欢迎来宠狗！")
	fmt.Printf("%s您好，欢迎来宠狗！\n", os.Args[1])
}
