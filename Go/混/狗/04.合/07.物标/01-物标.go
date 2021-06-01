/* JSON(Java脚本物标) <-> 狗数据结构
* Java脚本又名爪脚语
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type D电影 struct {
	M片名  string   `json:"片名"`
	R发布日 string   `json:"发布日期"`
	J级别  int      `json:"分级"`
	J禁   bool     `json:"禁止未成年观看,omitempty"`
	Y演员表 []string `json:"演员表"`
}

func main() {
	var 影集 = []D电影{
		{
			M片名:  "拆弹专家2",
			R发布日: "2020-12-24",
			J级别:  3, J禁: false,
			Y演员表: []string{"刘德华", "刘青云", "倪妮", "谢君豪", "姜皓文", "吴卓羲", "马浴柯"},
		},
		{
			M片名:  "色，戒",
			R发布日: "2007-11-1",
			J级别:  5, J禁: true,
			Y演员表: []string{"梁朝伟", "汤唯", "陈冲", "王力宏", "庹宗华", "钱嘉乐", "朱芷莹"},
		},
		{
			M片名:  "地道战",
			R发布日: "1965-10-1",
			J级别:  1, J禁: false,
			Y演员表: []string{"朱龙广", "王炳彧", "张勇手"},
		},
	}

	// 狗数据结构 -> 爪脚语物标
	fmt.Print("\n狗数据结构 -> 爪脚语物标:\n")
	if 物标, 障 := json.Marshal(影集); 障 != nil {
		log.Fatalf("转爪脚语物标失败: %s", 障)
	} else {
		fmt.Printf("\n无格式爪脚语物标:\n")
		fmt.Printf("%s\n", 物标)
	}

	物标, 障 := json.MarshalIndent(影集, "", "  ")
	if 障 != nil {
		log.Fatalf("转爪脚语物标失败: %s", 障)
	} else {
		fmt.Print("\n缩进式爪脚语物标:\n")
		fmt.Printf("%s\n", 物标)
	}

	// 爪脚语物标 -> 狗数据结构， 仅取片名
	fmt.Print("\n爪脚语物标 -> 狗数据结构， 仅取片名:\n")

	var 影集名 []struct {
		M片名 string `json:"片名"` // 爪签不能省略，否则不能获取片名
	}
	if 障 := json.Unmarshal(物标, &影集名); 障 != nil {
		log.Fatalf("爪脚语物标转狗语结构败: %s", 障)
	}
	fmt.Println(影集名)
}
