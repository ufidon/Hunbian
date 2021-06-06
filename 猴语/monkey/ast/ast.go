/* 抽象语法树(ast)，由结(node)连接而成
* 结描述句(sentence)与式(expression)
* 榫(interface)
 */

package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

// 结榫
type J结 interface {
	C词文() string
	String() string
}

// 句榫
type J句 interface {
	J结
	句结()
}

// 式榫
type S式 interface {
	J结
	式结()
}

// 程序由句子组成
type C程 struct {
	J句集 []J句
}

// 程序词文为其首句词文
func (程 *C程) C词文() string {
	if len(程.J句集) > 0 {
		return 程.J句集[0].C词文()
	} else {
		return ""
	}
}

// 程序串由其语句串接而成
func (程 *C程) String() string {
	var 出 bytes.Buffer

	for _, 句 := range 程.J句集 {
		出.WriteString(句.String())
	}

	return 出.String()
}

// J句集含各种句型
// 令句
type L令句 struct {
	C词 token.C词 // "L令" 词
	M名 *B标式
	Z值 S式
}

func (令句 *L令句) 句结()         {}
func (令句 *L令句) C词文() string { return 令句.C词.W文 }
func (令句 *L令句) String() string {
	var 出 bytes.Buffer

	出.WriteString(令句.C词文())
	出.WriteString(" ")
	出.WriteString(令句.M名.String())
	出.WriteString(" = ")
	if 令句.Z值 != nil {
		出.WriteString(令句.Z值.String())
	}

	出.WriteString(";")

	return 出.String()
}

// 返句
type F返句 struct {
	C词  token.C词 // “F返”词
	F返值 S式
}

func (返句 *F返句) 句结()         {}
func (返句 *F返句) C词文() string { return 返句.C词.W文 }
func (返句 *F返句) String() string {
	var 出 bytes.Buffer

	出.WriteString(返句.C词文())

	if 返句.F返值 != nil {
		出.WriteString(返句.F返值.String())
	}

	出.WriteString(";")

	return 出.String()
}

type S式句 struct {
	C词 token.C词 // 式首词
	S式 S式
}

func (式句 *S式句) 句结()         {}
func (式句 *S式句) C词文() string { return 式句.C词.W文 }
func (式句 *S式句) String() string {
	if 式句.S式 != nil {
		return 式句.S式.String()
	}
	return ""
}

type K块句 struct {
	C词  token.C词 // "Z左曲" { 词
	J句集 []J句
}

func (块句 *K块句) 句结()         {}
func (块句 *K块句) C词文() string { return 块句.C词.W文 }
func (块句 *K块句) String() string {
	var 出 bytes.Buffer

	for _, 句 := range 块句.J句集 {
		出.WriteString(句.String())
	}

	return 出.String()
}

// 式，各种算式（表达式）
// 标式
type B标式 struct {
	C词 token.C词 // "B标识"词
	Z值 string
}

func (标 *B标式) 式结()            {}
func (标 *B标式) C词文() string    { return 标.C词.W文 }
func (标 *B标式) String() string { return 标.Z值 }

// 贰式
type E贰式 struct {
	C词 token.C词
	Z值 bool
}

func (贰式 *E贰式) 式结()            {}
func (贰式 *E贰式) C词文() string    { return 贰式.C词.W文 }
func (贰式 *E贰式) String() string { return 贰式.C词.W文 }

// 整式
type Z整式 struct {
	C词 token.C词
	Z值 int64
}

func (整式 *Z整式) 式结()            {}
func (整式 *Z整式) C词文() string    { return 整式.C词.W文 }
func (整式 *Z整式) String() string { return 整式.C词.W文 }

// 前式，算符前置式
type Q前式 struct {
	C词  token.C词 // 前置算符，如 +,-, !
	S算符 string
	Y右式 S式
}

func (前式 *Q前式) 式结()         {}
func (前式 *Q前式) C词文() string { return 前式.C词.W文 }
func (前式 *Q前式) String() string {
	var 出 bytes.Buffer

	出.WriteString("(")
	出.WriteString(前式.S算符)
	出.WriteString(前式.Y右式.String())
	出.WriteString(")")

	return 出.String()
}

// 中式，算符中置式
type Z中式 struct {
	C词  token.C词 // 二元算符，如 +,-,*,/
	Z左式 S式
	S算符 string
	Y右式 S式
}

