/* 管道拼接成管线，流水线
* 遍历管线元素
 */
package main

import (
	"fmt"
	"time"
)

func 造数(出管 chan<- int, 总数 int) {
	for i := 0; i < 总数; i++ {
		出管 <- i
		time.Sleep(1 * time.Second)
	}
	close(出管)
}

func 平方(出管 chan<- int, 入管 <-chan int) {
	for i := range 入管 {
		出管 <- i * i
	}
}

func 显示(入管 <-chan int) {
	for i := range 入管 {
		fmt.Println(i)
	}
}

func main() {
	const 总数 int = 10

	整数管 := make(chan int, 0)
	平方管 := make(chan int, 0)

	go 造数(整数管, 总数)
	go 平方(平方管, 整数管)
	go 显示(平方管)

}
