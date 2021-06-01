/*
 */

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	链, 障 := net.Dial("tcp", "localhost:12345")
	if 障 != nil {
		log.Fatal(障)
	}
	defer 链.Close()
	go io.Copy(os.Stdout, 链)
	io.Copy(链, os.Stdin)
}
