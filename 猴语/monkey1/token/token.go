package token

import "fmt"

type C词类 string

const (
	T非法 = "非法"
	T档尾 = "档尾"

	// 标识及文值
	X标识 = "标识" // 名字, 身高, x, y, ...
	X整型 = "整型" // 1343456

	// 算符
	S赋 = "="
	S加 = "+"
	S减 = "-"
	S非 = "!"
	S乘 = "*"
	S除 = "/"

	S小于 = "<"
	S大于 = ">"

	S等于 = "=="
	S不等 = "!="

	// 界符
	J逗号 = ","
	J分号 = ";"

	J左括 = "("
	J右括 = ")"
	J左曲 = "{"
	J右曲 = "}"

	// 键词
	C函 = "函"
	C令 = "令"
	C真 = "真"
	C假 = "假"
	C若 = "若"
	C别 = "别"
	C返 = "返"
)

type C词 struct {
	C类 C词类
	C值 string
}

func (词 *C词) String() string {
	return fmt.Sprintf("%-10v:\t%-20v", 词.C类, 词.C值)
}

var 键词集 = map[string]C词类{
	"函": C函,
	"令": C令,
	"真": C真,
	"假": C假,
	"若": C若,
	"别": C别,
	"返": C返,
}

func C查词类(标 string) C词类 {
	if 词类, 有 := 键词集[标]; 有 {
		return 词类
	}
	return X标识
}
