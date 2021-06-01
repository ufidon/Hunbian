/* 汇率
* 甲币对乙币汇率 = 多少甲币每乙币
 */

package huilv

import "fmt"

func H换汇(汇 H汇对) float64 {
	汇款 := H汇率表(汇)
	if 汇款 <= 0.0 {
		fmt.Printf("%s 暂时不能换 %s\n", H币名[汇.H汇出], H币名[汇.H汇入])
		return 0.0
	}
	return 汇款
}
