package evaluator

import (
	"fmt"
	"monkey/object"
)

var 内函集 = map[string]*object.N内型{
	"节数": &object.N内型{
		H函: func(参集 ...object.H盒) object.H盒 {
			if len(参集) != 1 {
				return 造障("需一个参数，不是%d个",
					len(参集))
			}

			switch 参 := 参集[0].(type) {
			case *object.D队型:
				return &object.Z整型{Z值: int64(len(参.Y元集))}
			case *object.C串型:
				return &object.Z整型{Z值: int64(len(参.Z值))}
			default:
				return 造障("节数不支持%s",
					参集[0].X型())
			}
		},
	},
	"字数": &object.N内型{
		H函: func(参集 ...object.H盒) object.H盒 {
			if len(参集) != 1 {
				return 造障("需一个参数，不是%d个",
					len(参集))
			}

			switch 参 := 参集[0].(type) {
			case *object.D队型:
				return &object.Z整型{Z值: int64(len(参.Y元集))}
			case *object.C串型:
				return &object.Z整型{Z值: int64(len([]rune(参.Z值)))}
			default:
				return 造障("节数不支持%s",
					参集[0].X型())
			}
		},
	},
	"显示": &object.N内型{
		H函: func(参集 ...object.H盒) object.H盒 {
			for _, 参 := range 参集 {
				fmt.Println(参.C察值())
			}

			return K空值
		},
	},
	"首元": &object.N内型{
		H函: func(参集 ...object.H盒) object.H盒 {
			if len(参集) != 1 {
				return 造障("需一个参数，不是%d个",
					len(参集))
			}
			if 参集[0].X型() != object.D队盒 {
				return 造障("首元参数应为队而不是%s",
					参集[0].X型())
			}

			队 := 参集[0].(*object.D队型)
			if len(队.Y元集) > 0 {
				return 队.Y元集[0]
			}

			return K空值
		},
	},
	"尾元": &object.N内型{
		H函: func(参集 ...object.H盒) object.H盒 {
			if len(参集) != 1 {
				return 造障("需一个参数，不是%d个",
					len(参集))
			}
			if 参集[0].X型() != object.D队盒 {
				return 造障("尾元参数应为队而不是%s",
					参集[0].X型())
			}

			队 := 参集[0].(*object.D队型)
			长度 := len(队.Y元集)
			if 长度 > 0 {
				return 队.Y元集[长度-1]
			}

			return K空值
		},
	},
	"掐头": &object.N内型{
		H函: func(参集 ...object.H盒) object.H盒 {
			if len(参集) != 1 {
				return 造障("需一个参数，不是%d个",
					len(参集))
			}
			if 参集[0].X型() != object.D队盒 {
				return 造障("掐头参数应为队而不是%s",
					参集[0].X型())
			}

			队 := 参集[0].(*object.D队型)
			长度 := len(队.Y元集)
			if 长度 > 0 {
				新元集 := make([]object.H盒, 长度-1, 长度-1)
				copy(新元集, 队.Y元集[1:长度])
				return &object.D队型{Y元集: 新元集}
			}

			return K空值
		},
	},
	"附尾": &object.N内型{
		H函: func(参集 ...object.H盒) object.H盒 {
			if len(参集) != 2 {
				return 造障("需两个参数，不是%d个",
					len(参集))
			}
			if 参集[0].X型() != object.D队盒 {
				return 造障("附尾参数应为队而不是%s",
					参集[0].X型())
			}

			队 := 参集[0].(*object.D队型)
			长度 := len(队.Y元集)

			新元集 := make([]object.H盒, 长度+1, 长度+1)
			copy(新元集, 队.Y元集)
			新元集[长度] = 参集[1]

			return &object.D队型{Y元集: 新元集}
		},
	},
}
