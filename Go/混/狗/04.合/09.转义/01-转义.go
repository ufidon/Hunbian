/* html/template默认自动转义保留字符，也可禁止其转义
* 跑：
* ./01-转义 >转义.html
 */

package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const 模板1 = `<p>甲: {{.S甲}}</p><p>乙: {{.S乙}}</p>`
	模 := template.Must(template.New("escape").Parse(模板1))

	var 数据 struct {
		S甲 string        // 无自动转义，有安全漏洞
		S乙 template.HTML // 自动转义
	}

	数据.S甲 = "<b>自动转义禁止</b>"
	数据.S乙 = "<b>自动转义启用</b>"

	if 障 := 模.Execute(os.Stdout, 数据); 障 != nil {
		log.Fatal(障)
	}
}
