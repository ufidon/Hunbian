/* 析词成句
* 取词器(lexer)组字成词，
* 树句器(parser)读取分析词语(token)及其关系，组成句子，
* 判断句子是否合乎语法的, 并以语法树表示合法句式
 */
package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

// 算序，计算顺序,算符优先级,重在相对级别，不在其值
const (
	_   int = iota
	Z最后     // 最后计算，最低优先级
	X相等     // ==
	D大小     // > or <
	Q求和     // +
	Q求积     // *
	Q前缀     // -X , !X
	S使函     // 使用函数 求和(X)
	H号项     // 项[号]，以号取项
)

var 算序集 = map[token.C词类]int{
	token.D等于: X相等,
	token.B不等: X相等,
	token.X小于: D大小,
	token.D大于: D大小,
	token.J加:  Q求和,
	token.J减:  Q求和,
	token.C乘:  Q求积,
	token.C除:  Q求积,
	token.Z左括: S使函,
	token.Z左方: H号项,
}

type (
	前析函 func() ast.S式
	中析函 func(ast.S式) ast.S式
)

type S树句器 struct {
	取  *lexer.Q取词器
	障集 []string

	现词 token.C词
	翌词 token.C词

	前析函集 map[token.C词类]前析函
	中析函集 map[token.C词类]中析函
}

func (树 *S树句器) 绑前析函(词类 token.C词类, fn 前析函) {
	树.前析函集[词类] = fn
}

func (树 *S树句器) 绑中析函(词类 token.C词类, fn 中析函) {
	树.中析函集[词类] = fn
}

func Z造(取 *lexer.Q取词器) *S树句器 {
	树 := &S树句器{
		取:  取,
		障集: []string{},
	}

	树.前析函集 = make(map[token.C词类]前析函)
	树.绑前析函(token.B标识, 树.析标式)
	树.绑前析函(token.Z整型, 树.析整式)
	树.绑前析函(token.C串型, 树.析串式)
	树.绑前析函(token.F非, 树.析前式)
	树.绑前析函(token.J减, 树.析前式)
	树.绑前析函(token.Z真, 树.析贰式)
	树.绑前析函(token.J假, 树.析贰式)
	树.绑前析函(token.Z左括, 树.析群式)
	树.绑前析函(token.R若, 树.析若式)
	树.绑前析函(token.H函, 树.析函式)
	树.绑前析函(token.Z左方, 树.析队式)
	树.绑前析函(token.Z左曲, 树.析耦式)

	树.中析函集 = make(map[token.C词类]中析函)
	树.绑中析函(token.J加, 树.析中式)
	树.绑中析函(token.J减, 树.析中式)
	树.绑中析函(token.C除, 树.析中式)
	树.绑中析函(token.C乘, 树.析中式)
	树.绑中析函(token.D等于, 树.析中式)
	树.绑中析函(token.B不等, 树.析中式)
	树.绑中析函(token.X小于, 树.析中式)
	树.绑中析函(token.D大于, 树.析中式)

	树.绑中析函(token.Z左括, 树.析呼式)
	树.绑中析函(token.Z左方, 树.析号式)

	// 连读两词设置现词与翌词
	树.取词()
	树.取词()

	return 树
}

func (树 *S树句器) 取词() {
	树.现词 = 树.翌词
	树.翌词 = 树.取.Q取词()
}

func (树 *S树句器) 现词类是(词类 token.C词类) bool {
	return 树.现词.L类 == 词类
}

func (树 *S树句器) 翌词类是(词类 token.C词类) bool {
	return 树.翌词.L类 == 词类
}

func (树 *S树句器) 期翌词类(词类 token.C词类) bool {
	if 树.翌词类是(词类) {
		树.取词()
		return true
	} else {
		树.翌词类障(词类)
		return false
	}
}

func (树 *S树句器) Z障集() []string {
	return 树.障集
}

func (树 *S树句器) 翌词类障(词类 token.C词类) {
	志 := fmt.Sprintf("翌词期望为%s,却得%s。", 词类, 树.翌词.L类)
	树.障集 = append(树.障集, 志)
}

