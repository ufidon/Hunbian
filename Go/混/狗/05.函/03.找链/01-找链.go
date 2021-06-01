/* 改自找链1，演示函数返回多值
* 跑：
* ./01-找链 https://golang.org
 */
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, 网址 := range os.Args[1:] {
		链集, 障 := 找链(网址)
		if 障 != nil {
			fmt.Fprintf(os.Stderr, "找链：%v\n", 障)
			continue
		}

		for _, 链 := range 链集 {
			fmt.Println(链)
		}
	}
}

func 找链(网址 string) ([]string, error) {
	回, 障 := http.Get(网址)
	if 障 != nil {
		return nil, 障
	}

	if 回.StatusCode != http.StatusOK {
		回.Body.Close()
		return nil, fmt.Errorf("结果 %s: %s", 网址, 回.Status)
	}
	网页, 障 := html.Parse(回.Body)
	回.Body.Close()

	if 障 != nil {
		return nil, fmt.Errorf("把%s作网页分析出错：%v", 网址, 障)
	}
	return 访问(nil, 网页), nil
}

func 访问(链集 []string, 节点 *html.Node) []string {
	if 节点.Type == html.ElementNode && 节点.Data == "a" {
		for _, 锚点 := range 节点.Attr {
			if 锚点.Key == "href" {
				链集 = append(链集, 锚点.Val)
			}
		}
	}
	// 递归查看节点之所有子节点
	for 子 := 节点.FirstChild; 子 != nil; 子 = 子.NextSibling {
		链集 = 访问(链集, 子)
	}

	return 链集
}
