package lexer

import "monkey/token"

type Q取词器 struct {
	码流 []rune
	位  int  // 当前文（字）之位，文可以是字、数、符
	翌位 int  // 位之次位
	文  rune // 当前文（字）
}

func Q造(流 string) *Q取词器 {
	取 := &Q取词器{码流: []rune(流)}
	取.读字()
	return 取
}

func (取 *Q取词器) Q翌词() token.C词 {
	var 词 token.C词

	取.跨空()

	switch 取.文 {
	case '=':
		if 取.眺字() == '=' {
			文 := 取.文
			取.读字()
			文值 := string(文) + string(取.文)
			词 = token.C词{C类: token.S加, C值: 文值}
		} else {
			词 = 新词(token.S赋, 取.文)
		}
	case '+':
		词 = 新词(token.S加, 取.文)
	case '-':
		词 = 新词(token.S减, 取.文)
	case '!':
		if 取.眺字() == '=' {
			文 := 取.文
			取.读字()
			文值 := string(文) + string(取.文)
			词 = token.C词{C类: token.S不等, C值: 文值}
		} else {
			词 = 新词(token.S非, 取.文)
		}
	case '/':
		词 = 新词(token.S除, 取.文)
	case '*':
		词 = 新词(token.S乘, 取.文)
	case '<':
		词 = 新词(token.S小于, 取.文)
	case '>':
		词 = 新词(token.S大于, 取.文)
	case ';':
		词 = 新词(token.J分号, 取.文)
	case ',':
		词 = 新词(token.J逗号, 取.文)
	case '{':
		词 = 新词(token.J左曲, 取.文)
	case '}':
		词 = 新词(token.J右曲, 取.文)
	case '(':
		词 = 新词(token.J左括, 取.文)
	case ')':
		词 = 新词(token.J右括, 取.文)
	case 0:
		词.C值 = ""
		词.C类 = token.T档尾
	default:
		if 是字(取.文) {
			词.C值 = 取.读标识()
			词.C类 = token.C查词类(词.C值)
			return 词
		} else if 是数(取.文) {
			词.C类 = token.X整型
			词.C值 = 取.读数()
			return 词
		} else {
			词 = 新词(token.T非法, 取.文)
		}
	}

	取.读字()
	return 词
}

func (取 *Q取词器) 跨空() {
	for 取.文 == ' ' || 取.文 == '\t' || 取.文 == '\n' || 取.文 == '\r' {
		取.读字()
	}
}

func (取 *Q取词器) 读字() {
	if 取.翌位 >= len(取.码流) {
		取.文 = 0
	} else {
		取.文 = 取.码流[取.翌位]
	}
	取.位 = 取.翌位
	取.翌位 += 1
}

func (取 *Q取词器) 眺字() rune {
	if 取.翌位 >= len(取.码流) {
		return 0
	} else {
		return 取.码流[取.翌位]
	}
}

func (取 *Q取词器) 读标识() string {
	位 := 取.位
	for 是字(取.文) {
		取.读字()
	}
	return string(取.码流[位:取.位])
}

func (取 *Q取词器) 读数() string {
	位 := 取.位
	for 是数(取.文) {
		取.读字()
	}
	return string(取.码流[位:取.位])
}

func 是字(文 rune) bool {
	return 'a' <= 文 && 文 <= 'z' || 'A' <= 文 && 文 <= 'Z' || 文 == '_' || 文 >= 0x4e00
}

func 是数(文 rune) bool {
	return '0' <= 文 && 文 <= '9'
}

func 新词(词类 token.C词类, 文 rune) token.C词 {
	return token.C词{C类: 词类, C值: string(文)}
}
