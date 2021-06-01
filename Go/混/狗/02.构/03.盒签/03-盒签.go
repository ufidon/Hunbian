/* 行参处理
* 例：
* go run 02-盒签.go -h
* go run 02-盒签.go -尾=1 -隔 ：  李白 杜甫 白居易
* go run 02-盒签.go -尾=0 -隔 ：  李白 杜甫 白居易
 */
package main

import (
	"flag"
	"fmt"
	"strings"
)

var 尾 = flag.Bool("尾", false, "带行尾符")
var 隔 = flag.String("隔", " ", "隔字符")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *隔))
	if *尾 {
		fmt.Println("。")
	}
}
