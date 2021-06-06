/* 计算语法树时需要各种大小的存储盒（区）来存相应的数据结构
 */
package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"monkey/ast"
	"strings"
)

type N内函 func(参集 ...H盒) H盒

type H盒型 int

const (
	K空盒 H盒型 = iota
	Z障盒
	Z整盒
	E贰盒
	C串盒
	F返盒
	H函盒
	N内盒
	D队盒
	O耦盒
)

var 盒名 = [...][2]string{
	K空盒: {"空盒", "NULL"},
	Z障盒: {"障盒", "ERROR"},
	Z整盒: {"整盒", "INTEGER"},
	E贰盒: {"贰盒", "BOOLEAN"},
	C串盒: {"串盒", "STRING"},
	F返盒: {"返盒", "RETURN_VALUE"},
	H函盒: {"函盒", "FUNCTION"},
	N内盒: {"内盒", "BUILTIN"},
	D队盒: {"队盒", "ARRAY"},
	O耦盒: {"耦盒", "HASH"},
}

func (盒 H盒型) String() string {
	return 盒名[盒][0]
}

type O耦键 struct {
	X型 H盒型
	Z值 uint64
}

type K可耦合 interface {
	O耦键() O耦键
}

type H盒 interface {
	X型() H盒型
	C察值() string
}

type Z整型 struct {
	Z值 int64
}

func (整 *Z整型) X型() H盒型     { return Z整盒 }
func (整 *Z整型) C察值() string { return fmt.Sprintf("%d", 整.Z值) }
func (整 *Z整型) O耦键() O耦键 {
	return O耦键{X型: 整.X型(), Z值: uint64(整.Z值)}
}

type E贰型 struct {
	Z值 bool
}

func (贰 *E贰型) X型() H盒型 { return E贰盒 }
func 贰值文(贰值 bool) string {
	if 贰值 {
		return "真"
	} else {
		return "假"
	}
}
func (贰 *E贰型) C察值() string { return 贰值文(贰.Z值) }
func (贰 *E贰型) O耦键() O耦键 {
	var 值 uint64

	if 贰.Z值 {
		值 = 1
	} else {
		值 = 0
	}

	return O耦键{X型: 贰.X型(), Z值: 值}
}

type K空型 struct{}

func (空 *K空型) X型() H盒型     { return K空盒 }
func (空 *K空型) C察值() string { return "空" }

type F返型 struct {
	Z值 H盒
}

func (返 *F返型) X型() H盒型     { return F返盒 }
func (返 *F返型) C察值() string { return 返.Z值.C察值() }

type Z障型 struct {
	Z志 string
}

func (障 *Z障型) X型() H盒型     { return Z障盒 }
func (障 *Z障型) C察值() string { return "障:" + 障.Z志 }

type H函型 struct {
	X形参集 []*ast.B标式
	T体   *ast.K块句
	J境   *S算境
}

func (函 *H函型) X型() H盒型 { return H函盒 }
func (函 *H函型) C察值() string {
	var out bytes.Buffer

	形参集 := []string{}
	for _, 形参 := range 函.X形参集 {
		形参集 = append(形参集, 形参.String())
	}

	out.WriteString("函")
	out.WriteString("(")
	out.WriteString(strings.Join(形参集, ", "))
	out.WriteString(") {\n")
	out.WriteString(函.T体.String())
	out.WriteString("\n}")

	return out.String()
}

type C串型 struct {
	Z值 string
}

func (串 *C串型) X型() H盒型     { return C串盒 }
func (串 *C串型) C察值() string { return 串.Z值 }
func (串 *C串型) O耦键() O耦键 {
	耦 := fnv.New64a()
	耦.Write([]byte(串.Z值))

	return O耦键{X型: 串.X型(), Z值: 耦.Sum64()}
}

type N内型 struct {
	H函 N内函
}

func (内 *N内型) X型() H盒型     { return N内盒 }
func (内 *N内型) C察值() string { return "内建函数" }

type D队型 struct {
	Y元集 []H盒
}

func (队 *D队型) X型() H盒型 { return D队盒 }
func (队 *D队型) C察值() string {
	var out bytes.Buffer

	元集 := []string{}
	for _, 元 := range 队.Y元集 {
		元集 = append(元集, 元.C察值())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(元集, ", "))
	out.WriteString("]")

	return out.String()
}

type O耦双型 struct {
	J键 H盒
	Z值 H盒
}

type O耦型 struct {
	S双 map[O耦键]O耦双型
}

func (耦 *O耦型) X型() H盒型 { return O耦盒 }
func (耦 *O耦型) C察值() string {
	var out bytes.Buffer

	双群 := []string{}
	for _, pair := range 耦.S双 {
		双群 = append(双群, fmt.Sprintf("%s: %s",
			pair.J键.C察值(), pair.Z值.C察值()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(双群, ", "))
	out.WriteString("}")

	return out.String()
}
