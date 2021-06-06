package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func Test令句群(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期标 string
		期值 interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = y;", "foobar", "y"},
		{"令 甲 = 5;", "甲", 5},
		{"令 乙 = 真;", "乙", true},
		{"令 谁啊 = 乙;", "谁啊", "乙"},
	}

	for _, 测试 := range 测试集 {
		取 := lexer.Z造(测试.待测)
		树 := Z造(取)
		程 := 树.X析程()
		检树障(测, 树)

		if len(程.J句集) != 1 {
			测.Fatalf("程.J句集应有1句，却得%d",
				len(程.J句集))
		}

		句 := 程.J句集[0]
		if !测令句(测, 句, 测试.期标) {
			return
		}

		值 := 句.(*ast.L令句).Z值
		if !测值式(测, 值, 测试.期值) {
			return
		}
	}
}

func Test返句群(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期值 interface{}
	}{
		{"return 5;", 5},
		{"return true;", true},
		{"return foobar;", "foobar"},
		{"返 5;", 5},
		{"返 真;", true},
		{"返 假;", false},
		{"返 销售额;", "销售额"},
	}

	for _, 测试 := range 测试集 {
		取 := lexer.Z造(测试.待测)
		树 := Z造(取)
		程 := 树.X析程()
		检树障(测, 树)

		if len(程.J句集) != 1 {
			测.Fatalf("程.J句集应有1句，却得%d", len(程.J句集))
		}

		句 := 程.J句集[0]
		返句, 是 := 句.(*ast.F返句)
		if !是 {
			测.Fatalf("stmt应为*ast.F返句，却得%T", 句)
		}
		if 返句.C词文() != "返" && 返句.C词文() != "return" {
			测.Fatalf("返句.C词文应为'返'或'return',却得%q",
				返句.C词文())
		}
		if 测值式(测, 返句.F返值, 测试.期值) {
			return
		}
	}
}

func Test式句群(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期值 interface{}
	}{
		{"销售额;", "销售额"},
		{"sales;", "sales"},
	}

	for _, 测试 := range 测试集 {

		取 := lexer.Z造(测试.待测)
		树 := Z造(取)
		程 := 树.X析程()
		检树障(测, 树)

		if len(程.J句集) != 1 {
			测.Fatalf("程应含1句，却得%d",
				len(程.J句集))
		}
		句, 是 := 程.J句集[0].(*ast.S式句)
		if !是 {
			测.Fatalf("程.J句集[0]应为ast.S式句，却得%T",
				程.J句集[0])
		}

		标式, 是 := 句.S式.(*ast.B标式)
		if !是 {
			测.Fatalf("exp应为*ast.B标式，却得%T", 句.S式)
		}
		if 标式.Z值 != 测试.期值 {
			测.Errorf("标式.Z值应为%s，却得%s", 测试.期值, 标式.Z值)
		}
		if 标式.C词文() != 测试.期值 {
			测.Errorf("标式.C词文应为%s，却得%s", 测试.期值,
				标式.C词文())
		}
	}
}

func Test整式群(测 *testing.T) {
	待测 := "5;"

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	if len(程.J句集) != 1 {
		测.Fatalf("程应含1句，却得%d",
			len(程.J句集))
	}
	句, 是 := 程.J句集[0].(*ast.S式句)
	if !是 {
		测.Fatalf("程.J句集[0]应为ast.S式句，却得%T",
			程.J句集[0])
	}

	整式, 是 := 句.S式.(*ast.Z整式)
	if !是 {
		测.Fatalf("exp应为*ast.Z整式，却得%T", 句.S式)
	}
	if 整式.Z值 != 5 {
		测.Errorf("整式.Z值应为%d，却得%d", 5, 整式.Z值)
	}
	if 整式.C词文() != "5" {
		测.Errorf("整式.C词文应为%s，却得%s", "5",
			整式.C词文())
	}
}

