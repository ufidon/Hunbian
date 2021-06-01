/* 管道拼接成管线，流水线
* 遍历管线元素
 */
package main

import (
	"fmt"
	"time"
)

const 总数 int = 10

func main() {

	整数管 := make(chan int, 0)
	平方管 := make(chan int, 0)

	// 生成自然数
	go func() {
		for i := 0; i < 总数; i++ {
			整数管 <- i
			time.Sleep(1 * time.Second)
		}
		close(整数管)
	}()

	// 把入数平方并输出
	go func() {
		for i := range 整数管 {
			平方管 <- i * i
		}
		close(平方管)
	}()

	// 主狗程
	for i := range 平方管 {
		fmt.Println(i)
	}
}
