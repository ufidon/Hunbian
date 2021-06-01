/* 狗慌了啥表现？
 */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer 看叠()
	慌(3)
}

func 慌(x int) {
	fmt.Printf("慌(%d)\n", 2*x+3/x) // 当x==0狗就慌了
	defer fmt.Printf("延迟 %d\n", x) // 注意延迟函数的执行顺序
	慌(x - 1)
}

func 看叠() {
	var 格 [4096]byte
	信长 := runtime.Stack(格[:], false)
	fmt.Println("叠始：---------------",
		string(格[:信长]),
		"叠终：---------------")
}
