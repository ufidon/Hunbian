package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

const 提示 = "🐒语:>> "

func H会话(入 io.Reader, 出 io.Writer) {
	览器 := bufio.NewScanner(入)
	境 := object.Z造境()

	for {
		fmt.Printf(提示)
		取完 := 览器.Scan()
		if !取完 {
			return
		}

		行 := 览器.Text()
		取 := lexer.Z造(行)
		树 := parser.Z造(取)

		程 := 树.X析程()
		if len(树.Z障集()) != 0 {
			印树障(出, 树.Z障集())
			continue
		}

		估完 := evaluator.G估(程, 境)
		if 估完 != nil {
			io.WriteString(出, 估完.C察值())
			io.WriteString(出, "\n")
		}
	}
}

const 猴脸 = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func 印树障(出 io.Writer, 树障 []string) {
	io.WriteString(出, 猴脸)
	io.WriteString(出, "糟糕！猴子🐵蒙了🙈请说🐒语，勿说☺话。\n")
	io.WriteString(出, "语法错误:\n")
	for _, 障 := range 树障 {
		io.WriteString(出, "\t"+障+"\n")
	}
}
