/* TCP(网控协议)时间服务器
* 向客户重复发送当前时间
 */

package main

import (
	"fmt"
	"io"
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
		//伺客(客) // 一次伺候一个客户
		go 伺客(客) // 并行伺候多个客户
	}
}

func 伺客(客 net.Conn) {
	defer 客.Close()
	for {
		时 := time.Now()

		if _, 障 := io.WriteString(客, fmt.Sprintf("城市：%s 时间：%s\n", 时.Location().String(), 时)); 障 != nil {
			return
		}

		if 市, 障 := time.LoadLocation("Asia/Shanghai"); 障 != nil {
			return
		} else {
			时 := time.Now().In(市)
			if _, 障 := io.WriteString(客, fmt.Sprintf("城市：%s 时间：%s\n", 时.Location(), 时)); 障 != nil {
				return
			}
		}

		time.Sleep(2 * time.Second)
	}
}
