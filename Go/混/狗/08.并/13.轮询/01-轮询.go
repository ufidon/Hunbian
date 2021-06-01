/* 用select轮询事件
 */

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	取消 := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		取消 <- struct{}{}
	}()

	fmt.Println("导弹发射倒计时。按任意键取消！")
	嘀嗒 := time.NewTicker(1 * time.Second)
	defer 嘀嗒.Stop()

	for 倒计时 := 10; 倒计时 > 0; 倒计时-- {
		fmt.Printf("%d秒\n", 倒计时)
		select {
		case 时 := <-嘀嗒.C:
			// 继续倒计时
			fmt.Println(时.Clock())
		case <-取消:
			fmt.Println("导弹发射取消！")
			return
		}
	}

	fmt.Println("发射！")
}
