/* 回应请求次数及请求信息
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var 锁 sync.Mutex
var 请次 int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法：", os.Args[0], "端号")
		os.Exit(1)
	}

	http.HandleFunc("/", 应声)
	http.HandleFunc("/请次", 计请)
	log.Fatal(http.ListenAndServe("localhost:"+os.Args[1], nil))
}

func 应声(回 http.ResponseWriter, 请 *http.Request) {
	锁.Lock()
	请次++
	锁.Unlock()

	fmt.Fprintf(回, "网址路径 = %q\n", 请.URL.Path)
	fmt.Fprintf(回, "请求方法：%v\n请求协议：%v\n", 请.Method, 请.Proto)

	fmt.Fprintf(回, "请求首部详情：\n")
	for 项, 值 := range 请.Header {
		fmt.Fprintf(回, "首[%q] = %q\n", 项, 值)
	}
	fmt.Fprintf(回, "伺者 = %q\n", 请.Host)
	fmt.Fprintf(回, "客址 = %q\n", 请.RemoteAddr)

	if 障 := 请.ParseForm(); 障 != nil {
		log.Printf("请求表障：%q\n", 障)
		return
	}
	for 项, 值 := range 请.Form {
		fmt.Fprintf(回, "表项[%q] = %q\n", 项, 值)
	}
}

func 计请(回 http.ResponseWriter, 请 *http.Request) {
	锁.Lock()
	fmt.Fprintf(回, "请求%d次\n", 请次)
	锁.Unlock()
}
