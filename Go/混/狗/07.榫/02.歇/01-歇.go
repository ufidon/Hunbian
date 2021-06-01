/* 自定义接口
 */
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var 一会儿 = flag.Duration("多久", 1*time.Second, "休息时间")
	flag.Parse()
	fmt.Printf("休息%v┄┄", 一会儿)
	time.Sleep(*一会儿)
	fmt.Println("休息完毕")
}
