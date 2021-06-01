/* 统计字、词、短语、句和行出现频率
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func 行频(文 io.Reader) map[string]int {
	行次 := make(map[string]int)
	入 := bufio.NewScanner(文)
	for 入.Scan() {
		行 := 入.Text()
		行次[行]++
	}
	return 行次
}

func 字频(文 io.Reader) (map[rune]int, int) {
	字次 := make(map[rune]int)
	入 := bufio.NewReader(文)
	无效字数 := 0
	for {
		字, 字长, 障 := 入.ReadRune()
		if 障 == io.EOF {
			break
		}
		if 障 != nil {
			fmt.Fprintf(os.Stderr, "读字障：%v\n", 障)
			return nil, 无效字数
		}
		if 字 == unicode.ReplacementChar && 字长 == 1 {
			无效字数++
			continue
		}
		字次[字]++

	}
	return 字次, 无效字数
}

func main() {
	fmt.Println("行频统计：")
	行次 := 行频(os.Stdin)

	for 行, 次 := range 行次 {
		fmt.Printf("%s出现%02d次\n", 行, 次)
	}

	fmt.Println("\n字频统计：")
	字次, _ := 字频(os.Stdin)

	for 字, 次 := range 字次 {
		fmt.Printf("%c出现%02d次\n", 字, 次)
	}
}
