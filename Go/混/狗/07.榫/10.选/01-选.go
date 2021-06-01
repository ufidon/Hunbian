/* 选印xml档元素
* 跑：./07-下载 http://www.w3.org/TR/2006/REC-xml11-20060816 | ./01-选 div div h2
 */

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	解 := xml.NewDecoder(os.Stdin)
	var 元叠 []string

	for {
		令, 障 := 解.Token()
		if 障 == io.EOF {
			break
		} else if 障 != nil {
			fmt.Fprintf(os.Stderr, "01-选：%v\n", 障)
			os.Exit(1)
		}

		switch 令 := 令.(type) {
		case xml.StartElement:
			元叠 = append(元叠, 令.Name.Local) //入叠
		case xml.EndElement:
			元叠 = 元叠[:len(元叠)-1] // 出叠
		case xml.CharData:
			if 全含(元叠, os.Args[1:]) {
				fmt.Printf("%s:%s\n", strings.Join(元叠, " "), 令)
			}
		}
	}
}

// 甲是否按序含乙之全部元素
func 全含(甲, 乙 []string) bool {
	for len(乙) <= len(甲) {
		if len(乙) == 0 {
			return true
		}
		if 甲[0] == 乙[0] {
			乙 = 乙[1:]
		}
		甲 = 甲[1:]
	}
	return false
}
