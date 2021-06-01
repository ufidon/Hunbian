/* 显复行及复次
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	复次 := make(map[string]int, 0)
	输入 := bufio.NewScanner(os.Stdin)

	// 统计每行出现次数
	for 输入.Scan() {
		复次[输入.Text()]++
	}

	// 显复行及其复次
	for 行, 次 := range 复次 {
		if 次 > 1 {
			fmt.Printf("%02d: %s\n", 次, 行)
		}
	}
}
