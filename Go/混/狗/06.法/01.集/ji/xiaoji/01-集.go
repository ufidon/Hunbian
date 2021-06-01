/* 位图实现的小非负整数集
 */
package xiaoji

import (
	"bytes"
	"fmt"
)

type Z非负整集 struct { // 有大写字母前缀，出包界
	字列 []uint64 // 无大写字母前缀，未出包界
}

// 数在集中否？
func (集 *Z非负整集) Z有(数 uint) bool {
	号, 位 := 数/64, 数%64
	return int(号) < len(集.字列) && 集.字列[号]&(1<<位) != 0
}

// 添数入集
func (集 *Z非负整集) Z添(数 uint) {
	号, 位 := 数/64, 数%64
	for int(号) >= len(集.字列) {
		集.字列 = append(集.字列, 0)
	}
	集.字列[号] |= 1 << uint64(位)
}

// 取并集
func (甲 *Z非负整集) Z并(乙 *Z非负整集) {
	for 号, 字 := range 乙.字列 {
		if 号 < len(甲.字列) {
			甲.字列[号] |= 字
		} else {
			甲.字列 = append(甲.字列, 字)
		}
	}
}

// 显集合
func (集 *Z非负整集) String() string {
	var 串 bytes.Buffer
	串.WriteByte('{')
	for 号, 字 := range 集.字列 {
		if 字 == 0 {
			continue
		}
		for 位 := 0; 位 < 64; 位++ {
			if 字&(1<<uint64(位)) != 0 {
				if 串.Len() > len("{") {
					串.WriteByte(',')
				}
				fmt.Fprintf(&串, "%d", 64*号+位)
			}
		}
	}
	串.WriteByte('}')
	return 串.String()
}
