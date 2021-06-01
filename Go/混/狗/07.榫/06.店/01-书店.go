/* 一个简单的网上书店
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type 人民币 float64

func (价 人民币) String() string {
	return fmt.Sprintf("￥%.2f", 价)
}

type 图书库 map[string]人民币

func (库 图书库) ServeHTTP(回 http.ResponseWriter, 请 *http.Request) {
	switch 请.URL.Path {
	case "/书":
		for 书, 价 := range 库 {
			fmt.Fprintf(回, "书名：%s 单价：%s\n", 书, 价)
		}
	case "/价":
		书名 := 请.URL.Query().Get("书名")
		价, 有 := 库[书名]
		if !有 {
			回.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(回, "店里暂时没有《%s》\n", 书名)
			return
		}
		fmt.Fprintf(回, "%s售价为%s\n", 书名, 价)
	default:
		回.WriteHeader(http.StatusNotFound)
		问啥, _ := url.QueryUnescape(请.URL.String())
		fmt.Fprintf(回, "抱歉，您请求的网页“%s”不存在。\n", 问啥)
	}

}

func main() {
	库 := 图书库{"新白娘子传奇": 20, "时代浪潮": 30, "归去来兮": 25}
	log.Fatal(http.ListenAndServe("localhost:12345", 库))
}
