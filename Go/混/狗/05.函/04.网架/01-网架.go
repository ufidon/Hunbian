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

	逐节(网页, 节始, 节终)
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

func 逐节(节点 *html.Node, 预备, 善后 func(节点 *html.Node)) {
	if 预备 != nil {
		预备(节点)
	}
	for 子 := 节点.FirstChild; 子 != nil; 子 = 子.NextSibling {
		逐节(子, 预备, 善后)
	}
	if 善后 != nil {
		善后(节点)
	}
}

var 节点深度 int

func 节始(节点 *html.Node) {
	if 节点.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", 节点深度*2, "", 节点.Data)
		节点深度++
	}
}

func 节终(节点 *html.Node) {
	if 节点.Type == html.ElementNode {
		节点深度--
		fmt.Printf("%*s</%s>\n", 节点深度*2, "", 节点.Data)
	}
}
