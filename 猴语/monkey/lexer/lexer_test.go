package lexer

import (
	"testing"

	"monkey/token"
)

func TestQ取词(测 *testing.T) {
	待测 := `let five = 5;
令 五 = 5;	
let ten = 10;
令 十 = 10;

let add = fn(x, y) {
  x + y;
};
令 加 = 函(x,y){
	x+y;
};

let result = add(five, ten);
令 结果 = 加(五,十);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}
若(5<10){
	返 true;
}否{
	返 false;
}

10 == 10;
10 != 9;
"foobar"
"身高"
"foo bar"
"身 高"
[1, 2];
{"foo": "bar"}
{"身高": 172}
`

	测试集 := []struct {
		期词类 token.C词类
		期词文 string
	}{
		{token.L令, "let"},
		{token.B标识, "five"},
		{token.F赋, "="},
		{token.Z整型, "5"},
		{token.F分号, ";"},
		{token.L令, "令"},
		{token.B标识, "五"},
		{token.F赋, "="},
		{token.Z整型, "5"},
		{token.F分号, ";"},
		{token.L令, "let"},
		{token.B标识, "ten"},
		{token.F赋, "="},
		{token.Z整型, "10"},
		{token.F分号, ";"},
		{token.L令, "令"},
		{token.B标识, "十"},
		{token.F赋, "="},
		{token.Z整型, "10"},
		{token.F分号, ";"},
		{token.L令, "let"},
		{token.B标识, "add"},
		{token.F赋, "="},
		{token.H函, "fn"},
		{token.Z左括, "("},
		{token.B标识, "x"},
		{token.D逗号, ","},
		{token.B标识, "y"},
		{token.Y右括, ")"},
		{token.Z左曲, "{"},
		{token.B标识, "x"},
		{token.J加, "+"},
		{token.B标识, "y"},
		{token.F分号, ";"},
		{token.Y右曲, "}"},
		{token.F分号, ";"},
		{token.L令, "令"},
		{token.B标识, "加"},
		{token.F赋, "="},
		{token.H函, "函"},
		{token.Z左括, "("},
		{token.B标识, "x"},
		{token.D逗号, ","},
		{token.B标识, "y"},
		{token.Y右括, ")"},
		{token.Z左曲, "{"},
		{token.B标识, "x"},
		{token.J加, "+"},
		{token.B标识, "y"},
		{token.F分号, ";"},
		{token.Y右曲, "}"},
		{token.F分号, ";"},
		{token.L令, "let"},
		{token.B标识, "result"},
		{token.F赋, "="},
		{token.B标识, "add"},
		{token.Z左括, "("},
		{token.B标识, "five"},
		{token.D逗号, ","},
		{token.B标识, "ten"},
		{token.Y右括, ")"},
		{token.F分号, ";"},
		{token.L令, "令"},
		{token.B标识, "结果"},
		{token.F赋, "="},
		{token.B标识, "加"},
		{token.Z左括, "("},
		{token.B标识, "五"},
		{token.D逗号, ","},
		{token.B标识, "十"},
		{token.Y右括, ")"},
		{token.F分号, ";"},
		{token.F非, "!"},
		{token.J减, "-"},
		{token.C除, "/"},
		{token.C乘, "*"},
		{token.Z整型, "5"},
		{token.F分号, ";"},
		{token.Z整型, "5"},
		{token.X小于, "<"},
		{token.Z整型, "10"},
		{token.D大于, ">"},
		{token.Z整型, "5"},
		{token.F分号, ";"},
		{token.R若, "if"},
		{token.Z左括, "("},
		{token.Z整型, "5"},
		{token.X小于, "<"},
		{token.Z整型, "10"},
		{token.Y右括, ")"},
		{token.Z左曲, "{"},
		{token.F返, "return"},
		{token.Z真, "true"},
		{token.F分号, ";"},
		{token.Y右曲, "}"},
		{token.F否, "else"},
		{token.Z左曲, "{"},
		{token.F返, "return"},
		{token.J假, "false"},
		{token.F分号, ";"},
		{token.Y右曲, "}"},
		{token.R若, "若"},
		{token.Z左括, "("},
		{token.Z整型, "5"},
		{token.X小于, "<"},
		{token.Z整型, "10"},
		{token.Y右括, ")"},
		{token.Z左曲, "{"},
		{token.F返, "返"},
		{token.Z真, "true"},
		{token.F分号, ";"},
		{token.Y右曲, "}"},
		{token.F否, "否"},
		{token.Z左曲, "{"},
		{token.F返, "返"},
		{token.J假, "false"},
		{token.F分号, ";"},
		{token.Y右曲, "}"},
		{token.Z整型, "10"},
		{token.D等于, "=="},
		{token.Z整型, "10"},
		{token.F分号, ";"},
		{token.Z整型, "10"},
		{token.B不等, "!="},
		{token.Z整型, "9"},
		{token.F分号, ";"},
		{token.C串型, "foobar"},
		{token.C串型, "身高"},
		{token.C串型, "foo bar"},
		{token.C串型, "身 高"},
		{token.Z左方, "["},
		{token.Z整型, "1"},
		{token.D逗号, ","},
		{token.Z整型, "2"},
		{token.Y右方, "]"},
		{token.F分号, ";"},
		{token.Z左曲, "{"},
		{token.C串型, "foo"},
		{token.M冒号, ":"},
		{token.C串型, "bar"},
		{token.Y右曲, "}"},
		{token.Z左曲, "{"},
		{token.C串型, "身高"},
		{token.M冒号, ":"},
		{token.Z整型, "172"},
		{token.Y右曲, "}"},
		{token.D档尾, ""},
	}

	取 := Z造(待测)

	for 号, 测试 := range 测试集 {
		词 := 取.Q取词()

		if 词.L类 != 测试.期词类 {
			测.Fatalf("%d号测试 词类应为%q, 却是%q",
				号, 测试.期词类, 词.L类)
		}

		if 词.W文 != 测试.期词文 {
			测.Fatalf("%d号测试 词文应为%q, 却是%q",
				号, 测试.期词文, 词.W文)
		}
	}
}
