package object

func Z造闭境(W外境 *S算境) *S算境 {
	算境 := Z造境()
	算境.W外境 = W外境
	return 算境
}

func Z造境() *S算境 {
	仓 := make(map[string]H盒)
	return &S算境{C仓: 仓, W外境: nil}
}

type S算境 struct {
	C仓  map[string]H盒
	W外境 *S算境
}

func (境 *S算境) Q取(名 string) (H盒, bool) {
	盒, 有 := 境.C仓[名]
	if !有 && 境.W外境 != nil {
		盒, 有 = 境.W外境.Q取(名)
	}
	return 盒, 有
}

func (境 *S算境) S设(名 string, 值 H盒) H盒 {
	境.C仓[名] = 值
	return 值
}
