/* 鹦鹉学舌的服务器
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	伺者, 障 := net.Listen("tcp", "localhost:12345")
	if 障 != nil {
		log.Fatal(障)
	}
	for {
		客, 障 := 伺者.Accept()
		if 障 != nil {
			log.Println(障)
			continue
		}
		go 仿客(客) // 并行模仿多个客户
	}
}

func 仿客(客 net.Conn) {
	客语 := bufio.NewScanner(客)
	for 客语.Scan() {
		// 模仿(客, 客语.Text(), 2*time.Second)
		go 模仿(客, 客语.Text(), 2*time.Second)
	}
	客.Close()
}

func 模仿(客 net.Conn, 戏语 string, 延 time.Duration) {
	fmt.Fprintln(客, 戏语, "😄")
	time.Sleep(延)
	fmt.Fprintln(客, 戏语, "😂")
	time.Sleep(延)
	fmt.Fprintln(客, 戏语, "😋")
}
