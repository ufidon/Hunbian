package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	程 := &C程{
		J句集: []J句{
			&L令句{
				C词: token.C词{L类: token.L令, W文: "令"},
				M名: &B标式{
					C词: token.C词{L类: token.B标识, W文: "变量"},
					Z值: "变量",
				},
				Z值: &B标式{
					C词: token.C词{L类: token.B标识, W文: "身高"},
					Z值: "身高",
				},
			},
		},
	}

	if 程.String() != "令 变量 = 身高;" {
		t.Errorf("程.String() 错,得%q", 程.String())
	}
}
