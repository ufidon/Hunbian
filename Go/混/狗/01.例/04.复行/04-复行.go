/* 显飙入(标入)及档中的复行和复次
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	复次 := make(map[string]int, 0)
	档集 := os.Args[1:]

	// 统计复行及复次
	if len(档集) == 0 {
		计行(os.Stdin, 复次)
	} else {
		for _, 档 := range 档集 {
			丿, 障 := os.Open(档)
			if 障 != nil {
				fmt.Fprintf(os.Stderr, "04-复行: %v\n", 障)
				continue
			}
			计行(丿, 复次)
		}
	}
	// 显复行及复次
	for 行, 次 := range 复次 {
		if 次 > 1 {
			fmt.Printf("%02d: %s\n", 次, 行)
		}
	}

}

func 计行(丿 *os.File, 复次 map[string]int) {
	入 := bufio.NewScanner(丿)
	for 入.Scan() {
		复次[入.Text()]++
	}
	// 待做: 入.Err()障处理
}
