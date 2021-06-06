package parser

import (
	"fmt"
	"strings"
)

var 溯层 int = 0

const 层缩占位符 string = "\t"

func 标层() string {
	return strings.Repeat(层缩占位符, 溯层-1)
}

func 溯印(溯志 string) {
	fmt.Printf("%s%s\n", 标层(), 溯志)
}

func 进溯() { 溯层 = 溯层 + 1 }
func 退溯() { 溯层 = 溯层 - 1 }

func 溯(志 string) string {
	进溯()
	溯印("开始 " + 志)
	return 志
}

func 退(志 string) {
	溯印("结束 " + 志)
	退溯()
}