func Test前式群(测 *testing.T) {
	前式测集 := []struct {
		待测 string
		算符 string
		值  interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!foobar;", "!", "foobar"},
		{"-foobar;", "-", "foobar"},
		{"!true;", "!", true},
		{"!false;", "!", false},
		{"!产量;", "!", "产量"},
		{"-产量;", "-", "产量"},
		{"!真;", "!", true},
		{"!假;", "!", false},
	}

	for _, 测试 := range 前式测集 {
		取 := lexer.Z造(测试.待测)
		树 := Z造(取)
		程 := 树.X析程()
		检树障(测, 树)

		if len(程.J句集) != 1 {
			测.Fatalf("程.J句集应含1句，却得%d\n", len(程.J句集))
		}

		句, 是 := 程.J句集[0].(*ast.S式句)
		if !是 {
			测.Fatalf("程.J句集[0]应为ast.S式句，却得%T",
				程.J句集[0])
		}

		前式, 是 := 句.S式.(*ast.Q前式)
		if !是 {
			测.Fatalf("应为ast.Q前式，却得%T", 句.S式)
		}
		if 前式.S算符 != 测试.算符 {
			测.Fatalf("前式.S算符应为'%s'，却得%s",
				测试.算符, 前式.S算符)
		}
		if !测值式(测, 前式.Y右式, 测试.值) {
			return
		}
	}
}

func Test中式群(测 *testing.T) {
	中式测集 := []struct {
		待测 string
		左值 interface{}
		算符 string
		右值 interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"foobar + barfoo;", "foobar", "+", "barfoo"},
		{"foobar - barfoo;", "foobar", "-", "barfoo"},
		{"foobar * barfoo;", "foobar", "*", "barfoo"},
		{"foobar / barfoo;", "foobar", "/", "barfoo"},
		{"foobar > barfoo;", "foobar", ">", "barfoo"},
		{"foobar < barfoo;", "foobar", "<", "barfoo"},
		{"foobar == barfoo;", "foobar", "==", "barfoo"},
		{"foobar != barfoo;", "foobar", "!=", "barfoo"},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
		{"产量 + 量产;", "产量", "+", "量产"},
		{"产量 - 量产;", "产量", "-", "量产"},
		{"产量 * 量产;", "产量", "*", "量产"},
		{"产量 / 量产;", "产量", "/", "量产"},
		{"产量 > 量产;", "产量", ">", "量产"},
		{"产量 < 量产;", "产量", "<", "量产"},
		{"产量 == 量产;", "产量", "==", "量产"},
		{"产量 != 量产;", "产量", "!=", "量产"},
		{"真 == 真", true, "==", true},
		{"真 != 假", true, "!=", false},
		{"假 == 假", false, "==", false},
	}

	for _, 测试 := range 中式测集 {
		取 := lexer.Z造(测试.待测)
		树 := Z造(取)
		程 := 树.X析程()
		检树障(测, 树)

		if len(程.J句集) != 1 {
			测.Fatalf("程.J句集应含1句，却得%d\n", len(程.J句集))
		}

		句, 是 := 程.J句集[0].(*ast.S式句)
		if !是 {
			测.Fatalf("程.J句集[0]应为ast.S式句，却得%T",
				程.J句集[0])
		}

		if !测中式(测, 句.S式, 测试.左值,
			测试.算符, 测试.右值) {
			return
		}
	}
}

func Test算序群(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"三 + 四 * 五 == 三 * 一 + 四 * 五",
			"((三 + (四 * 五)) == ((三 * 一) + (四 * 五)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"真",
			"真",
		},
		{
			"假",
			"假",
		},
		{
			"3 > 5 == 假",
			"((3 > 5) == 假)",
		},
		{
			"3 < 5 == 真",
			"((3 < 5) == 真)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"(5 + 5) * 2 * (5 + 5)",
			"(((5 + 5) * 2) * (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
		{
			"!(真 == 真)",
			"(!(真 == 真))",
		},
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"求和(甲, 乙, 1, 2 * 3, 4 + 5, 求和(6, 7 * 8))",
			"求和(甲, 乙, 1, (2 * 3), (4 + 5), 求和(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
		{
			"a * [1, 2, 3, 4][b * c] * d",
			"((a * ([1, 2, 3, 4][(b * c)])) * d)",
		},
		{
			"add(a * b[2], b[1], 2 * [1, 2][1])",
			"add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))",
		},
		{
			"某函(产量 * 成本[2], 成本[1], 2 * [1, 2][1])",
			"某函((产量 * (成本[2])), (成本[1]), (2 * ([1, 2][1])))",
		},
	}

	for _, 测试 := range 测试集 {
		取 := lexer.Z造(测试.待测)
		树 := Z造(取)
		程 := 树.X析程()
		检树障(测, 树)

		实得 := 程.String()
		if 实得 != 测试.期望 {
			测.Errorf("期望%q, 却得%q", 测试.期望, 实得)
		}
	}
}

