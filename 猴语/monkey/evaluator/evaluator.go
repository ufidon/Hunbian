/* 现写现算
* 即时计算输入的猴语程序
 */
package evaluator

import (
	"fmt"
	"monkey/ast"
	"monkey/object"
)

var (
	K空值 = &object.K空型{}
	K真值 = &object.E贰型{Z值: true}
	K假值 = &object.E贰型{Z值: false}
)

func G估(结 ast.J结, 算境 *object.S算境) object.H盒 {
	switch 结 := 结.(type) {

	// 各种句
	case *ast.C程:
		return 估程(结, 算境)

	case *ast.K块句:
		return 估块句(结, 算境)

	case *ast.S式句:
		return G估(结.S式, 算境)

	case *ast.F返句:
		值 := G估(结.F返值, 算境)
		if 是障(值) {
			return 值
		}
		return &object.F返型{Z值: 值}

	case *ast.L令句:
		值 := G估(结.Z值, 算境)
		if 是障(值) {
			return 值
		}
		算境.S设(结.M名.Z值, 值)

	// 各种式
	case *ast.Z整式:
		return &object.Z整型{Z值: 结.Z值}

	case *ast.C串式:
		return &object.C串型{Z值: 结.Z值}

	case *ast.E贰式:
		return 狗贰型转猴贰型(结.Z值)

	case *ast.Q前式:
		右盒 := G估(结.Y右式, 算境)
		if 是障(右盒) {
			return 右盒
		}
		return 估前式(结.S算符, 右盒)

	case *ast.Z中式:
		左盒 := G估(结.Z左式, 算境)
		if 是障(左盒) {
			return 左盒
		}

		右盒 := G估(结.Y右式, 算境)
		if 是障(右盒) {
			return 右盒
		}

		return 估中式(结.S算符, 左盒, 右盒)

	case *ast.R若式:
		return 估若式(结, 算境)

	case *ast.B标式:
		return 估标式(结, 算境)

	case *ast.H函式:
		形参集 := 结.X形参集
		函体 := 结.T体
		return &object.H函型{X形参集: 形参集, J境: 算境, T体: 函体}

	case *ast.H呼式:
		函型 := G估(结.H函名, 算境)
		if 是障(函型) {
			return 函型
		}

		实参集 := 估式集(结.S实参集, 算境)
		if len(实参集) == 1 && 是障(实参集[0]) {
			return 实参集[0]
		}

		return 用函(函型, 实参集)

	case *ast.D队式:
		元集 := 估式集(结.Y元集, 算境)
		if len(元集) == 1 && 是障(元集[0]) {
			return 元集[0]
		}
		return &object.D队型{Y元集: 元集}

	case *ast.H号式:
		左盒 := G估(结.Z左式, 算境)
		if 是障(左盒) {
			return 左盒
		}
		号 := G估(结.H号, 算境)
		if 是障(号) {
			return 号
		}
		return 估队式(左盒, 号)

	case *ast.O耦式:
		return 估耦式(结, 算境)

	}

	return nil
}

func 估程(程 *ast.C程, 算境 *object.S算境) object.H盒 {
	var 果 object.H盒

	for _, 句 := range 程.J句集 {
		果 = G估(句, 算境)

		switch 果 := 果.(type) {
		case *object.F返型:
			return 果.Z值
		case *object.Z障型:
			return 果
		}
	}

	return 果
}

func 估块句(
	句群 *ast.K块句,
	算境 *object.S算境,
) object.H盒 {
	var 果 object.H盒

	for _, 句 := range 句群.J句集 {
		果 = G估(句, 算境)

		if 果 != nil {
			盒 := 果.X型()
			if 盒 == object.F返盒 || 盒 == object.Z障盒 {
				return 果
			}
		}
	}

	return 果
}

func 狗贰型转猴贰型(贰 bool) *object.E贰型 {
	if 贰 {
		return K真值
	}
	return K假值
}

func 估前式(算符 string, 右盒 object.H盒) object.H盒 {
	switch 算符 {
	case "!":
		return 估非前式(右盒)
	case "-":
		return 估负前式(右盒)
	default:
		return 造障("未知算符: %s %s", 算符, 右盒.X型())
	}
}

func 估中式(
	算符 string,
	左盒, 右盒 object.H盒,
) object.H盒 {
	switch {
	case 左盒.X型() == object.Z整盒 && 右盒.X型() == object.Z整盒:
		return 估整中式(算符, 左盒, 右盒)
	case 左盒.X型() == object.C串盒 && 右盒.X型() == object.C串盒:
		return 估串中式(算符, 左盒, 右盒)
	case 算符 == "==":
		return 狗贰型转猴贰型(左盒 == 右盒)
	case 算符 == "!=":
		return 狗贰型转猴贰型(左盒 != 右盒)
	case 左盒.X型() != 右盒.X型():
		return 造障("型不配: %s %s %s",
			左盒.X型(), 算符, 右盒.X型())
	default:
		return 造障("未知算符: %s %s %s",
			左盒.X型(), 算符, 右盒.X型())
	}
}

func 估非前式(右盒 object.H盒) object.H盒 {
	switch 右盒 {
	case K真值:
		return K假值
	case K假值:
		return K真值
	case K空值:
		return K真值
	default:
		return K假值
	}
}

func 估负前式(右盒 object.H盒) object.H盒 {
	if 右盒.X型() != object.Z整盒 {
		return 造障("未知算符: -%s", 右盒.X型())
	}

	值 := 右盒.(*object.Z整型).Z值
	return &object.Z整型{Z值: -值}
}

