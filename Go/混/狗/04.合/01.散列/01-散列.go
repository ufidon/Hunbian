/* 算消息散列值
 */
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	m1 := sha256.Sum256([]byte("明"))
	m2 := sha256.Sum256([]byte("文"))

	fmt.Printf("'明'之sha256散列值：%x\n‘文’之sha256散列值：%x\n两散列值相等否？%t\n散列值类型：%T\n", m1, m2, m1 == m2, m1)
}
