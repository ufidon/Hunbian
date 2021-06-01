/* 回应请求次数
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
}

func 计请(回 http.ResponseWriter, 请 *http.Request) {
	锁.Lock()
	fmt.Fprintf(回, "请求%d次\n", 请次)
	锁.Unlock()
}
