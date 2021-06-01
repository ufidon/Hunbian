/* 汇率
* 甲币对乙币汇率 = 多少甲币每乙币
 */

package huilv

import "fmt"

type H国 int

const (
	中 = iota
	美
	欧
	日
	韩
)

var H国名 = [...]string{
	中: "中国",
	美: "美国",
	欧: "欧盟",
	日: "日本",
	韩: "韩国",
}

var H币名 = [...]string{
	中: "人民币",
	美: "美刀",
	欧: "欧元",
	日: "日元",
	韩: "韩元",
}

type H汇对 struct {
	H汇出 H国
	H汇入 H国
}

func (汇 H汇对) String() string {
	return fmt.Sprintf("%s换%s", H币名[汇.H汇出], H币名[汇.H汇入])
}

func H汇率表(汇 H汇对) float64 {
	switch 汇 {
	case H汇对{中, 美}:
		return 1 / 6.42
	case H汇对{欧, 美}:
		return 1 / 0.82
	case H汇对{日, 美}:
		return 1 / 108.91
	case H汇对{韩, 美}:
		return 1 / 1128.38
	default:
		return 0.0
	}
}
