/* 测试Github之版控狗宝
* 运行：
* 1) ./02-报告 repo:golang/go commenter:gopherbot json encoder >报告.html
* html/template自动处理(转义) HTML 保留符号如 &,>, <等，如报告2.html的标题
* 2) ./02-报告 repo:golang/go 3133 10535 >报告2.html
 */
package main

import (
	"html/template"
	"log"
	"os"

	"vcs/github"
)

func main() {

	// 创建网页报告模板
	var 报模1 = template.Must(template.New("障报").Parse(`
	<h1>共{{.Z障数}}个故障</h1>
	<table>
	<tr style='text-align: left'>
		<th>号</th>
		<th>状态</th>
		<th>用户</th>
		<th>标题</th>
	</tr>
	{{range .Z障集}}
	<tr>
		<td><a href='{{.W网址}}'>{{.S数}}</a></td>
		<td>{{.Z状态}}</td>
		<td><a href='{{.Y用户.W网址}}'>{{.Y用户.D登录名}}</a></td>
		<td><a href='{{.W网址}}'>{{.B标题}}</a></td>
	</tr>
	{{end}}
	</table>
	`))

	果, 障 := github.X寻障(os.Args[1:])
	if 障 != nil {
		log.Fatal(障)
	}
	if 障 := 报模1.Execute(os.Stdout, 果); 障 != nil {
		log.Fatal(障)
	}
}
