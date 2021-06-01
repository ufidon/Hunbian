/* 改自找链1，处理狗慌
* 跑：
* ./01-网标 https://golang.org
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

	标题, 障 := 单标(网页)
	if 障 != nil {
		return 障
	}
	fmt.Printf("网页标题：%s\n", 标题)

	return nil
}

func 单标(网节 *html.Node) (标题 string, 障 error) {
	type 急救 struct{}

	defer func() { // 延迟处理狗慌
		switch 慌 := recover(); 慌 {
		case nil:
			// 没慌
		case 急救{}:
			// 预料到的慌，应作错误处理，勿慌
			障 = fmt.Errorf("有多个标题")
		default:
			panic(慌) // 慌不知所措，狗慌不止
		}
	}()

	逐节(网节, func(节 *html.Node) {
		if 节.Type == html.ElementNode && 节.Data == "title" &&
			节.FirstChild != nil {
			if 标题 != "" {
				panic(急救{})
			}
			标题 = 节.FirstChild.Data
		}
	}, nil)

	if 标题 == "" {
		return "", fmt.Errorf("无标题元素")
	}
	return 标题, nil
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
