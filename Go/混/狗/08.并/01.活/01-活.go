/* 狗程
 */
package main

import (
	"fmt"
	"time"
)

func main() { // 主狗程
	go 活(100 * time.Millisecond) // 子狗程
	const 号 = 50
	数 := 临和列(号)
	fmt.Printf("\r临和列第%d号数是%d\n", 号, 数)
}

func 活(延 time.Duration) {
	for {
		for _, 转 := range `-\|/` {
			fmt.Printf("\r%c", 转)
			time.Sleep(延)
		}
	}
}

func 临和列(n int) int {
	if n < 2 {
		return n
	}
	return 临和列(n-1) + 临和列(n-2)
}
