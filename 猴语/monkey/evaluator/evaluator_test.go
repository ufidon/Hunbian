package evaluator

import (
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)

func Test估整式(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, 测试 := range 测试集 {
		估得 := 测估(测试.待测)
		测整盒(测, 估得, 测试.期望)
	}
}

func Test估贰式(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"真", true},
		{"假", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"真 == 真", true},
		{"假 == 假", true},
		{"真 == 假", false},
		{"真 != 假", true},
		{"假 != 真", true},
		{"(1 < 2) == 真", true},
		{"(1 < 2) == 假", false},
		{"(1 > 2) == 真", false},
		{"(1 > 2) == 假", true},
	}

	for _, 测试 := range 测试集 {
		估得 := 测估(测试.待测)
		测贰盒(测, 估得, 测试.期望)
	}
}

func Test非算符(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"!真", false},
		{"!假", true},
		{"!!真", true},
		{"!!假", false},
	}

	for _, 测试 := range 测试集 {
		估得 := 测估(测试.待测)
		测贰盒(测, 估得, 测试.期望)
	}
}

func Test若式(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"若 (真) { 10 }", 10},
		{"若 (假) { 10 }", nil},
		{"若 (1) { 10 }", 10},
		{"若 (1 < 2) { 10 }", 10},
		{"若 (1 > 2) { 10 }", nil},
		{"若 (1 > 2) { 10 } 否 { 20 }", 20},
		{"若 (1 < 2) { 10 } 否 { 20 }", 10},
	}

	for _, 测试 := range 测试集 {
		估得 := 测估(测试.待测)
		整数, 是 := 测试.期望.(int)
		if 是 {
			测整盒(测, 估得, int64(整数))
		} else {
			测空盒(测, 估得)
		}
	}
}

func Test返句(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{"if (10 > 1) { return 10; }", 10},
		{
			`
if (10 > 1) {
  if (10 > 1) {
    return 10;
  }

  return 1;
}
`,
			10,
		},
		{
			`
let f = fn(x) {
  return x;
  x + 10;
};
f(10);`,
			10,
		},
		{
			`
let f = fn(x) {
   let result = x + 10;
   return result;
   return 10;
};
f(10);`,
			20,
		},
		{"返 10;", 10},
		{"返 10; 9;", 10},
		{"返 2 * 5; 9;", 10},
		{"9; 返 2 * 5; 9;", 10},
		{"若 (10 > 1) { 返 10; }", 10},
		{
			`
若 (10 > 1) {
  若 (10 > 1) {
    返 10;
  }

  返 1;
}
`,
			10,
		},

		{
			`令 甲函 = 函(甲) {
					  返 甲;
					  甲 + 10;
					};
					甲函(10);`,
			10,
		},
		{
			`
				令 乙函 = 函(乙) {
				   令 果 = 乙 + 10;
				   返 果;
				   返 10;
				};
				乙函(10);`,
			20,
		},
	}

	for _, 测试 := range 测试集 {
		估得 := 测估(测试.待测)
		测整盒(测, 估得, 测试.期望)
	}
}

func Test理障(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期志 string
	}{
		{
			"5 + true;",
			"型不配: 整盒 + 贰盒",
		},
		{
			"5 + true; 5;",
			"型不配: 整盒 + 贰盒",
		},
		{
			"-true",
			"未知算符: -贰盒",
		},
		{
			"true + false;",
			"未知算符: 贰盒 + 贰盒",
		},
		{
			"true + false + true + false;",
			"未知算符: 贰盒 + 贰盒",
		},
		{
			"5; true + false; 5",
			"未知算符: 贰盒 + 贰盒",
		},
		{
			`"Hello" - "World"`,
			"未知算符: 串盒 - 串盒",
		},
		{
			"if (10 > 1) { true + false; }",
			"未知算符: 贰盒 + 贰盒",
		},
		{
			`
if (10 > 1) {
  if (10 > 1) {
    return true + false;
  }

  return 1;
}
`,
			"未知算符: 贰盒 + 贰盒",
		},
		{
			"foobar",
			"未知标识: foobar",
		},
		{
			`{"name": "Monkey"}[fn(x) { x }];`,
			"函盒不可用作耦键",
		},
		{
			`999[1]`,
			"整盒号型不支持",
		},
	}

	for _, 测试 := range 测试集 {
		估得 := 测估(测试.待测)

		障盒, 是 := 估得.(*object.Z障型)
		if !是 {
			测.Errorf("no error object returned. got=%T(%+v)",
				估得, 估得)
			continue
		}

		if 障盒.Z志 != 测试.期志 {
			测.Errorf("wrong error message. 期望=%q, got=%q",
				测试.期志, 障盒.Z志)
		}
	}
}

func Test令句(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
		{"令 甲 = 5; 甲;", 5},
		{"令 甲 = 5 * 5; 甲;", 25},
		{"令 甲 = 5; 令 乙 = 甲; 乙;", 5},
		{"令 甲 = 5; 令 乙 = 甲; 令 丙 = 甲 + 乙 + 5; 丙;", 15},
	}

	for _, 测试 := range 测试集 {
		测整盒(测, 测估(测试.待测), 测试.期望)
	}
}

