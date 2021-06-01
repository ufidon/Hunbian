/* 并行取显网址
* 报告耗时及节数
 */
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法：", os.Args[0], "网址")
		os.Exit(1)
	}

	始 := time.Now()
	管 := make(chan string)

	for _, 网址 := range os.Args[1:] {
		go 下载(网址, 管)
	}

	fmt.Print("耗时\t 节数\t 网址\n")
	for range os.Args[1:] {
		fmt.Println(<-管)
	}

	fmt.Printf("总耗时：%.2f秒\n", time.Since(始).Seconds())
}

func 下载(网址 string, 管 chan<- string) {
	始 := time.Now()
	回, 障 := http.Get(网址)

	if 障 != nil {
		管 <- fmt.Sprint(障)
		return
	}

	节数, 障 := io.Copy(ioutil.Discard, 回.Body)
	回.Body.Close()
	if 障 != nil {
		管 <- fmt.Sprintf("读网址%s: %v\n", 网址, 障)
		return
	}

	耗时 := time.Since(始).Seconds()
	管 <- fmt.Sprintf("%.f秒\t%7d节\t%s", 耗时, 节数, 网址)
}