func Test贰式群(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 bool
	}{
		{"true;", true},
		{"false;", false},
		{"真;", true},
		{"假;", false},
	}

	for _, 测试 := range 测试集 {
		取 := lexer.Z造(测试.待测)
		树 := Z造(取)
		程 := 树.X析程()
		检树障(测, 树)

		if len(程.J句集) != 1 {
			测.Fatalf("程应含1句，却得%d",
				len(程.J句集))
		}

		句, 是 := 程.J句集[0].(*ast.S式句)
		if !是 {
			测.Fatalf("程.J句集[0]应为ast.S式句，却得%T",
				程.J句集[0])
		}

		贰式, 是 := 句.S式.(*ast.E贰式)
		if !是 {
			测.Fatalf("应为*ast.E贰式，却得%T", 句.S式)
		}
		if 贰式.Z值 != 测试.期望 {
			测.Errorf("贰式.Z值应为%t，却得%t", 测试.期望,
				贰式.Z值)
		}
	}
}

func Test若式群(测 *testing.T) {
	待测 := `若 (x < y) { x }`
	//待测 := `if (x < y) { x }`

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	if len(程.J句集) != 1 {
		测.Fatalf("程.T体应含1句，却得%d\n", len(程.J句集))
	}

	句, 是 := 程.J句集[0].(*ast.S式句)
	if !是 {
		测.Fatalf("程.J句集[0]应为ast.S式句，却得%T",
			程.J句集[0])
	}

	若式, 是 := 句.S式.(*ast.R若式)
	if !是 {
		测.Fatalf("句.S式应为ast.R若式，却得%T",
			句.S式)
	}

	if !测中式(测, 若式.T条件, "x", "<", "y") {
		return
	}

	if len(若式.S是块.J句集) != 1 {
		测.Errorf("S是块应含1句，却得%d\n",
			len(若式.S是块.J句集))
	}

	是块, 是 := 若式.S是块.J句集[0].(*ast.S式句)
	if !是 {
		测.Fatalf("J句集[0]应为ast.S式句，却得%T",
			若式.S是块.J句集[0])
	}

	if !测标式(测, 是块.S式, "x") {
		return
	}

	if 若式.F非块 != nil {
		测.Errorf("若式.F非块.J句集应为nil，却得%+v", 若式.F非块)
	}
}

func Test若否式群(测 *testing.T) {
	待测 := `若 (x < y) { x } 否 { y }`
	//待测 := `if (x < y) { x } else { y }`

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	if len(程.J句集) != 1 {
		测.Fatalf("程.T体应含1句，却得%d\n", len(程.J句集))
	}

	句, 是 := 程.J句集[0].(*ast.S式句)
	if !是 {
		测.Fatalf("程.J句集[0]应为ast.S式句，却得%T",
			程.J句集[0])
	}

	若否式, 是 := 句.S式.(*ast.R若式)
	if !是 {
		测.Fatalf("句.S式应为ast.若式，却得%T", 句.S式)
	}

	if !测中式(测, 若否式.T条件, "x", "<", "y") {
		return
	}

	if len(若否式.S是块.J句集) != 1 {
		测.Errorf("是块应含1句，却得%d\n",
			len(若否式.S是块.J句集))
	}

	是块, 是 := 若否式.S是块.J句集[0].(*ast.S式句)
	if !是 {
		测.Fatalf("J句集[0]应为，ast.S式句却得%T",
			若否式.S是块.J句集[0])
	}

	if !测标式(测, 是块.S式, "x") {
		return
	}

	if len(若否式.F非块.J句集) != 1 {
		测.Errorf("若否式.F非块.J句集应含1句，却得%d\n",
			len(若否式.F非块.J句集))
	}

	非块, 是 := 若否式.F非块.J句集[0].(*ast.S式句)
	if !是 {
		测.Fatalf("J句集[0]应为ast.S式句，却得%T",
			若否式.F非块.J句集[0])
	}

	if !测标式(测, 非块.S式, "y") {
		return
	}
}

