/* 改自找链1，演示匿名函数
* 跑：
* ./01-找链 https://golang.org
 */
package main

import (
	"fmt"
	"os"
	"zl/qulian"
)

func main() {
	for _, 网址 := range os.Args[1:] {
		链集, 障 := qulian.H取链(网址)
		if 障 != nil {
			fmt.Fprintf(os.Stderr, "找链：%v\n", 障)
			continue
		}

		for _, 链 := range 链集 {
			fmt.Println(链)
		}
	}
}
