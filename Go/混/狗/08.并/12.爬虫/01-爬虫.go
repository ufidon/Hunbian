/* 使用狗程并行下载的爬虫
* 用有限狗程,因为系统各种资源有限
* 造程：
* go mod init pc
* go mod tidy
* go build 01-爬虫.go
* 跑：
* ./01-爬虫 https://golang.org
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 并行爬虫数为20
const 爬虫数 int = 20

func 爬(链 string) []string {
	fmt.Println(链)
	链列, 障 := 取链(链)
	if 障 != nil {
		log.Println(障)
	}
	return 链列
}

func main() {
	网址集 := make(chan []string) // 可能含重复链
	未爬链集 := make(chan string)  // 无重复链

	go func() {
		网址集 <- os.Args[1:] //起始网址集从命令行获取
	}()

	// 造有限狗程
	for i := 0; i < 爬虫数; i++ {
		go func() {
			for 链 := range 未爬链集 {
				爬得链集 := 爬(链)
				go func() {
					网址集 <- 爬得链集
				}()
			}
		}()
	}

	// 主狗程去除重复链，将未爬链送给爬程
	爬过链集 := make(map[string]bool)
	for 链列 := range 网址集 {
		for _, 链 := range 链列 {
			if !爬过链集[链] {
				爬过链集[链] = true
				未爬链集 <- 链
			}
		}
	}
}

func 取链(网址 string) ([]string, error) {
	回, 障 := http.Get(网址)

	if 障 != nil {
		return nil, 障
	}

	if 回.StatusCode != http.StatusOK {
		回.Body.Close()
		return nil, fmt.Errorf("取%s障%s", 网址, 回.Status)
	}

	网页, 障 := html.Parse(回.Body)
	回.Body.Close()

	if 障 != nil {
		return nil, fmt.Errorf("析%s障%v", 网址, 障)
	}

	var 链集 []string
	访节 := func(节 *html.Node) { // 匿函
		if 节.Type == html.ElementNode && 节.Data == "a" {
			for _, 锚 := range 节.Attr {
				if 锚.Key != "href" {
					continue
				}
				链, 障 := 回.Request.URL.Parse(锚.Val) // 绝对链址
				if 障 != nil {
					continue
				}
				链集 = append(链集, 链.String())
			}
		}
	}

	逐节(网页, 访节, nil)

	return 链集, nil
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
