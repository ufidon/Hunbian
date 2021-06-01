/* 接口
 */

package main

import "fmt"

type X字节计数器 int

func (计 *X字节计数器) Write(列 []byte) (int, error) {
	*计 += X字节计数器(len(列))
	return len(列), nil
}

func main() {
	var 计 X字节计数器
	计.Write([]byte("字节计数器"))
	fmt.Println(计)

	var 串 = "计数列占多少字节"
	fmt.Fprintf(&计, "%s", 串)
	fmt.Println(计)

	计 = 0
	串 = "清零重计"
	fmt.Fprintf(&计, "%s", 串)
	fmt.Println(计)
}