func (中式 *Z中式) 式结()         {}
func (中式 *Z中式) C词文() string { return 中式.C词.W文 }
func (中式 *Z中式) String() string {
	var 出 bytes.Buffer

	出.WriteString("(")
	出.WriteString(中式.Z左式.String())
	出.WriteString(" " + 中式.S算符 + " ")
	出.WriteString(中式.Y右式.String())
	出.WriteString(")")

	return 出.String()
}

// 若式
type R若式 struct {
	C词  token.C词 // "R若"词
	T条件 S式
	S是块 *K块句
	F非块 *K块句
}

func (若式 *R若式) 式结()         {}
func (若式 *R若式) C词文() string { return 若式.C词.W文 }
func (若式 *R若式) String() string {
	var 出 bytes.Buffer

	出.WriteString("若")
	出.WriteString(若式.T条件.String())
	出.WriteString(" ")
	出.WriteString(若式.S是块.String())

	if 若式.F非块 != nil {
		出.WriteString("否 ")
		出.WriteString(若式.F非块.String())
	}

	return 出.String()
}

// 函式，函数当变量

type H函式 struct {
	C词   token.C词 // "H函"词
	X形参集 []*B标式
	T体   *K块句
}

func (函式 *H函式) 式结()         {}
func (函式 *H函式) C词文() string { return 函式.C词.W文 }
func (函式 *H函式) String() string {
	var 出 bytes.Buffer

	形参集 := []string{}
	for _, 形参 := range 函式.X形参集 {
		形参集 = append(形参集, 形参.String())
	}

	出.WriteString(函式.C词文())
	出.WriteString("(")
	出.WriteString(strings.Join(形参集, ", "))
	出.WriteString(") ")
	出.WriteString(函式.T体.String())

	return 出.String()
}

// 呼式，呼叫函式
type H呼式 struct {
	C词   token.C词 // “(”词
	H函名  S式       // 函式名
	S实参集 []S式
}

func (呼式 *H呼式) 式结()         {}
func (呼式 *H呼式) C词文() string { return 呼式.C词.W文 }
func (呼式 *H呼式) String() string {
	var 出 bytes.Buffer

	实参集 := []string{}
	for _, 实参 := range 呼式.S实参集 {
		实参集 = append(实参集, 实参.String())
	}

	出.WriteString(呼式.H函名.String())
	出.WriteString("(")
	出.WriteString(strings.Join(实参集, ", "))
	出.WriteString(")")

	return 出.String()
}

// 串式，串值
type C串式 struct {
	C词 token.C词
	Z值 string
}

func (串式 *C串式) 式结()            {}
func (串式 *C串式) C词文() string    { return 串式.C词.W文 }
func (串式 *C串式) String() string { return 串式.C词.W文 }

// 队式
type D队式 struct {
	C词  token.C词 // "["词
	Y元集 []S式
}

func (队式 *D队式) 式结()         {}
func (队式 *D队式) C词文() string { return 队式.C词.W文 }
func (队式 *D队式) String() string {
	var 出 bytes.Buffer

	元集 := []string{}
	for _, 元 := range 队式.Y元集 {
		元集 = append(元集, 元.String())
	}

	出.WriteString("[")
	出.WriteString(strings.Join(元集, ", "))
	出.WriteString("]")

	return 出.String()
}

// 号式，以号取项
type H号式 struct {
	C词  token.C词 //  [ 词
	Z左式 S式
	H号  S式
}

func (号式 *H号式) 式结()         {}
func (号式 *H号式) C词文() string { return 号式.C词.W文 }
func (号式 *H号式) String() string {
	var 出 bytes.Buffer

	出.WriteString("(")
	出.WriteString(号式.Z左式.String())
	出.WriteString("[")
	出.WriteString(号式.H号.String())
	出.WriteString("])")

	return 出.String()
}

// 耦式
type O耦式 struct {
	C词  token.C词 //  '{' 词
	O耦集 map[S式]S式
}

func (耦式 *O耦式) 式结()         {}
func (耦式 *O耦式) C词文() string { return 耦式.C词.W文 }
func (耦式 *O耦式) String() string {
	var 出 bytes.Buffer

	耦集 := []string{}
	for 键, 值 := range 耦式.O耦集 {
		耦集 = append(耦集, 键.String()+":"+值.String())
	}

	出.WriteString("{")
	出.WriteString(strings.Join(耦集, ", "))
	出.WriteString("}")

	return 出.String()
}
