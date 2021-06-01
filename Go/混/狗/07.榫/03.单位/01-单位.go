/* 度量单位设计与命令行输入
 */
package main

import (
	"bytes"
	"flag"
	"fmt"
)

type D长度 float64

func (长度 D长度) String() string {
	return fmt.Sprintf("%g米", 长度)
}

type 长度行参 struct {
	D长度
}

func (长度 *长度行参) Set(行参 string) error {
	var 单位 string
	var 值 float64
	fmt.Sscanf(行参, "%f%s", &值, &单位)
	switch 单位 {
	case "米", "m":
		长度.D长度 = D长度(值)
		return nil
	case "分米", "dm":
		长度.D长度 = D长度(值 / 10)
		return nil
	case "厘米", "cm":
		长度.D长度 = D长度(值 / 100)
		return nil
	}
	return fmt.Errorf("无效长度%q", 行参)
}

func D长度行参(名 string, 值 D长度, 用法 string) *D长度 {
	长度 := 长度行参{值}
	flag.CommandLine.Var(&长度, 名, 用法)
	return &长度.D长度
}

func main() {
	var 身高 = D长度行参("身高", 1.75, "平均男性身高")

	flag.Parse()
	fmt.Println("您输入的身高：", 身高)

	var 空 interface{}
	fmt.Printf("空接口之型：%T与值%#[1]v\n", 空)
	空 = new(bytes.Buffer)
	fmt.Printf("空接口之型：%T与值%#[1]v\n", 空)

}
