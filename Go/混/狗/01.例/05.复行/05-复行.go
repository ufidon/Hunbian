/* 显飙入(标入)及档中的复行和复次
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	复次 := make(map[string]int, 0)
	档集 := os.Args[1:]

	// 统计复行及复次
	for _, 档 := range 档集 {
		文, 障 := ioutil.ReadFile(档)
		if 障 != nil {
			fmt.Fprintf(os.Stderr, "05-复行: %v\n", 障)
			continue
		}
		for _, 行 := range strings.Split(string(文), "\n") {
			复次[行]++
		}
	}

	// 显复行及复次
	for 行, 次 := range 复次 {
		if 次 > 1 {
			fmt.Printf("%02d: %s\n", 次, 行)
		}
	}

}
