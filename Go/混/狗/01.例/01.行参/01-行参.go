// 独立程序
package main

// 导入包
import (
	"fmt"
	"os"
)

func main() {

	// 屏显第〇个行参
	fmt.Println("第一个行参是程序名:", os.Args[0])

	// 屏显其余行参
	var 隔, 行参串 string
	for 丶 := 1; 丶 < len(os.Args); 丶++ {
		行参串 += 隔 + os.Args[丶]
		隔 = " "
	}
	fmt.Printf("其余行参:\n%s\n", 行参串)
}
