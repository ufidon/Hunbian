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

func (库 图书库) 询书(回 http.ResponseWriter, 请 *http.Request) {
	for 书, 价 := range 库 {
		fmt.Fprintf(回, "书名：%s 单价：%s\n", 书, 价)
	}
}

func (库 图书库) 询价(回 http.ResponseWriter, 请 *http.Request) {
	书名 := 请.URL.Query().Get("书名")
	价, 有 := 库[书名]
	if !有 {
		回.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(回, "店里暂时没有《%s》\n", 书名)
		return
	}
	fmt.Fprintf(回, "《%s》售价为%s\n", 书名, 价)
}

func main() {
	库 := 图书库{"新白娘子传奇": 20, "时代浪潮": 30, "归去来兮": 25}

	http.HandleFunc("/书", http.HandlerFunc(库.询书))
	http.HandleFunc("/价", http.HandlerFunc(库.询价))

	log.Fatal(http.ListenAndServe("localhost:12345", nil))
}
