/* 狗慌了啥表现？
 */

package main

import (
	"fmt"
)

func main() {
	慌(3)
}

func 慌(x int) {
	fmt.Printf("慌(%d)\n", 2*x+3/x) // 当x==0狗就慌了
	defer fmt.Printf("延迟 %d\n", x) // 注意延迟函数的执行顺序
	慌(x - 1)
}
