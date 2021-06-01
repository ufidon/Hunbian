/* 一个简单的网上书店
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

type 人民币 float64

func (价 人民币) String() string {
	return fmt.Sprintf("￥%.2f", 价)
}

type 图书库 map[string]人民币

func (库 图书库) ServeHTTP(回 http.ResponseWriter, 请 *http.Request) {
	for 书, 价 := range 库 {
		fmt.Fprintf(回, "书名：%s 单价：%s\n", 书, 价)
	}
}

func main() {
	库 := 图书库{"新白娘子传奇": 20, "时代浪潮": 30, "归去来兮": 25}
	log.Fatal(http.ListenAndServe("localhost:12345", 库))
}
