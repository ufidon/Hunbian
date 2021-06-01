/* 网页算式绘图
* 跑：./01-网图
* 在浏览器中输入：http://localhost:12345/绘?式=正弦(-x)*幂(1.5,-r)
 */

package main

import (
	"fmt"
	"io"
	"jsq/ss"
	"log"
	"math"
	"net/http"
)

const (
	画宽, 画高 = 600, 320     // 画布宽高,单位像素
	网格数    = 100          //网格数
	xy域    = 30.0         // xy轴范围， 定义域 (-xy域..+xy域)
	xy刻度   = 画宽 / 2 / xy域 // x, y轴每单位像素数
	z刻度    = 画高 * 0.4     // z轴每单位像素数
	xy夹角   = math.Pi / 6  // x,y轴夹角(=30°)
)

var sin30, cos30 = math.Sin(xy夹角), math.Cos(xy夹角)

func 曲面图(出 io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(出, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill:white; stroke-width:0.7' "+
		"width='%d' height='%d'>", 画宽, 画高)
	for i := 0; i < 网格数; i++ {
		for j := 0; j < 网格数; j++ {
			ax, ay := 隅(f, i+1, j)
			bx, by := 隅(f, i, j)
			cx, cy := 隅(f, i, j+1)
			dx, dy := 隅(f, i+1, j+1)
			fmt.Fprintf(出, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(出, "</svg>")
}

func 隅(f func(x, y float64) float64, i, j int) (float64, float64) {
	// 确定网格隅点坐标
	x := xy域 * (float64(i)/网格数 - 0.5)
	y := xy域 * (float64(j)/网格数 - 0.5)

	// 算函数值z
	z := f(x, y)

	// 将(x,y,z)等距投影到2维的SVG画布上(sx,sy)
	sx := 画宽/2 + (x-y)*cos30*xy刻度
	sy := 画高/2 + (x+y)*sin30*xy刻度 - z*z刻度
	return sx, sy
}

func 析检(串 string) (ss.S式, error) {
	if 串 == "" {
		return nil, fmt.Errorf("空算式")
	}
	式, 障 := ss.S析(串)
	if 障 != nil {
		return nil, 障
	}
	变量集 := make(map[ss.S变量]bool)
	if 障 := 式.S检(变量集); 障 != nil {
		return nil, 障
	}
	for 量 := range 变量集 {
		if 量 != "x" && 量 != "y" && 量 != "r" {
			return nil, fmt.Errorf("未定义变量：%s", 量)
		}
	}
	return 式, 障
}

func 绘(回 http.ResponseWriter, 请 *http.Request) {
	请.ParseForm()
	算式 := 请.Form.Get("式")
	log.Printf("您输入的算式是:%s\n", 算式)

	式, 障 := 析检(算式)
	if 障 != nil {
		http.Error(回, "错误算式："+障.Error(), http.StatusBadRequest)
		return
	}

	回.Header().Set("Content-type", "image/svg+xml")
	曲面图(回, func(x, y float64) float64 {
		r := math.Hypot(x, y) // 距至(0,0)
		return 式.S算(ss.S境{"x": x, "y": y, "r": r})
	})
}

func main() {
	http.HandleFunc("/绘", 绘)
	log.Fatal(http.ListenAndServe("localhost:12345", nil))
}
