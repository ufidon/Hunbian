/* 恒量生成函数iota
 */
package main

import (
	"fmt"
	"math/big"
)

func main() {
	const (
		B = 1 << (iota * 10)
		KB
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fmt.Printf("B : %T %#[1]v\n", B)
	fmt.Printf("KB: %T %#[1]v\n", KB)
	fmt.Printf("MB: %T %#[1]v\n", MB)
	fmt.Printf("GB: %T %#[1]v\n", GB)
	fmt.Printf("TB: %T %#[1]v\n", TB)
	fmt.Printf("PB: %T %#[1]v\n", PB)
	fmt.Printf("EB: %T %#[1]v\n", EB)
	/* 溢出 int64
	fmt.Printf("ZB: %T %#[1]v\n", ZB)
	fmt.Printf("YB: %T %#[1]v\n", YB)
	*/
	if zb, 成 := new(big.Int).SetString("1180591620717411303424", 10); 成 {
		fmt.Println("ZB: ", zb)
	}
	if yb, 成 := new(big.Int).SetString("1208925819614629174706176", 10); 成 {
		fmt.Println("YB: ", yb)
	}
}