func Test函式群(测 *testing.T) {
	待测 := `函(x, y) { x + y; }`
	//待测 := `fn(x, y) { x + y; }`

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	if len(程.J句集) != 1 {
		测.Fatalf("程.T体应含1句，却得%d\n", len(程.J句集))
	}

	句, 是 := 程.J句集[0].(*ast.S式句)
	if !是 {
		测.Fatalf("程.J句集[0]应为ast.S式句，却得%T",
			程.J句集[0])
	}

	函式, 是 := 句.S式.(*ast.H函式)
	if !是 {
		测.Fatalf("句.S式应为ast.H函式，却得%T",
			句.S式)
	}

	if len(函式.X形参集) != 2 {
		测.Fatalf("函式形参数应为2, 却得%d\n",
			len(函式.X形参集))
	}

	测值式(测, 函式.X形参集[0], "x")
	测值式(测, 函式.X形参集[1], "y")

	if len(函式.T体.J句集) != 1 {
		测.Fatalf("函式.T体.J句集应含1句却得%d\n",
			len(函式.T体.J句集))
	}

	体句, 是 := 函式.T体.J句集[0].(*ast.S式句)
	if !是 {
		测.Fatalf("函体首句应为ast.S式句，却得%T",
			函式.T体.J句集[0])
	}

	测中式(测, 体句.S式, "x", "+", "y")
}

func Test函形参群(测 *testing.T) {
	测试集 := []struct {
		待测   string
		期形参集 []string
	}{
		{待测: "fn() {};", 期形参集: []string{}},
		{待测: "fn(x) {};", 期形参集: []string{"x"}},
		{待测: "fn(x, y, z) {};", 期形参集: []string{"x", "y", "z"}},
		{待测: "函() {};", 期形参集: []string{}},
		{待测: "函(甲) {};", 期形参集: []string{"甲"}},
		{待测: "函(甲,乙,丙) {};", 期形参集: []string{"甲", "乙", "丙"}},
	}

	for _, 测试 := range 测试集 {
		取 := lexer.Z造(测试.待测)
		树 := Z造(取)
		程 := 树.X析程()
		检树障(测, 树)

		句 := 程.J句集[0].(*ast.S式句)
		函式 := 句.S式.(*ast.H函式)

		if len(函式.X形参集) != len(测试.期形参集) {
			测.Errorf("形参数应为%d, 却得%d\n",
				len(测试.期形参集), len(函式.X形参集))
		}

		for 号, 标 := range 测试.期形参集 {
			测值式(测, 函式.X形参集[号], 标)
		}
	}
}

func Test呼式群(测 *testing.T) {
	待测 := "某函(1, 2 * 3, 4 + 5);"

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	if len(程.J句集) != 1 {
		测.Fatalf("程.J句集应含1句,却得%d\n", len(程.J句集))
	}

	句, 是 := 程.J句集[0].(*ast.S式句)
	if !是 {
		测.Fatalf("句应为ast.S式句，却得%T",
			程.J句集[0])
	}

	呼式, 是 := 句.S式.(*ast.H呼式)
	if !是 {
		测.Fatalf("句.S式应为ast.H呼式，却得%T",
			句.S式)
	}

	if !测标式(测, 呼式.H函名, "某函") {
		return
	}

	if len(呼式.S实参集) != 3 {
		测.Fatalf("实参数应为3个，却得%d", len(呼式.S实参集))
	}

	测值式(测, 呼式.S实参集[0], 1)
	测中式(测, 呼式.S实参集[1], 2, "*", 3)
	测中式(测, 呼式.S实参集[2], 4, "+", 5)
}

