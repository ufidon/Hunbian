/* 取词器，lexer，读取码流，可以是字、数和符
* 据其模式取出词，按词类分门别类
 */
package lexer

import "monkey/token"

type Q取词器 struct {
	码流 []rune
	现位 int  // 当前文（字）之位，文可以是字、数、符
	翌位 int  // 位之次位
	文  rune // 当前文（字）
}

func Z造(流 string) *Q取词器 {
	取 := &Q取词器{码流: []rune(流)}
	取.读文()
	return 取
}

func (取 *Q取词器) Q取词() token.C词 {
	var 词 token.C词

	取.跨空()

	switch 取.文 {
	case '=':
		if 取.眺文() == '=' {
			文 := 取.文
			取.读文()
			文串 := string(文) + string(取.文)
			词 = token.C词{L类: token.D等于, W文: 文串}
		} else {
			词 = 造词(token.F赋, 取.文)
		}
	case '+':
		词 = 造词(token.J加, 取.文)
	case '-':
		词 = 造词(token.J减, 取.文)
	case '!':
		if 取.眺文() == '=' {
			文 := 取.文
			取.读文()
			文串 := string(文) + string(取.文)
			词 = token.C词{L类: token.B不等, W文: 文串}
		} else {
			词 = 造词(token.F非, 取.文)
		}
	case '/':
		词 = 造词(token.C除, 取.文)
	case '*':
		词 = 造词(token.C乘, 取.文)
	case '<':
		词 = 造词(token.X小于, 取.文)
	case '>':
		词 = 造词(token.D大于, 取.文)
	case ';':
		词 = 造词(token.F分号, 取.文)
	case ':':
		词 = 造词(token.M冒号, 取.文)
	case ',':
		词 = 造词(token.D逗号, 取.文)
	case '{':
		词 = 造词(token.Z左曲, 取.文)
	case '}':
		词 = 造词(token.Y右曲, 取.文)
	case '(':
		词 = 造词(token.Z左括, 取.文)
	case ')':
		词 = 造词(token.Y右括, 取.文)
	case '"':
		词.L类 = token.C串型
		词.W文 = 取.读串()
	case '[':
		词 = 造词(token.Z左方, 取.文)
	case ']':
		词 = 造词(token.Y右方, 取.文)
	case 0:
		词.W文 = ""
		词.L类 = token.D档尾
	default:
		if 是字(取.文) {
			词.W文 = 取.读标者()
			词.L类 = token.C查词类(词.W文)
			return 词
		} else if 是数(取.文) {
			词.L类 = token.Z整型
			词.W文 = 取.读数()
			return 词
		} else {
			词 = 造词(token.F非法, 取.文)
		}
	}

	取.读文()
	return 词
}

func (取 *Q取词器) 跨空() {
	for 取.文 == ' ' || 取.文 == '\t' || 取.文 == '\n' || 取.文 == '\r' {
		取.读文()
	}
}

func (取 *Q取词器) 读文() {
	if 取.翌位 >= len(取.码流) {
		取.文 = 0
	} else {
		取.文 = 取.码流[取.翌位]
	}
	取.现位 = 取.翌位
	取.翌位 += 1
}

func (取 *Q取词器) 眺文() rune {
	if 取.翌位 >= len(取.码流) {
		return 0
	} else {
		return 取.码流[取.翌位]
	}
}

func (取 *Q取词器) 读标者() string {
	现位 := 取.现位
	for 是字(取.文) {
		取.读文()
	}
	return string(取.码流[现位:取.现位])
}

func (取 *Q取词器) 读数() string {
	现位 := 取.现位
	for 是数(取.文) {
		取.读文()
	}
	return string(取.码流[现位:取.现位])
}

func (取 *Q取词器) 读串() string {
	现位 := 取.现位 + 1
	for {
		取.读文()
		if 取.文 == '"' || 取.文 == 0 {
			break
		}
	}
	return string(取.码流[现位:取.现位])
}

func 是字(文 rune) bool {
	return 'a' <= 文 && 文 <= 'z' || 'A' <= 文 && 文 <= 'Z' || 文 == '_' || 文 >= 0x4e00
}

func 是数(文 rune) bool {
	return '0' <= 文 && 文 <= '9'
}

func 造词(词类 token.C词类, 文 rune) token.C词 {
	return token.C词{L类: 词类, W文: string(文)}
}