func Test函盒(测 *testing.T) {
	待测 := "fn(x) { x + 2; };"
	//待测 := "函(x) { x + 2; };"

	估得 := 测估(待测)
	次函, 是 := 估得.(*object.H函型)
	if !是 {
		测.Fatalf("object is not H函型. got=%T (%+v)", 估得, 估得)
	}

	if len(次函.X形参集) != 1 {
		测.Fatalf("function has wrong parameters. X形参集=%+v",
			次函.X形参集)
	}

	if 次函.X形参集[0].String() != "x" {
		测.Fatalf("parameter is not 'x'. got=%q", 次函.X形参集[0])
	}

	期体 := "(x + 2)"

	if 次函.T体.String() != 期体 {
		测.Fatalf("body is not %q. got=%q", 期体, 次函.T体.String())
	}
}

func Test用函(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 int64
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},
		{"令 恒等 = 函(子) { 子; }; 恒等(5);", 5},
		{"令 恒等 = 函(子) { 返 子; }; 恒等(5);", 5},
		{"令 翻番 = 函(子) { 子 * 2; }; 翻番(5);", 10},
		{"令 求和 = 函(子, 丑) { 子 + 丑; }; 求和(5, 5);", 10},
		{"令 求和 = 函(子, 丑) { 子 + 丑; }; 求和(5 + 5, 求和(5, 5));", 20},
		{"函(子) { 子; }(5)", 5},
	}

	for _, 测试 := range 测试集 {
		测整盒(测, 测估(测试.待测), 测试.期望)
	}
}

func Test闭境(测 *testing.T) {
	待测集 := []string{
		`
let first = 10;
let second = 10;
let third = 10;

let ourFunction = fn(first) {
  let second = 20;

  first + second + third;
};

ourFunction(20) + first + second;`,
		`
令 甲 = 10;
令 乙 = 10;
令 丙 = 10;

令 某函 = fn(甲) {
  令 乙 = 20;

  甲 + 乙 + 丙;
};

某函(20) + 甲 + 乙;`,
	}
	for _, 待测 := range 待测集 {
		测整盒(测, 测估(待测), 70)
	}

}

func Test闭(测 *testing.T) {
	待测集 := []string{`
let newAdder = fn(x) {
  fn(y) { x + y };
};

let addTwo = newAdder(2);
addTwo(2);`,
		`
令 含参函 = 函(函参) {
	函(甲) { 函参 + 甲 };
  };
  
  令 用参函 = 含参函(2);
  用参函(2);`,
	}
	for _, 待测 := range 待测集 {
		测整盒(测, 测估(待测), 4)
	}

}

func Test串式(测 *testing.T) {
	待测集 := []struct {
		输入 string
		期望 string
	}{
		{`"Hello World!"`, "Hello World!"},
		{`"你好，世界！"`, "你好，世界！"},
	}

	for _, 待测 := range 待测集 {
		估得 := 测估(待测.输入)
		串, 是 := 估得.(*object.C串型)
		if !是 {
			测.Fatalf("object is not C串型. got=%T (%+v)", 估得, 估得)
		}

		if 串.Z值 != 待测.期望 {
			测.Errorf("C串型 has wrong value. got=%q", 串.Z值)
		}
	}

}

func Test串串(测 *testing.T) {
	待测 := `"Hello" + " " + "World!"`

	估得 := 测估(待测)
	串, 是 := 估得.(*object.C串型)
	if !是 {
		测.Fatalf("object is not C串型. got=%T (%+v)", 估得, 估得)
	}

	if 串.Z值 != "Hello World!" {
		测.Errorf("C串型 has wrong value. got=%q", 串.Z值)
	}
}

func Test内函(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 interface{}
	}{
		{`节数("")`, 0},
		{`节数("four")`, 4},
		{`节数("hello world")`, 11},
		{`节数(1)`, "节数不支持整盒"},
		{`节数("one", "two")`, "需一个参数，不是2个"},
		{`节数([1, 2, 3])`, 3},
		{`节数([])`, 0},
		{`字数("")`, 0},
		{`字数("你好")`, 2},
		{`字数("hello world")`, 11},
		{`字数(1)`, "节数不支持整盒"},
		{`字数("你好", "世界")`, "需一个参数，不是2个"},
		{`字数([1, 2, 3])`, 3},
		{`字数([])`, 0},
		{`显示("hello", "world!")`, nil},
		{`首元([1, 2, 3])`, 1},
		{`首元([])`, nil},
		{`首元(1)`, "首元参数应为队而不是整盒"},
		{`尾元([1, 2, 3])`, 3},
		{`尾元([])`, nil},
		{`尾元(1)`, "尾元参数应为队而不是整盒"},
		{`掐头([1, 2, 3])`, []int{2, 3}},
		{`掐头([])`, nil},
		{`附尾([], 1)`, []int{1}},
		{`附尾(1, 1)`, "附尾参数应为队而不是整盒"},
	}

	for _, 测试 := range 测试集 {
		估得 := 测估(测试.待测)

		switch 期望 := 测试.期望.(type) {
		case int:
			测整盒(测, 估得, int64(期望))
		case nil:
			测空盒(测, 估得)
		case string:
			障盒, 是 := 估得.(*object.Z障型)
			if !是 {
				测.Errorf("object is not Z障型. got=%T (%+v)",
					估得, 估得)
				continue
			}
			if 障盒.Z志 != 期望 {
				测.Errorf("wrong error message. 期望=%q, got=%q",
					期望, 障盒.Z志)
			}
		case []int:
			队, 是 := 估得.(*object.D队型)
			if !是 {
				测.Errorf("obj not D队型. got=%T (%+v)", 估得, 估得)
				continue
			}

			if len(队.Y元集) != len(期望) {
				测.Errorf("wrong num of elements. want=%d, got=%d",
					len(期望), len(队.Y元集))
				continue
			}

			for 号, 期元 := range 期望 {
				测整盒(测, 队.Y元集[号], int64(期元))
			}
		}
	}
}