func (树 *S树句器) 无前析函障(词类 token.C词类) {
	志 := fmt.Sprintf("%s未绑定前析函", 词类)
	树.障集 = append(树.障集, 志)
}

func (树 *S树句器) X析程() *ast.C程 {
	程 := &ast.C程{}
	程.J句集 = []ast.J句{}

	for !树.现词类是(token.D档尾) {
		句 := 树.析句()
		if 句 != nil {
			程.J句集 = append(程.J句集, 句)
		}
		树.取词()
	}

	return 程
}

func (树 *S树句器) 析句() ast.J句 {
	switch 树.现词.L类 {
	case token.L令:
		return 树.析令句()
	case token.F返:
		return 树.析返句()
	default:
		return 树.析式句()
	}
}

func (树 *S树句器) 析令句() *ast.L令句 {
	句 := &ast.L令句{C词: 树.现词}

	if !树.期翌词类(token.B标识) {
		return nil
	}

	句.M名 = &ast.B标式{C词: 树.现词, Z值: 树.现词.W文}

	if !树.期翌词类(token.F赋) {
		return nil
	}

	树.取词()

	句.Z值 = 树.析式(Z最后)

	if 树.翌词类是(token.F分号) {
		树.取词()
	}

	return 句
}

func (树 *S树句器) 析返句() *ast.F返句 {
	句 := &ast.F返句{C词: 树.现词}

	树.取词()

	句.F返值 = 树.析式(Z最后)

	if 树.翌词类是(token.F分号) {
		树.取词()
	}

	return 句
}

func (树 *S树句器) 析式句() *ast.S式句 {
	句 := &ast.S式句{C词: 树.现词}

	句.S式 = 树.析式(Z最后)

	if 树.翌词类是(token.F分号) {
		树.取词()
	}

	return 句
}

func (树 *S树句器) 析式(算序 int) ast.S式 {
	前析函 := 树.前析函集[树.现词.L类]
	if 前析函 == nil {
		树.无前析函障(树.现词.L类)
		return nil
	}
	左式 := 前析函()

	for !树.翌词类是(token.F分号) && 算序 < 树.翌词算序() {
		中析函 := 树.中析函集[树.翌词.L类]
		if 中析函 == nil {
			return 左式
		}

		树.取词()

		左式 = 中析函(左式)
	}

	return 左式
}

func (树 *S树句器) 翌词算序() int {
	if 树, 有 := 算序集[树.翌词.L类]; 有 {
		return 树
	}

	return Z最后
}

func (树 *S树句器) 现词算序() int {
	if 树, 有 := 算序集[树.现词.L类]; 有 {
		return 树
	}

	return Z最后
}

func (树 *S树句器) 析标式() ast.S式 {
	return &ast.B标式{C词: 树.现词, Z值: 树.现词.W文}
}

func (树 *S树句器) 析整式() ast.S式 {
	整式 := &ast.Z整式{C词: 树.现词}

	值, 障 := strconv.ParseInt(树.现词.W文, 0, 64)
	if 障 != nil {
		志 := fmt.Sprintf("%q不能作整式", 树.现词.W文)
		树.障集 = append(树.障集, 志)
		return nil
	}

	整式.Z值 = 值

	return 整式
}

func (树 *S树句器) 析串式() ast.S式 {
	return &ast.C串式{C词: 树.现词, Z值: 树.现词.W文}
}

func (树 *S树句器) 析前式() ast.S式 {
	式 := &ast.Q前式{
		C词:  树.现词,
		S算符: 树.现词.W文,
	}

	树.取词()

	式.Y右式 = 树.析式(Q前缀)

	return 式
}

func (树 *S树句器) 析中式(左式 ast.S式) ast.S式 {
	式 := &ast.Z中式{
		C词:  树.现词,
		S算符: 树.现词.W文,
		Z左式: 左式,
	}

	算序 := 树.现词算序()
	树.取词()
	式.Y右式 = 树.析式(算序)

	return 式
}

func (树 *S树句器) 析贰式() ast.S式 {
	return &ast.E贰式{C词: 树.现词, Z值: 树.现词类是(token.Z真)}
}