func Test呼式实参群(测 *testing.T) {
	测试集 := []struct {
		待测   string
		期函名  string
		期实参集 []string
	}{
		{
			待测:   "add();",
			期函名:  "add",
			期实参集: []string{},
		},
		{
			待测:   "add(1);",
			期函名:  "add",
			期实参集: []string{"1"},
		},
		{
			待测:   "add(1, 2 * 3, 4 + 5);",
			期函名:  "add",
			期实参集: []string{"1", "(2 * 3)", "(4 + 5)"},
		},
		{
			待测:   "某函();",
			期函名:  "某函",
			期实参集: []string{},
		},
		{
			待测:   "某函(金);",
			期函名:  "某函",
			期实参集: []string{"金"},
		},
		{
			待测:   "某函(金, 木 * 水, 火 + 土);",
			期函名:  "某函",
			期实参集: []string{"金", "(木 * 水)", "(火 + 土)"},
		},
	}

	for _, 测试 := range 测试集 {
		取 := lexer.Z造(测试.待测)
		树 := Z造(取)
		程 := 树.X析程()
		检树障(测, 树)

		句 := 程.J句集[0].(*ast.S式句)
		呼式, 是 := 句.S式.(*ast.H呼式)
		if !是 {
			测.Fatalf("句.S式应为ast.H呼式却得%T",
				句.S式)
		}

		if !测标式(测, 呼式.H函名, 测试.期函名) {
			return
		}

		if len(呼式.S实参集) != len(测试.期实参集) {
			测.Fatalf("期望实参数为%d, 却得%d",
				len(测试.期实参集), len(呼式.S实参集))
		}

		for 号, 实参 := range 测试.期实参集 {
			if 呼式.S实参集[号].String() != 实参 {
				测.Errorf("%d号实参应为%q, 却得%q", 号,
					实参, 呼式.S实参集[号].String())
			}
		}
	}
}

func Test串式群(测 *testing.T) {
	待测 := `"你好，猴子！Hello monkey！";`

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	句 := 程.J句集[0].(*ast.S式句)
	串, 是 := 句.S式.(*ast.C串式)
	if !是 {
		测.Fatalf("串应为*ast.C串式，却得%T", 句.S式)
	}

	if 串.Z值 != "你好，猴子！Hello monkey！" {
		测.Errorf("串.Z值应为%q，却得%q", "你好，猴子！Hello monkey！", 串.Z值)
	}
}

func Test空队群(测 *testing.T) {
	待测 := "[]"

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	句, _ := 程.J句集[0].(*ast.S式句)
	队, 是 := 句.S式.(*ast.D队式)
	if !是 {
		测.Fatalf("队应为ast.D队式，却得%T", 句.S式)
	}

	if len(队.Y元集) != 0 {
		测.Errorf("空队长度应为0，却得%d", len(队.Y元集))
	}
}

func Test队式群(测 *testing.T) {
	待测 := "[1, 2 * 2, 3 + 3]"

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	句, _ := 程.J句集[0].(*ast.S式句)
	队, 是 := 句.S式.(*ast.D队式)
	if !是 {
		测.Fatalf("应为ast.D队式，却得%T", 句.S式)
	}

	if len(队.Y元集) != 3 {
		测.Fatalf("len(队.Y元集)应为3，却得%d", len(队.Y元集))
	}

	测整式(测, 队.Y元集[0], 1)
	测中式(测, 队.Y元集[1], 2, "*", 2)
	测中式(测, 队.Y元集[2], 3, "+", 3)
}

func Test号式群(测 *testing.T) {
	待测 := "车队[1 + 1]"

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	句, _ := 程.J句集[0].(*ast.S式句)
	号式, 是 := 句.S式.(*ast.H号式)
	if !是 {
		测.Fatalf("号式应为*ast.H号式，却得%T", 句.S式)
	}

	if !测标式(测, 号式.Z左式, "车队") {
		return
	}

	if !测中式(测, 号式.H号, 1, "+", 1) {
		return
	}
}

func Test空耦式群(测 *testing.T) {
	待测 := "{}"

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	句 := 程.J句集[0].(*ast.S式句)
	耦式, 是 := 句.S式.(*ast.O耦式)
	if !是 {
		测.Fatalf("耦式应为ast.O耦式，却得%T", 句.S式)
	}

	if len(耦式.O耦集) != 0 {
		测.Errorf("空耦式.O耦集长度应为0，却得%d", len(耦式.O耦集))
	}
}

