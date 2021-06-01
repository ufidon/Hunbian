/* 画三维图
z = sin(r/r)
r = x^2+y^2
参考：
1. 等轴测投影：https://zh.wikipedia.org/zh-cn/等轴测投影
2. svg多边形：https://www.w3schools.com/graphics/svg_polygon.asp
3. svg颜色：https://css-tricks.com/almanac/properties/s/stroke/
*/
package main

import (
	"fmt"
	"math"
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

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"width='%d' height='%d'>", 画宽, 画高)
	for i := 0; i < 网格数; i++ {
		for j := 0; j < 网格数; j++ {
			ax, ay, _ := 隅(i+1, j)
			bx, by, z := 隅(i, j)
			cx, cy, _ := 隅(i, j+1)
			dx, dy, _ := 隅(i+1, j+1)
			s := 色(z)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'  style='stroke:#%s; fill:#%s; stroke-width:0.7' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, s, s)
		}
	}
	fmt.Printf("</svg>")
}

func 隅(i, j int) (float64, float64, float64) {
	// 确定网格隅点坐标
	x := xy域 * (float64(i)/网格数 - 0.5)
	y := xy域 * (float64(j)/网格数 - 0.5)

	// 算函数值z
	z := f(x, y)

	// 将(x,y,z)等距投影到2维的SVG画布上(sx,sy)
	sx := 画宽/2 + (x-y)*cos30*xy刻度
	sy := 画高/2 + (x+y)*sin30*xy刻度 - z*z刻度
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // r 为(x,y)至原点(0,0)的距离
	return math.Sin(r) / r
}

func 色(z float64) string {
	/*
		// 法1
		// 把区间[-0.2172336,1]映射至[0x0000ff, 0xff0000]
		s := float64(0xff) + float64(0xff0000-0xff)*z/1.2172336
		return fmt.Sprintf("%06X", uint(s))
	*/
	/* 法2：https://www.schemecolor.com/blue-to-red-gradient.php
	-0.21723362821118974
	-0.01436135684265813
	0.18851091452587349
	0.39138318589440513
	0.5942554572629367
	0.7971277286314684
	1.0
	*/
	switch {
	case z <= -0.01436135684265813:
		return "1034A6"
	case -0.01436135684265813 < z && z <= 0.18851091452587349:
		return "412F88"
	case 0.18851091452587349 < z && z <= 0.39138318589440513:
		return "722B6A"
	case 0.39138318589440513 < z && z <= 0.5942554572629367:
		return "A2264B"
	case 0.5942554572629367 < z && z <= 0.7971277286314684:
		return "D3212D"
	default:
		return "F62D2D"
	}
}