func (树 *S树句器) 析群式() ast.S式 {
	树.取词()

	式 := 树.析式(Z最后)

	if !树.期翌词类(token.Y右括) {
		return nil
	}

	return 式
}

func (树 *S树句器) 析若式() ast.S式 {
	式 := &ast.R若式{C词: 树.现词}

	if !树.期翌词类(token.Z左括) {
		return nil
	}

	树.取词()
	式.T条件 = 树.析式(Z最后)

	if !树.期翌词类(token.Y右括) {
		return nil
	}

	if !树.期翌词类(token.Z左曲) {
		return nil
	}

	式.S是块 = 树.析块句()

	if 树.翌词类是(token.F否) {
		树.取词()

		if !树.期翌词类(token.Z左曲) {
			return nil
		}

		式.F非块 = 树.析块句()
	}

	return 式
}

func (树 *S树句器) 析块句() *ast.K块句 {
	块 := &ast.K块句{C词: 树.现词}
	块.J句集 = []ast.J句{}

	树.取词()

	for !树.现词类是(token.Y右曲) && !树.现词类是(token.D档尾) {
		句 := 树.析句()
		if 句 != nil {
			块.J句集 = append(块.J句集, 句)
		}
		树.取词()
	}

	return 块
}

func (树 *S树句器) 析函式() ast.S式 {
	函式 := &ast.H函式{C词: 树.现词}

	if !树.期翌词类(token.Z左括) {
		return nil
	}

	函式.X形参集 = 树.析形参集()

	if !树.期翌词类(token.Z左曲) {
		return nil
	}

	函式.T体 = 树.析块句()

	return 函式
}

func (树 *S树句器) 析形参集() []*ast.B标式 {
	标式集 := []*ast.B标式{}

	if 树.翌词类是(token.Y右括) {
		树.取词()
		return 标式集
	}

	树.取词()

	标式 := &ast.B标式{C词: 树.现词, Z值: 树.现词.W文}
	标式集 = append(标式集, 标式)

	for 树.翌词类是(token.D逗号) {
		树.取词()
		树.取词()
		标式 := &ast.B标式{C词: 树.现词, Z值: 树.现词.W文}
		标式集 = append(标式集, 标式)
	}

	if !树.期翌词类(token.Y右括) {
		return nil
	}

	return 标式集
}

func (树 *S树句器) 析呼式(函名 ast.S式) ast.S式 {
	式 := &ast.H呼式{C词: 树.现词, H函名: 函名}
	式.S实参集 = 树.析式列(token.Y右括)
	return 式
}

func (树 *S树句器) 析式列(尾词 token.C词类) []ast.S式 {
	式列 := []ast.S式{}

	if 树.翌词类是(尾词) {
		树.取词()
		return 式列
	}

	树.取词()
	式列 = append(式列, 树.析式(Z最后))

	for 树.翌词类是(token.D逗号) {
		树.取词()
		树.取词()
		式列 = append(式列, 树.析式(Z最后))
	}

	if !树.期翌词类(尾词) {
		return nil
	}

	return 式列
}

func (树 *S树句器) 析队式() ast.S式 {
	队 := &ast.D队式{C词: 树.现词}

	队.Y元集 = 树.析式列(token.Y右方)

	return 队
}

func (树 *S树句器) 析号式(左式 ast.S式) ast.S式 {
	式 := &ast.H号式{C词: 树.现词, Z左式: 左式}

	树.取词()
	式.H号 = 树.析式(Z最后)

	if !树.期翌词类(token.Y右方) {
		return nil
	}

	return 式
}

func (树 *S树句器) 析耦式() ast.S式 {
	耦式 := &ast.O耦式{C词: 树.现词}
	耦式.O耦集 = make(map[ast.S式]ast.S式)

	for !树.翌词类是(token.Y右曲) {
		树.取词()
		键 := 树.析式(Z最后)

		if !树.期翌词类(token.M冒号) {
			return nil
		}

		树.取词()
		值 := 树.析式(Z最后)

		耦式.O耦集[键] = 值

		if !树.翌词类是(token.Y右曲) && !树.期翌词类(token.D逗号) {
			return nil
		}
	}

	if !树.期翌词类(token.Y右曲) {
		return nil
	}

	return 耦式
}
