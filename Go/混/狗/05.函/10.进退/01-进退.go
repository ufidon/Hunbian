/* 展示延迟在互反操作中的应用
 */

package main

import (
	"log"
	"time"
)

func main() {
	长操作()
}

func 长操作() {
	defer 测时("长操作")() // 延迟的是测时返回的函数
	time.Sleep(10 * time.Second)
}

func 测时(志 string) func() {
	起 := time.Now()
	log.Printf("开始%s\n", 志)
	return func() {
		log.Printf("完成%s于%s之后\n", 志, time.Since(起))
	}
}
