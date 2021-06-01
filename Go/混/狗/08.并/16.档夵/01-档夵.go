/* 计算目录所占磁盘空间
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var 进度 = flag.Bool("进度", false, "显示求目录所占磁盘空间进度")

func main() {
	// 从命令行获取所求目录
	flag.Parse()
	始录 := flag.Args()
	if len(始录) == 0 {
		始录 = []string{"."}
	}

	档夵 := make(chan int64)
	var 等 sync.WaitGroup

	for _, 始 := range 始录 {
		等.Add(1)
		go 历录(始, &等, 档夵)
	}

	go func() {
		等.Wait()
		close(档夵)
	}()

	var 嘀嗒 <-chan time.Time
	if *进度 {
		嘀嗒 = time.Tick(500 * time.Millisecond)
	}

	var 档数, 储夵 int64
环:
	for {
		select {
		case 夵, 成 := <-档夵:
			if !成 {
				break 环
			}
			档数++
			储夵 += 夵
		case <-嘀嗒:
			磁盘用量(档数, 储夵)
		}
	}
	磁盘用量(档数, 储夵)
}

func 磁盘用量(档数, 储夵 int64) {
	fmt.Printf("共%d个文件，占用空间%.2gGB，或%.2gMB，或%.2gKB，或%dB\n",
		档数, float64(储夵)/(1<<30), float64(储夵)/(1<<20), float64(储夵)/(1<<10), 储夵)
}

var 旗号 = make(chan struct{}, 20)

func 录项集(录 string) []os.FileInfo {
	旗号 <- struct{}{}
	defer func() {
		<-旗号
	}()

	项集, 障 := ioutil.ReadDir(录)
	if 障 != nil {
		fmt.Fprintf(os.Stderr, "01-档夵：%v\n", 障)
		return nil
	}
	return 项集
}

func 历录(录 string, 等 *sync.WaitGroup, 档夵 chan<- int64) {
	defer 等.Done()
	for _, 项 := range 录项集(录) {
		if 项.IsDir() {
			等.Add(1)
			子录 := filepath.Join(录, 项.Name())
			go 历录(子录, 等, 档夵)
		} else {
			档夵 <- 项.Size()
		}
	}
}