func Test串键耦式群(测 *testing.T) {
	待测 := `{"甲子": 1, "乙丑": 2, "丙寅": 3}`

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	句 := 程.J句集[0].(*ast.S式句)
	耦式, 是 := 句.S式.(*ast.O耦式)
	if !是 {
		测.Fatalf("耦式应为ast.O耦式，却得%T", 句.S式)
	}

	期望 := map[string]int64{
		"甲子": 1,
		"乙丑": 2,
		"丙寅": 3,
	}

	if len(耦式.O耦集) != len(期望) {
		测.Errorf("耦式.O耦集长度应为3，却得%d", len(耦式.O耦集))
	}

	for 键, 值 := range 耦式.O耦集 {
		串式, 是 := 键.(*ast.C串式)
		if !是 {
			测.Errorf("串键应为ast.C串式，却得%T", 键)
			continue
		}

		期值 := 期望[串式.String()]
		测整式(测, 值, 期值)
	}
}

func Test贰键耦式群(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 map[string]int64
	}{
		{`{true: 1, false: 2}`, map[string]int64{
			"true":  1,
			"false": 2,
		}},
		{`{真: 1, 假: 2}`, map[string]int64{
			"真": 1,
			"假": 2,
		}},
	}

	for _, 测试 := range 测试集 {

		取 := lexer.Z造(测试.待测)
		树 := Z造(取)
		程 := 树.X析程()
		检树障(测, 树)

		句 := 程.J句集[0].(*ast.S式句)
		耦式, 是 := 句.S式.(*ast.O耦式)
		if !是 {
			测.Fatalf("耦式应为ast.O耦式，却得%T", 句.S式)
		}

		if len(耦式.O耦集) != len(测试.期望) {
			测.Errorf("耦式.O耦集长度应为2，却得%d", len(耦式.O耦集))
		}

		for 键, 值 := range 耦式.O耦集 {
			贰式, 是 := 键.(*ast.E贰式)
			if !是 {
				测.Errorf("贰键应为ast.E贰式，却得%T", 键)
				continue
			}

			期值 := 测试.期望[贰式.String()]
			测整式(测, 值, 期值)
		}
	}
}

