// 独立程序
package main

// 导入包
import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 屏显行参
	隔, 行参串 := "", ""
	for _, 参 := range os.Args {
		行参串 += 隔 + 参
		隔 = " "
	}
	fmt.Println("取行参法一：", 行参串)

	// 屏显程名外之行参
	fmt.Println("取行参法二：", strings.Join(os.Args[1:], " "))
	fmt.Println("取行参法三：", os.Args[1:])
}
