/* 断词判句 - 取
* 语义分析 - parse
 */
package ss

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// 取词器
type 取词器 struct {
	扫 scanner.Scanner
	令 rune // 当前令
}

func (取 *取词器) 翌()         { 取.令 = 取.扫.Scan() }
func (取 *取词器) 令文() string { return 取.扫.TokenText() }

type 词慌 string

// 叙述令文，用于错误信息.
func (取 *取词器) 叙令() string {
	switch 取.令 {
	case scanner.EOF:
		return "档尾"
	case scanner.Ident:
		return fmt.Sprintf("标识 %s", 取.令文())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("数 %s", 取.令文())
	}
	return fmt.Sprintf("%q", rune(取.令)) // 其余令
}

func 优先级(算符 rune) int {
	switch 算符 { //全角半角需可分辨字体
	case '*', '×', 'x', 'X', '/', '÷':
		return 2
	case '+', '＋', '-', '－':
		return 1
	}
	return 0
}

// 造句器

// S析 从输入文本中分析出句子，如算式.
//
//   式 = 数                         数值如, 3.14159
//        | 量名                          变量名如, 身高
//        | 量名 '(' 式 ',' ... ')'     函数调用
//        | '-' 式                    独元算子 (+-)
//        | 式 '+' 式               二元算子 (+-*/)
//
func S析(输入 string) (_ S式, 障 error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// 无慌
		case 词慌:
			障 = fmt.Errorf("%s", x)
		default:
			// 未料之慌: 恢慌态.
			panic(x)
		}
	}()
	取 := new(取词器)
	取.扫.Init(strings.NewReader(输入))
	取.扫.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	取.翌() // initial lookahead
	e := 析S式(取)
	if 取.令 != scanner.EOF {
		return nil, fmt.Errorf("未知 %s", 取.叙令())
	}
	return e, nil
}

func 析S式(取 *取词器) S式 { return 析二元式(取, 1) }

// 二元 = 独元 ('+' 二元)*
// 析二元式 停于算子其优先级低于优级.
func 析二元式(取 *取词器, 优级 int) S式 {
	左式 := 析独元式(取)
	for 级 := 优先级(取.令); 级 >= 优级; 级-- {
		for 优先级(取.令) == 级 {
			算符 := 取.令
			取.翌() // 消化算子
			右式 := 析二元式(取, 级+1)
			左式 = 二元{算符, 左式, 右式}
		}
	}
	return 左式
}

// 独元 = '+' 式 | 主式
func 析独元式(取 *取词器) S式 {
	if 取.令 == '+' || 取.令 == '-' {
		算符 := 取.令
		取.翌() // 消化 '+' 或 '-'
		return 独元{算符, 析独元式(取)}
	}
	return 析主式(取)
}

// 主式 = 量名
//         | 量名 '(' 式 ',' ... ',' 式 ')'
//         | 数
//         | '(' 式 ')'
func 析主式(取 *取词器) S式 {
	switch 取.令 {
	case scanner.Ident:
		量名 := 取.令文()
		取.翌() // 消化标识
		if 取.令 != '(' {
			return S变量(量名)
		}
		取.翌() // 消化 '('
		var 参集 []S式
		if 取.令 != ')' {
			for {
				参集 = append(参集, 析S式(取))
				if 取.令 != ',' {
					break
				}
				取.翌() // 消化 ','
			}
			if 取.令 != ')' {
				志 := fmt.Sprintf("得 %s, 要 ')'", 取.叙令())
				panic(词慌(志))
			}
		}
		取.翌() // 消化 ')'
		return 函数{量名, 参集}

	case scanner.Int, scanner.Float:
		f, 障 := strconv.ParseFloat(取.令文(), 64)
		if 障 != nil {
			panic(词慌(障.Error()))
		}
		取.翌() // 消化数
		return 值(f)

	case '(':
		取.翌() // 消化 '('
		e := 析S式(取)
		if 取.令 != ')' {
			志 := fmt.Sprintf("得 %s, 要 ')'", 取.叙令())
			panic(词慌(志))
		}
		取.翌() // 消化 ')'
		return e
	}
	志 := fmt.Sprintf("未知 %s", 取.叙令())
	panic(词慌(志))
}
