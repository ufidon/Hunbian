/* 算式
* 支持64位标准浮点的：
1. 二元运算: +,-,*,/ （加减乘除）
2. 独元运算: +, - （正,负）
3. 函数：幂(底，指数), 正弦(角度)和平方根(数)
例：
（指数（利率率+1，年）+1）*本金
正弦（角）/角
平方根（面积/π）
*/

package ss

import (
	"fmt"
	"math"
	"strings"
)

type S变量 string  //标识变量，如半径
type 值 float64   // 标识数值，如0.707
type 独元 struct { // 标识独元算式，如 -温度
	算符 rune // +,-
	式  S式
}
type 二元 struct {
	算符   rune // +,-,*,/
	甲, 乙 S式
}
type 函数 struct { // 平方根(5)
	函 string // 幂，正弦，平方根
	参 []S式   // 自变量
}

type S境 map[S变量]float64 // 算式计算环境

type S式 interface {
	// 在境中计算本表达式并返回其值
	S算(境 S境) float64
	// 检报此式错误并收集其变量于映射表
	S检(变量集 map[S变量]bool) error
}

func (b S变量) S算(境 S境) float64 {
	return 境[b]
}
func (b S变量) S检(变量集 map[S变量]bool) error {
	变量集[b] = true
	return nil
}

func (z 值) S算(_ S境) float64 {
	return float64(z)
}
func (z 值) S检(变量集 map[S变量]bool) error {
	return nil
}

func (d 独元) S算(境 S境) float64 {
	switch d.算符 {
	case '+', '＋':
		return +d.式.S算(境)
	case '-', '－':
		return -d.式.S算(境)
	default:
		panic(fmt.Sprintf("不支持的独元算符：%q", d.算符))
	}
}
func (d 独元) S检(变量集 map[S变量]bool) error {
	if !strings.ContainsRune("+-＋－", d.算符) {
		return fmt.Errorf("未知独元算符%q", d.算符)
	}
	return d.式.S检(变量集)
}
func (e 二元) S算(境 S境) float64 {
	switch e.算符 {
	case '+', '＋':
		return e.甲.S算(境) + e.乙.S算(境)
	case '-', '－':
		return e.甲.S算(境) - e.乙.S算(境)
	case '*', '×', 'x', 'X':
		return e.甲.S算(境) * e.乙.S算(境)
	case '/', '÷':
		return e.甲.S算(境) / e.乙.S算(境)
	default:
		panic(fmt.Sprintf("不支持的二元算符：%q", e.算符))
	}
}
func (e 二元) S检(变量集 map[S变量]bool) error {
	if !strings.ContainsRune("+＋-－*×xX/÷", e.算符) {
		return fmt.Errorf("未知二元算符%q", e.算符)
	}
	if 障 := e.甲.S检(变量集); 障 != nil {
		return 障
	}
	return e.乙.S检(变量集)
}

func (h 函数) S算(境 S境) float64 {
	switch h.函 {
	case "pow", "幂":
		return math.Pow(h.参[0].S算(境), h.参[1].S算(境))
	case "sin", "正弦":
		return math.Sin(h.参[0].S算(境))
	case "sqrt", "平方根":
		return math.Sqrt(h.参[0].S算(境))
	default:
		panic(fmt.Sprintf("不支持的函数：%q", h.函))
	}
}

var 入数 = map[string]int{"pow": 2, "幂": 2, "sin": 1, "正弦": 1, "sqrt": 1, "平方根": 1}

func (h 函数) S检(变量集 map[S变量]bool) error {
	参目, 有 := 入数[h.函]
	if !有 {
		return fmt.Errorf("未知函名%q", h.函)
	}
	if len(h.参) != 参目 {
		return fmt.Errorf("%s被供%d参，但只要%d个", h.函, len(h.参), 参目)
	}
	for _, 参 := range h.参 {
		if 障 := 参.S检(变量集); 障 != nil {
			return 障
		}
	}
	return nil
}
