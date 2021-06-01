/* 伺者生成随机谐振GIF动画
* 参考: https://zh.wikipedia.org/wiki/利萨茹曲线
 */
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

var 色板 = []color.Color{
	color.White,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0x00},
}

const (
	// 色板索引号
	白 = 0
	红 = 1
	绿 = 2
	蓝 = 3
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法：", os.Args[0], "端号")
		os.Exit(1)
	}

	http.HandleFunc("/", func(回 http.ResponseWriter, 请 *http.Request) {
		if 障 := 请.ParseForm(); 障 != nil {
			log.Printf("请求表障：%q\n", 障)
			回.Header().Set("Content-type", "image/gif")
			谐振图(回, 1.5)
		} else {
			if 频比, 障 := strconv.ParseFloat(请.FormValue("频比"), 64); 障 != nil {
				log.Printf("参数障：%q\n", 障)
				回.Header().Set("Content-type", "image/gif")
				谐振图(回, 1.5)
			} else {
				回.Header().Set("Content-type", "image/gif")
				谐振图(回, 频比)
			}
		}

	})
	log.Fatal(http.ListenAndServe("localhost:"+os.Args[1], nil))
}

func 谐振图(出 io.Writer, 频比 float64) {
	const (
		圈数 = 8
		角率 = 0.002
		图夵 = 100 // [-100..100]
		帧数 = 64
		帧隔 = 8 // 单位10ms
	)

	// 频比为纵横震动频率比, 相差为纵横震动相位差
	//频比 := rand.Float64() * 3.0
	动画 := gif.GIF{LoopCount: 帧数}
	相差 := 0.0

	for 帧 := 0; 帧 < 帧数; 帧++ {
		框 := image.Rect(0, 0, 2*图夵+1, 2*图夵+1)
		图 := image.NewPaletted(框, 色板)
		for 时 := 0.0; 时 < 圈数*2*math.Pi; 时 += 角率 {
			横标 := math.Sin(时)
			纵标 := math.Sin(时*频比 + 相差)
			图.SetColorIndex(图夵+int(横标*图夵+0.5), 图夵+int(纵标*图夵+0.5), (uint8(帧 % 4)))
		}
		相差 += 0.1
		动画.Delay = append(动画.Delay, 帧隔)
		动画.Image = append(动画.Image, 图)
	}
	gif.EncodeAll(出, &动画)
}
