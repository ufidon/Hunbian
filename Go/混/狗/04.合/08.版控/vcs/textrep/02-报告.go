/* 测试Github之版控狗宝
 */
package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"vcs/github"
)

const 模板 = `总故障数：{{.Z障数}} 
{{range .Z障集}}-------------------------------------------------
障号：		{{.S数}}
用户：		{{.Y用户.D登录名}}
标题：		{{.B标题 | printf "%.64s"}}
天数:		{{.C创建日 | 几天前}}天
{{end}}`

func 几天前(时 time.Time) int {
	return int(time.Since(时).Hours() / 24)
}

func main() {
	// 创建文本报告模板
	// 法1：
	// 报告, 障 := template.New("障报").
	// 	Funcs(template.FuncMap{"几天前": 几天前}).
	// 	Parse(模板)
	// if 障 != nil {
	// 	log.Fatal(障)
	// }
	// 法2：Must 带障理(故障处理)
	var 报告 = template.Must(template.New("障报").
		Funcs(template.FuncMap{"几天前": 几天前}).
		Parse(模板))

	果, 障 := github.X寻障(os.Args[1:])
	if 障 != nil {
		log.Fatal(障)
	}
	if 障 := 报告.Execute(os.Stdout, 果); 障 != nil {
		log.Fatal(障)
	}
}
