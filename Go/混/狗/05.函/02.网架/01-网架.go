/* 显示网页树形结构
* 可用第一章的07-下载下载网页灌入本程序
* 例：
* ./07-下载 https://golang.org | ./01-网架
 */

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	网页, 障 := html.Parse(os.Stdin)
	if 障 != nil {
		fmt.Fprintf(os.Stderr, "网架：%v\n", 障)
		os.Exit(1)
	}
	网架(nil, 网页)
}

func 网架(叠 []string, 节点 *html.Node) {
	if 节点.Type == html.ElementNode {
		叠 = append(叠, 节点.Data)
		fmt.Println(叠) // 从根到本节点
	}
	// 传给子节点的叠只是父叠的一份复制
	for 子 := 节点.FirstChild; 子 != nil; 子 = 子.NextSibling {
		网架(叠, 子)
	}
}
