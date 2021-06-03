package lexer

import (
	"testing"

	"monkey/token"
)

func TestQ翌词(测 *testing.T) {
	输入 := `令 五 = 5;
令 十 = 10;

令 加 = 函(x, y) {
  x + y;
};

令 结果 = 加(五, 十);
!-/*5;
5 < 10 > 5;

若 (5 < 10) {
	返 真;
} 别 {
	返 假;
}

10 == 10;
10 != 9;
`

	测试集 := []struct {
		期类 token.C词类
		期值 string
	}{
		{token.C令, "令"},
		{token.X标识, "五"},
		{token.S赋, "="},
		{token.X整型, "5"},
		{token.J分号, ";"},
		{token.C令, "令"},
		{token.X标识, "十"},
		{token.S赋, "="},
		{token.X整型, "10"},
		{token.J分号, ";"},
		{token.C令, "令"},
		{token.X标识, "加"},
		{token.S赋, "="},
		{token.C函, "函"},
		{token.J左括, "("},
		{token.X标识, "x"},
		{token.J逗号, ","},
		{token.X标识, "y"},
		{token.J右括, ")"},
		{token.J左曲, "{"},
		{token.X标识, "x"},
		{token.S加, "+"},
		{token.X标识, "y"},
		{token.J分号, ";"},
		{token.J右曲, "}"},
		{token.J分号, ";"},
		{token.C令, "令"},
		{token.X标识, "结果"},
		{token.S赋, "="},
		{token.X标识, "加"},
		{token.J左括, "("},
		{token.X标识, "五"},
		{token.J逗号, ","},
		{token.X标识, "十"},
		{token.J右括, ")"},
		{token.J分号, ";"},
		{token.S非, "!"},
		{token.S减, "-"},
		{token.S除, "/"},
		{token.S乘, "*"},
		{token.X整型, "5"},
		{token.J分号, ";"},
		{token.X整型, "5"},
		{token.S小于, "<"},
		{token.X整型, "10"},
		{token.S大于, ">"},
		{token.X整型, "5"},
		{token.J分号, ";"},
		{token.C若, "若"},
		{token.J左括, "("},
		{token.X整型, "5"},
		{token.S小于, "<"},
		{token.X整型, "10"},
		{token.J右括, ")"},
		{token.J左曲, "{"},
		{token.C返, "返"},
		{token.C真, "真"},
		{token.J分号, ";"},
		{token.J右曲, "}"},
		{token.C别, "别"},
		{token.J左曲, "{"},
		{token.C返, "返"},
		{token.C假, "假"},
		{token.J分号, ";"},
		{token.J右曲, "}"},
		{token.X整型, "10"},
		{token.S等于, "=="},
		{token.X整型, "10"},
		{token.J分号, ";"},
		{token.X整型, "10"},
		{token.S不等, "!="},
		{token.X整型, "9"},
		{token.J分号, ";"},
		{token.T档尾, ""},
	}

	取 := Q造(输入)

	for 号, 测试 := range 测试集 {
		词 := 取.Q翌词()

		if 词.C类 != 测试.期类 {
			测.Fatalf("测试集[%d] - 词类错误. 期望: %q, 却得: %q",
				号, 测试.期类, 词.C类)
		}

		if 词.C值 != 测试.期值 {
			测.Fatalf("测试集[%d] - 词值错误. 期望: %q, 却得: %q",
				号, 测试.期值, 词.C值)
		}
	}
}
