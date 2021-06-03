// 造：
// env GOOS=linux GOARCH=amd64 go build main.go
// env GOOS=windows GOARCH=amd64 go build main.gos
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
	fmt.Printf("您好%s，谢谢您跟猴子交流!\n",
		用户.Username)
	fmt.Printf("入乡随俗，请用猴语。\n")
	repl.H会话(os.Stdin, os.Stdout)
}
