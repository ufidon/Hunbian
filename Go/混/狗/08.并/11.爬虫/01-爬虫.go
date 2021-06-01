/* 使用狗程并行下载的爬虫
* 用旗号限制狗程数,因为系统各种资源有限
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

// 管模拟旗号, 限制并行狗程数为20
var 旗号 = make(chan struct{}, 20)

func 爬(链 string) []string {
	fmt.Println(链)
	旗号 <- struct{}{} // 取令旗
	链列, 障 := 取链(链)
	<-旗号 // 还令旗
	if 障 != nil {
		log.Println(障)
	}
	return 链列
}

func main() {
	网址集 := make(chan []string)
	var 待送链数 int

	待送链数++

	go func() {
		网址集 <- os.Args[1:]
	}()

	爬过 := make(map[string]bool)
	for ; 待送链数 > 0; 待送链数-- {
		网址 := <-网址集
		for _, 链 := range 网址 {
			if !爬过[链] {
				爬过[链] = true
				待送链数++
				go func(链 string) {
					网址集 <- 爬(链)
				}(链)
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
