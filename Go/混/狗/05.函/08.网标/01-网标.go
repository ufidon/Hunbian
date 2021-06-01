/* 改自找链1，演示判断网页页型
* 跑：
* ./01-找链 https://golang.org
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	for _, 网址 := range os.Args[1:] {
		if 障 := 网标(网址); 障 != nil {
			log.Fatalln(障)
		}
	}

}

func 网标(网址 string) error {
	回, 障 := http.Get(网址)

	if 障 != nil {
		return 障
	}

	// 检查回应状态码
	if 回.StatusCode != http.StatusOK {
		回.Body.Close()
		return fmt.Errorf("取%s障%s", 网址, 回.Status)
	}

	// 检查回应头部信息
	页型 := 回.Header.Get("Content-Type")
	if 页型 != "text/html" && !strings.HasPrefix(页型, "text/html;") {
		回.Body.Close()
		return fmt.Errorf("网址%s页型%s，非text/html", 网址, 页型)
	}

	网页, 障 := html.Parse(回.Body)
	回.Body.Close()

	if 障 != nil {
		return fmt.Errorf("析%s障%v", 网址, 障)
	}

	访节 := func(节 *html.Node) { // 匿函
		if 节.Type == html.ElementNode && 节.Data == "title" &&
			节.FirstChild != nil {
			fmt.Println(节.FirstChild.Data)
		}
	}

	逐节(网页, 访节, nil)

	return nil
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
