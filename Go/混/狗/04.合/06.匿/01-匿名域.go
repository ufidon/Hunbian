/* 结构之匿名域或嵌入域
 */

package main

import (
	"fmt"
	"image/color"
)

// D, 点
type D平面点 struct {
	Dx, Dy float64
}

type J矩形 struct {
	D左上点, D右下点 D平面点 // 同类型域不可匿名，否则无法区分
}

type J矩块 struct {
	//J框 J矩形
	J矩形                  // 匿名域（其实以类型名为名）或嵌入域
	S框色, S块色 color.Color // 同类型域不可匿名，否则无法区分
}

func main() {
	var 矩块1, 矩块2 J矩块

	// 复合结构逐域赋值
	矩块1.S框色 = color.RGBA{0xff, 0, 0, 0xff}
	矩块1.S块色 = color.RGBA{0, 0xff, 0, 0xff}

	// 匿名域附值，可省掉中间匿名域
	矩块1.D左上点 = D平面点{5, 6}     // 等价于 矩块1.J矩形.D左上点 = D平面点{5, 6}
	矩块1.J矩形.D右下点 = D平面点{7, 8} // 匿名域实则以构型名为域名

	// 不论匿名与非，始化者必须以结构赋值，匿名域不可略
	矩块2 = J矩块{ // 无名赋值，必须遵循各域定义顺序
		J矩形{
			D平面点{Dy: 2, Dx: 1}, // 名值对赋值，顺序无关
			D平面点{Dy: 4},        // 名值对赋值与无名赋值不可混用， 名值对可部分赋值，未赋值者取其类零值
		},
		color.White,
		color.Black,
	}
	// 下行错： 始化者赋值，匿名域不可略
	// 矩块2 = J矩块{{D平面点{1, 2}, D平面点{3, 4}}, color.White, color.Black}

	fmt.Printf("矩块1：%#v\n", 矩块1) // %#v显示结构域名与值
	fmt.Printf("矩块2：%#v\n", 矩块2)
}