func Test整键耦式群(测 *testing.T) {
	待测 := `{1: 1, 2: 2, 3: 3}`

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	句 := 程.J句集[0].(*ast.S式句)
	耦式, 是 := 句.S式.(*ast.O耦式)
	if !是 {
		测.Fatalf("耦式应为ast.O耦式，却得%T", 句.S式)
	}

	期望 := map[string]int64{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	if len(耦式.O耦集) != len(期望) {
		测.Errorf("耦式.O耦集长度应为3，却得%d", len(耦式.O耦集))
	}

	for 键, 值 := range 耦式.O耦集 {
		整式, 是 := 键.(*ast.Z整式)
		if !是 {
			测.Errorf("整式键应为ast.Z整式，却得%T", 键)
			continue
		}

		期值 := 期望[整式.String()]

		测整式(测, 值, 期值)
	}
}

func Test式值耦式群(测 *testing.T) {
	待测 := `{"乾卦": 0 + 1, "坤卦": 10 - 8, "屯卦": 15 / 5}`

	取 := lexer.Z造(待测)
	树 := Z造(取)
	程 := 树.X析程()
	检树障(测, 树)

	句 := 程.J句集[0].(*ast.S式句)
	耦式, 是 := 句.S式.(*ast.O耦式)
	if !是 {
		测.Fatalf("耦式应为ast.O耦式，却得%T", 句.S式)
	}

	if len(耦式.O耦集) != 3 {
		测.Errorf("耦式.O耦集长度应为3，却得%d", len(耦式.O耦集))
	}

	测试集 := map[string]func(ast.S式){
		"乾卦": func(e ast.S式) {
			测中式(测, e, 0, "+", 1)
		},
		"坤卦": func(e ast.S式) {
			测中式(测, e, 10, "-", 8)
		},
		"屯卦": func(e ast.S式) {
			测中式(测, e, 15, "/", 5)
		},
	}

	for 键, 值 := range 耦式.O耦集 {
		串键, 是 := 键.(*ast.C串式)
		if !是 {
			测.Errorf("串键应为ast.C串式，却得%T", 键)
			continue
		}

		函值, 是 := 测试集[串键.String()]
		if !是 {
			测.Errorf("%q键函值未找到", 串键.String())
			continue
		}

		函值(值)
	}
}

func 测令句(测 *testing.T, 令 ast.J句, 标名 string) bool {
	if 令.C词文() != "令" && 令.C词文() != "let" {
		测.Errorf("令.C词文应为'令'或'let'，却得%q", 令.C词文())
		return false
	}

	令句, 是 := 令.(*ast.L令句)
	if !是 {
		测.Errorf("s应为*ast.L令句，却得%T", 令)
		return false
	}

	if 令句.M名.Z值 != 标名 {
		测.Errorf("令句.M名.Z值应为'%s'，却得%s", 标名, 令句.M名.Z值)
		return false
	}

	if 令句.M名.C词文() != 标名 {
		测.Errorf("令句.M名应为'%s'，却得%s", 标名, 令句.M名)
		return false
	}

	return true
}

func 测中式(测 *testing.T, 中式 ast.S式, 左式 interface{},
	算符 string, 右式 interface{}) bool {

	算式, 是 := 中式.(*ast.Z中式)
	if !是 {
		测.Errorf("算式应为ast.Z中式，却得%T(%s)", 中式, 中式)
		return false
	}

	if !测值式(测, 算式.Z左式, 左式) {
		return false
	}

	if 算式.S算符 != 算符 {
		测.Errorf("中式.S算符应为'%s'，却得%q", 算符, 算式.S算符)
		return false
	}

	if !测值式(测, 算式.Y右式, 右式) {
		return false
	}

	return true
}

func 测值式(
	测 *testing.T,
	值式 ast.S式,
	期望 interface{},
) bool {
	switch v := 期望.(type) {
	case int:
		return 测整式(测, 值式, int64(v))
	case int64:
		return 测整式(测, 值式, v)
	case string:
		return 测标式(测, 值式, v)
	case bool:
		return 测贰式(测, 值式, v)
	}
	测.Errorf("%T类式未处理", 值式)
	return false
}

func 测整式(测 *testing.T, 整 ast.S式, 值 int64) bool {
	整式, 是 := 整.(*ast.Z整式)
	if !是 {
		测.Errorf("应为*ast.Z整式，却得%T", 整)
		return false
	}

	if 整式.Z值 != 值 {
		测.Errorf("整式.Z值应为%d，却得%d", 值, 整式.Z值)
		return false
	}

	if 整式.C词文() != fmt.Sprintf("%d", 值) {
		测.Errorf("整式.C词文应为%d，却得%s", 值,
			整式.C词文())
		return false
	}

	return true
}

func 测标式(测 *testing.T, 标 ast.S式, 值 string) bool {
	标式, 是 := 标.(*ast.B标式)
	if !是 {
		测.Errorf("应为*ast.B标式，却得%T", 标)
		return false
	}

	if 标式.Z值 != 值 {
		测.Errorf("标式.Z值应为%s，却得%s", 值, 标式.Z值)
		return false
	}

	if 标式.C词文() != 值 {
		测.Errorf("标式.C词文应为%s，却得%s", 值,
			标式.C词文())
		return false
	}

	return true
}

func 测贰式(测 *testing.T, 贰 ast.S式, 值 bool) bool {
	贰式, 是 := 贰.(*ast.E贰式)
	if !是 {
		测.Errorf("应为*ast.E贰式，却得%T", 贰)
		return false
	}

	if 贰式.Z值 != 值 {
		测.Errorf("贰式.Z值应为%t，却得%t", 值, 贰式.Z值)
		return false
	}

	if 值 && 贰式.C词文() != "真" && 贰式.C词文() != "true" ||
		!值 && 贰式.C词文() != "假" && 贰式.C词文() != "false" {
		测.Errorf("贰式.C词文应为%t，却得%s",
			值, 贰式.C词文())
		return false
	}

	return true
}

func 检树障(测 *testing.T, 树 *S树句器) {
	障集 := 树.Z障集()
	if len(障集) == 0 {
		return
	}

	测.Errorf("树句器有%d障", len(障集))
	for _, 障 := range 障集 {
		测.Errorf("树障: %q", 障)
	}
	测.FailNow()
}
