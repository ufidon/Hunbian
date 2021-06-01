/* 用管道协调狗程
 */

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	链, 障 := net.Dial("tcp", "127.0.0.1:12345")
	if 障 != nil {
		log.Fatal(障)
	}
	管 := make(chan struct{}, 0)

	go func() {
		io.Copy(os.Stdout, 链)
		log.Println("完毕")
		管 <- struct{}{} // 通知主狗程此子狗程已退出
	}()

	io.Copy(链, os.Stdin)
	链.Close()
	<-管
}
