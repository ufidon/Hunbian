/* 从路径中提取档名
 */

package main

import (
	"fmt"
	"strings"
)

// 广东/深圳/南山.风景.jpg -> 南山.风景
func 档名(串 string) string {
	// 从后往前逐字找'/'，去除第一个找到的'/'及其前面所有字符
	for i := len(串) - 1; i >= 0; i-- {
		if 串[i] == '/' {
			串 = 串[i+1:]
			break
		}
	}
	// 从后往前逐字找'.'，保留第一个找到的‘.’前所有字符
	for i := len(串) - 1; i >= 0; i-- {
		if 串[i] == '.' {
			串 = 串[:i]
			break
		}
	}
	return 串
}

func 档名2(串 string) string {
	末杠位 := strings.LastIndex(串, "/") //-1表示未找到
	串 = 串[末杠位+1:]
	if 末点位 := strings.LastIndex(串, "."); 末点位 >= 0 {
		串 = 串[:末点位]
	}
	return 串
}

func main() {
	档名集 := [...]string{
		"广东/深圳/南山.风景.jpg",
		"南山.风景.jpg",
		"深圳/南山.jpg",
	}

	for 号, 名 := range 档名集 {
		fmt.Printf("%d: %s -> %s\n", 号, 名, 档名(名))
	}

	fmt.Println("")
	for 号, 名 := range 档名集 {
		fmt.Printf("%d: %s -> %s\n", 号, 名, 档名2(名))
	}
}
