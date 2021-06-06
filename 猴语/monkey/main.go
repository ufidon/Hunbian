/* 猴语交互解释器
造：
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" main.go
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" main.go
参考：
1、https://golang.org/cmd/go/
*/
package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	用户, 障 := user.Current()
	if 障 != nil {
		panic(障)
	}
	fmt.Printf("您好%s，谢谢您跟🐒交流!\n",
		用户.Username)
	fmt.Printf("入乡随俗，请用🐒语。\n")
	repl.H会话(os.Stdin, os.Stdout)
}