func 估整中式(
	算符 string,
	左盒, 右盒 object.H盒,
) object.H盒 {
	左值 := 左盒.(*object.Z整型).Z值
	右值 := 右盒.(*object.Z整型).Z值

	switch 算符 {
	case "+":
		return &object.Z整型{Z值: 左值 + 右值}
	case "-":
		return &object.Z整型{Z值: 左值 - 右值}
	case "*":
		return &object.Z整型{Z值: 左值 * 右值}
	case "/":
		return &object.Z整型{Z值: 左值 / 右值}
	case "<":
		return 狗贰型转猴贰型(左值 < 右值)
	case ">":
		return 狗贰型转猴贰型(左值 > 右值)
	case "==":
		return 狗贰型转猴贰型(左值 == 右值)
	case "!=":
		return 狗贰型转猴贰型(左值 != 右值)
	default:
		return 造障("未知算符: %s %s %s",
			左盒.X型(), 算符, 右盒.X型())
	}
}

func 估串中式(
	算符 string,
	左盒, 右盒 object.H盒,
) object.H盒 {
	if 算符 != "+" {
		return 造障("未知算符: %s %s %s",
			左盒.X型(), 算符, 右盒.X型())
	}

	左值 := 左盒.(*object.C串型).Z值
	右值 := 右盒.(*object.C串型).Z值
	return &object.C串型{Z值: 左值 + 右值}
}

func 估若式(
	若式 *ast.R若式,
	算境 *object.S算境,
) object.H盒 {
	条件 := G估(若式.T条件, 算境)
	if 是障(条件) {
		return 条件
	}

	if 是真(条件) {
		return G估(若式.S是块, 算境)
	} else if 若式.F非块 != nil {
		return G估(若式.F非块, 算境)
	} else {
		return K空值
	}
}

func 估标式(
	结 *ast.B标式,
	算境 *object.S算境,
) object.H盒 {
	if 值, 得 := 算境.Q取(结.Z值); 得 {
		return 值
	}

	if builtin, 得 := 内函集[结.Z值]; 得 {
		return builtin
	}

	return 造障("未知标识: " + 结.Z值)
}

func 是真(盒 object.H盒) bool {
	switch 盒 {
	case K空值:
		return false
	case K真值:
		return true
	case K假值:
		return false
	default:
		return true
	}
}

func 造障(format string, a ...interface{}) *object.Z障型 {
	return &object.Z障型{Z志: fmt.Sprintf(format, a...)}
}

func 是障(盒 object.H盒) bool {
	if 盒 != nil {
		return 盒.X型() == object.Z障盒
	}
	return false
}

func 估式集(
	式集 []ast.S式,
	算境 *object.S算境,
) []object.H盒 {
	var 果 []object.H盒

	for _, 式 := range 式集 {
		估完 := G估(式, 算境)
		if 是障(估完) {
			return []object.H盒{估完}
		}
		果 = append(果, 估完)
	}

	return 果
}

func 用函(函 object.H盒, 实参集 []object.H盒) object.H盒 {
	switch 函 := 函.(type) {

	case *object.H函型:
		扩境 := 扩函境(函, 实参集)
		估完 := G估(函.T体, 扩境)
		return 解返值(估完)

	case *object.N内型:
		return 函.H函(实参集...)

	default:
		return 造障("非函型: %s", 函.X型())
	}
}

func 扩函境(
	函 *object.H函型,
	参集 []object.H盒,
) *object.S算境 {
	算境 := object.Z造闭境(函.J境)

	for 号, 参 := range 函.X形参集 {
		算境.S设(参.Z值, 参集[号])
	}

	return 算境
}

func 解返值(盒 object.H盒) object.H盒 {
	if 返值, 是 := 盒.(*object.F返型); 是 {
		return 返值.Z值
	}

	return 盒
}

func 估队式(左盒, 号 object.H盒) object.H盒 {
	switch {
	case 左盒.X型() == object.D队盒 && 号.X型() == object.Z整盒:
		return 估队号式(左盒, 号)
	case 左盒.X型() == object.O耦盒:
		return 估耦号式(左盒, 号)
	default:
		return 造障("%s号型不支持", 左盒.X型())
	}
}

func 估队号式(队盒, 号盒 object.H盒) object.H盒 {
	队 := 队盒.(*object.D队型)
	号 := 号盒.(*object.Z整型).Z值
	尾号 := int64(len(队.Y元集) - 1)

	if 号 < 0 || 号 > 尾号 {
		return K空值
	}

	return 队.Y元集[号]
}

func 估耦式(
	结 *ast.O耦式,
	算境 *object.S算境,
) object.H盒 {
	双 := make(map[object.O耦键]object.O耦双型)

	for 键结, 值结 := range 结.O耦集 {
		键 := G估(键结, 算境)
		if 是障(键) {
			return 键
		}

		耦键, 是 := 键.(object.K可耦合)
		if !是 {
			return 造障("%s不可用作耦键", 键.X型())
		}

		值 := G估(值结, 算境)
		if 是障(值) {
			return 值
		}

		得耦 := 耦键.O耦键()
		双[得耦] = object.O耦双型{J键: 键, Z值: 值}
	}

	return &object.O耦型{S双: 双}
}

func 估耦号式(耦, 号 object.H盒) object.H盒 {
	耦盒 := 耦.(*object.O耦型)

	键, 是 := 号.(object.K可耦合)
	if !是 {
		return 造障("%s不可用作耦键", 号.X型())
	}

	双, 有 := 耦盒.S双[键.O耦键()]
	if !有 {
		return K空值
	}

	return 双.Z值
}
