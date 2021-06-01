/* 网聊
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	听者, 障 := net.Listen("tcp", "localhost:12345")
	if 障 != nil {
		log.Fatal(障)
	}
	go 广播者()
	for {
		链, 障 := 听者.Accept()
		if 障 != nil {
			log.Print(障)
			continue
		}
		go 搞链(链)
	}
}

type 客 chan<- string

var (
	来客 = make(chan 客)
	离客 = make(chan 客)
	消息 = make(chan string) // 所有客户发入的消息，聊天室
)

func 广播者() {
	客们 := make(map[客]bool) // 所有在线客户
	for {
		select {
		case 息 := <-消息:
			for 宀 := range 客们 { // 消息广播给所有在线客户
				宀 <- 息
			}
		case 宀 := <-来客: // 客户加入
			客们[宀] = true
		case 宀 := <-离客: // 客户离开
			delete(客们, 宀)
			close(宀)
		}
	}
}

func 搞链(链 net.Conn) {
	播信 := make(chan string)
	go 客写者(链, 播信)

	谁 := 链.RemoteAddr().String()
	播信 <- "您是" + 谁
	消息 <- 谁 + "来了"
	来客 <- 播信

	入 := bufio.NewScanner(链)
	for 入.Scan() {
		消息 <- 谁 + ":" + 入.Text()
	}

	离客 <- 播信
	消息 <- 谁 + "走了"
	链.Close()
}

func 客写者(链 net.Conn, 播信 <-chan string) {
	for 信 := range 播信 {
		fmt.Fprintln(链, 信)
	}
}
