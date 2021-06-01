/* 找出网页中所有链接
* 可用第一章的07-下载下载网页灌入本程序
* 例：
* ./07-下载 https://golang.org | ./01-找链
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
		fmt.Fprintf(os.Stderr, "找链：%v\n", 障)
		os.Exit(1)
	}

	for _, 链集 := range 找链(nil, 网页) {
		fmt.Println(链集)
	}
}

// 把节点中的链接放入链集并返回之
func 找链(链集 []string, 节点 *html.Node) []string {
	if 节点.Type == html.ElementNode && 节点.Data == "a" {
		for _, 锚点 := range 节点.Attr {
			if 锚点.Key == "href" {
				链集 = append(链集, 锚点.Val)
			}
		}
	}
	// 递归查看节点之所有子节点
	for 子 := 节点.FirstChild; 子 != nil; 子 = 子.NextSibling {
		链集 = 找链(链集, 子)
	}

	return 链集
}
