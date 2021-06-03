package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = "猴语:>> "

func H会话(入 io.Reader, 出 io.Writer) {
	察器 := bufio.NewScanner(入)

	for {
		fmt.Print(PROMPT)
		察完 := 察器.Scan()
		if !察完 {
			return
		}

		行 := 察器.Text()
		取 := lexer.Q造(行)

		for 词 := 取.Q翌词(); 词.C类 != token.T档尾; 词 = 取.Q翌词() {
			fmt.Printf("词类：%-6v\t词：%-20v\n", 词.C类, 词.C值)
		}
	}
}