func Test队式(测 *testing.T) {
	待测 := "[1, 2 * 2, 3 + 3]"

	估得 := 测估(待测)
	果, 是 := 估得.(*object.D队型)
	if !是 {
		测.Fatalf("object is not D队型. got=%T (%+v)", 估得, 估得)
	}

	if len(果.Y元集) != 3 {
		测.Fatalf("array has wrong num of elements. got=%d",
			len(果.Y元集))
	}

	测整盒(测, 果.Y元集[0], 1)
	测整盒(测, 果.Y元集[1], 4)
	测整盒(测, 果.Y元集[2], 6)
}

func Test队号式(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"let i = 0; [1][i];",
			1,
		},
		{
			"[1, 2, 3][1 + 1];",
			3,
		},
		{
			"let myArray = [1, 2, 3]; myArray[2];",
			3,
		},
		{
			"let myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];",
			6,
		},
		{
			"let myArray = [1, 2, 3]; let i = myArray[0]; myArray[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
	}

	for _, 测试 := range 测试集 {
		估得 := 测估(测试.待测)
		整数, 是 := 测试.期望.(int)
		if 是 {
			测整盒(测, 估得, int64(整数))
		} else {
			测空盒(测, 估得)
		}
	}
}

func Test耦式(测 *testing.T) {
	待测 := `let two = "two";
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 6 / 2,
		4: 4,
		true: 5,
		false: 6
	}`

	估得 := 测估(待测)
	果, 是 := 估得.(*object.O耦型)
	if !是 {
		测.Fatalf("G估 didn'测 return O耦型. got=%T (%+v)", 估得, 估得)
	}

	期望 := map[object.O耦键]int64{
		(&object.C串型{Z值: "one"}).O耦键():   1,
		(&object.C串型{Z值: "two"}).O耦键():   2,
		(&object.C串型{Z值: "three"}).O耦键(): 3,
		(&object.Z整型{Z值: 4}).O耦键():       4,
		K真值.O耦键():                        5,
		K假值.O耦键():                        6,
	}

	if len(果.S双) != len(期望) {
		测.Fatalf("O耦型 has wrong num of pairs. got=%d", len(果.S双))
	}

	for 期键, 期值 := range 期望 {
		双, 有 := 果.S双[期键]
		if !有 {
			测.Errorf("no pair for given key in S双")
		}

		测整盒(测, 双.Z值, 期值)
	}
}

func Test耦号式(测 *testing.T) {
	测试集 := []struct {
		待测 string
		期望 interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`let key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{true: 5}[true]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
	}

	for _, 测试 := range 测试集 {
		估得 := 测估(测试.待测)
		整数, 是 := 测试.期望.(int)
		if 是 {
			测整盒(测, 估得, int64(整数))
		} else {
			测空盒(测, 估得)
		}
	}
}
func 测估(待测 string) object.H盒 {
	取 := lexer.Z造(待测)
	树 := parser.Z造(取)
	程 := 树.X析程()
	境 := object.Z造境()

	return G估(程, 境)
}

func 测整盒(测 *testing.T, 盒 object.H盒, 期望 int64) bool {
	果, 是 := 盒.(*object.Z整型)
	if !是 {
		测.Errorf("%T (%+v)非整型", 盒, 盒)
		return false
	}
	if 果.Z值 != 期望 {
		测.Errorf("盒含错值%d, 应为%d",
			果.Z值, 期望)
		return false
	}

	return true
}

func 测贰盒(测 *testing.T, 盒 object.H盒, 期望 bool) bool {
	果, 是 := 盒.(*object.E贰型)
	if !是 {
		测.Errorf("%T (%+v)非贰型", 盒, 盒)
		return false
	}
	if 果.Z值 != 期望 {
		测.Errorf("得%t, 期望%t",
			果.Z值, 期望)
		return false
	}
	return true
}

func 测空盒(测 *testing.T, 盒 object.H盒) bool {
	if 盒 != K空值 {
		测.Errorf("%T (%+v)非空值", 盒, 盒)
		return false
	}
	return true
}
