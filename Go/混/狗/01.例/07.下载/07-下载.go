/* 取显网址
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法：", os.Args[0], "网址")
	}

	for _, 网址 := range os.Args[1:] {
		if !strings.HasPrefix(网址, "http://") && !strings.HasPrefix(网址, "https://") {
			fmt.Println(网址, "缺少http:// 或 https://")
			continue
		}

		回, 障 := http.Get(网址)
		if 障 != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], 障)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "网址:%s\n应态:%v\n应码:%v\n", 网址, 回.Status, 回.StatusCode)

		// 法1：
		// 网页, 障 := ioutil.ReadAll(回.Body)
		// 回.Body.Close()
		// if 障 != nil {
		// 	fmt.Fprintf(os.Stderr, "%s: 取网址%s: %v\n", 网址, 障)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s\n", 网页)
		// 法2：
		_, 障 = io.Copy(os.Stdout, 回.Body)
		if 障 != nil {
			fmt.Fprintf(os.Stderr, "复制网页：%v\n", 障)
			os.Exit(1)
		}

	}
}
