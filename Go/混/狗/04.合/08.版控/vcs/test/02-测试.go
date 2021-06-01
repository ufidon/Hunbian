/* 测试Github之版控狗宝
 */
package main

import (
	"fmt"
	"log"
	"os"

	"vcs/github"
)

func main() {
	果, 障 := github.X寻障(os.Args[1:])
	if 障 != nil {
		log.Fatal(障)
	}
	fmt.Printf("%d 问题:\n", 果.Z障数)
	for _, 项 := range 果.Z障集 {
		fmt.Printf("#%-5d %9.9s %.55s\n", 项.S数, 项.Y用户.D登录名, 项.B标题)
	}
}
