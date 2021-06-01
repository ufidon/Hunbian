/* 延迟函数使用陷阱
* 跑：
* ./01-取链 https://golang.org/
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, 网址 := range os.Args[1:] {

		if 档名, 档夵, 障 := 取链(网址); 障 != nil {
			log.Println(障)
		} else {
			fmt.Printf("网址%s下主页保存为大小为%dB的%s\n", 网址, 档夵, 档名)
		}

	}
}

func 取链(网址 string) (档名 string, 档夵 int64, 障 error) {
	回, 障 := http.Get(网址)
	if 障 != nil {
		return "", 0, 障
	}

	defer 回.Body.Close()

	这 := path.Base(回.Request.URL.Path)
	if 这 == "/" {
		这 = "index.html"
	}

	档, 障 := os.Create(这)
	if 障 != nil {
		return "", 0, 障
	}

	档夵, 障 = io.Copy(档, 回.Body)
	// io.Copy，档.Close在不同文件系统中行为可能不同，io.Copy有可能在档.Close之后执行
	// 这是个使用延迟的陷阱
	if 闭障 := 档.Close(); 障 == nil {
		障 = 闭障
	}
	return 这, 档夵, 障
}
