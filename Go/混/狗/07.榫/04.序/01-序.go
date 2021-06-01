/* 排序接口
 */

package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// type X菜 struct {
// 	菜名 string
// 	配料 []string
// 	单价 float64
// 	耗时 time.Duration
// }
// var 川菜 = []*X菜{
// 	{"回锅肉", []string{"五花肉", "青红椒", "青蒜", "豆豉", "郫县豆瓣酱"}, 15.66, 20},
// 	{"麻豆腐", []string{"豆腐", "牛肉末", "豆瓣酱", "豆豉"}, 10.16, 10},
// 	{"水煮鱼", []string{"草鱼（黑鱼）", "盐", "生粉", "食用油", "大蒜", "香菜", "葱", "生姜", "麻椒", "花椒"}, 30.22, 30},
// 	{"红烧肉", []string{"品五花肉", "炖肉料包", "葱", "冰糖", "茶叶"}, 25.50, 28},
// }

type X菜 struct {
	菜名 string
	特价 bool
	单价 float64
	耗时 time.Duration
}

var 川菜 = []*X菜{
	{"回锅肉", true, 15.66, 8 * time.Minute},
	{"麻豆腐", true, 15.66, 10 * time.Minute},
	{"水煮鱼", false, 30.22, 30 * time.Minute},
	{"红烧肉", false, 25.50, 28 * time.Minute},
}

func 时长(串 string) time.Duration {
	时, 障 := time.ParseDuration(串)
	if 障 != nil {
		panic(串)
	}
	return 时
}

func 印菜单(点菜 []*X菜) {
	const 格式 = "%v\t%v\t%v\t%v\n"
	表 := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(表, 格式, "菜名", "特价", "单价", "耗时")
	fmt.Fprintf(表, 格式, "------", "------", "------", "------")
	for _, 菜 := range 点菜 {
		fmt.Fprintf(表, 格式, 菜.菜名, 菜.特价, 菜.单价, 菜.耗时)
	}
	表.Flush()
}

// 按菜名排序
type 菜名排序 []*X菜

func (菜 菜名排序) Len() int           { return len(菜) }
func (菜 菜名排序) Less(甲, 乙 int) bool { return 菜[甲].菜名 < 菜[乙].菜名 }
func (菜 菜名排序) Swap(甲, 乙 int)      { 菜[甲], 菜[乙] = 菜[乙], 菜[甲] }

// 按菜名排序
type 菜价排序 []*X菜

func (菜 菜价排序) Len() int           { return len(菜) }
func (菜 菜价排序) Less(甲, 乙 int) bool { return 菜[甲].单价 < 菜[乙].单价 }
func (菜 菜价排序) Swap(甲, 乙 int)      { 菜[甲], 菜[乙] = 菜[乙], 菜[甲] }

// 多重排序
type 多重排序 struct {
	菜  []*X菜
	小于 func(甲, 乙 *X菜) bool
}

func (菜 多重排序) Len() int           { return len(菜.菜) }
func (菜 多重排序) Less(甲, 乙 int) bool { return 菜.小于(菜.菜[甲], 菜.菜[乙]) }
func (菜 多重排序) Swap(甲, 乙 int)      { 菜.菜[甲], 菜.菜[乙] = 菜.菜[乙], 菜.菜[甲] }

func main() {
	fmt.Printf("\n1.按菜名顺序:\n")
	sort.Sort(菜名排序(川菜))
	印菜单(川菜)
	fmt.Printf("\n2.按菜名逆序:\n")
	sort.Sort(sort.Reverse(菜名排序(川菜)))
	印菜单(川菜)

	fmt.Printf("\n3.按菜价顺序:\n")
	sort.Sort(菜价排序(川菜))
	印菜单(川菜)

	fmt.Printf("\n4.多重排序:\n")
	sort.Sort(多重排序{川菜, func(甲, 乙 *X菜) bool {
		if 甲.单价 != 乙.单价 {
			return 甲.单价 < 乙.单价
		}
		if 甲.耗时 != 乙.耗时 {
			return 甲.耗时 < 乙.耗时
		}
		return false
	}})
	印菜单(川菜)
}
