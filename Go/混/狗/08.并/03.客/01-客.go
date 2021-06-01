/*网客，连接并读取服务器
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args[1:]) < 2 {
		fmt.Printf("用法：\n\t./客 网络协议[tcp或udp] 端口号\n\t例：./客 tcp 12345\n")
		os.Exit(1)
	}
	链, 障 := net.Dial(os.Args[1], "localhost:"+os.Args[2])
	if 障 != nil {
		log.Fatal(障)
	}
	defer 链.Close()
	读链(链, os.Stdout)
}

func 读链(源 io.Reader, 汇 io.Writer) {
	if _, 障 := io.Copy(汇, 源); 障 != nil {
		log.Fatal(障)
	}
}
