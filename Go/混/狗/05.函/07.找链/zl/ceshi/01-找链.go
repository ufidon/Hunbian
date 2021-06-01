/* 改自找链1，演示图之广度优先遍历
* 跑：
* ./01-找链 https://golang.org
 */
package main

import (
	"fmt"
	"log"
	"os"
	"zl/qulian"
)

func main() {
	横历(爬, os.Args[1:])
}

// 图之广度优先遍历
func 横历(函 func(点 string) []string, 列 []string) {
	完 := make(map[string]bool)

	for len(列) > 0 {
		点集 := 列
		列 = nil
		for _, 点 := range 点集 {
			if !完[点] {
				完[点] = true
				列 = append(列, 函(点)...)
			}
		}
	}
}

func 爬(网址 string) []string {
	fmt.Println(网址)
	列, 障 := qulian.H取链(网址)
	if 障 != nil {
		log.Println(障)
	}
	return 列
}
