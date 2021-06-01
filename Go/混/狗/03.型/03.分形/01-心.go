/* Mandelbrot心之分形集
* 参考：
* 1. 曼德博集合: https://zh.wikipedia.org/zh-cn/曼德博集合
 */
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		x丅, y丅, x丄, y丄 = -2, -2, 2, 2
		图宽, 图高         = 1024, 1024
	)

	心图 := image.NewRGBA(image.Rect(0, 0, 图宽, 图高))
	for by := 0; by < 图高; by++ {
		y := float64(by)/图高*(y丄-y丅) + y丅
		for bx := 0; bx < 图宽; bx++ {
			x := float64(bx)/图宽*(x丄-x丅) + x丅
			z := complex(x, y)
			心图.Set(bx, by, 心色(z))
		}
	}
	png.Encode(os.Stdout, 心图)
}

func 心色(z complex128) color.Color {
	const (
		迭代次数 = 200
		色梯   = 15
	)
	var v complex128
	for n := uint8(0); n < 迭代次数; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {

			//return color.Gray{255 - 色梯*n}
			return color.RGBA{0, 0, 255 - 色梯*n, 255}
		}
	}
	return color.RGBA{0, 0, 255, 255}
}
