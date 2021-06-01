/* 应声虫伺者
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法：", os.Args[0], "端号")
		os.Exit(1)
	}

	http.HandleFunc("/", 应声)
	log.Fatal(http.ListenAndServe("localhost:"+os.Args[1], nil))
}

func 应声(回 http.ResponseWriter, 请 *http.Request) {
	fmt.Fprintf(回, "网址路径 = %q\n", 请.URL.Path)
}
