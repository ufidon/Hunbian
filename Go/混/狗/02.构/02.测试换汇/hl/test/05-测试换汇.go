/* 测试huilv换算包
 */
package main

import (
	"fmt"
	"hl/huilv"
)

func main() {
	用法 := `
	换汇小程序:汇出款转汇入款
	支持币种: 0-人民币，1-美元，2-日元，3-欧元，4-韩元
	输入数字选择汇入汇出币种，然后输入汇出款。
	例：把1668.8块换成美元
	0 1 1668.8
	`
	var 汇入, 汇出 huilv.H国
	var 汇出款 float64

	fmt.Println(用法)
	fmt.Scanf("%d%d%f", &汇出, &汇入, &汇出款)
	汇率 := huilv.H汇率表(huilv.H汇对{汇出, 汇入})
	if 汇率 <= 0.0 {
		fmt.Printf("汇率表中没找到%s对%s的汇率，请更新汇率表。\n", huilv.H币名[汇出], huilv.H币名[汇入])
	} else {
		汇入款 := 汇率 * 汇出款
		fmt.Printf("%f%s可以换%f%s\n", 汇出款, huilv.H币名[汇出], 汇入款, huilv.H币名[汇入])
	}

}
