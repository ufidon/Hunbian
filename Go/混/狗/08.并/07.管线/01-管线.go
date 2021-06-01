/* 管道拼接成管线，流水线
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	const 总数 int = 10

	整数管 := make(chan int, 0)
	平方管 := make(chan int, 0)

	// 生成自然数
	go func() {
		for i := 0; ; i++ {
			整数管 <- i
			time.Sleep(1 * time.Second)
			if i >= 总数 {
				break
			}
		}
		close(整数管)
	}()

	// 把入数平方并输出
	go func() {
		for {
			i, 收到 := <-整数管
			if !收到 {
				break
			}
			平方管 <- i * i
		}
		close(平方管)
	}()

	// 主狗程
	for {
		平方数, 收到 := <-平方管
		if !收到 {
			return
		}
		fmt.Println(平方数)
	}
}
