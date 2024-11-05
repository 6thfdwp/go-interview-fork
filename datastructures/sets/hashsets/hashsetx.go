package hashsets

type HashSetStr struct {
	data map[string]struct{}
}

func NewHashSetStr() *HashSetStr {
	data := make(map[string]struct{})
	return &HashSetStr{
		data: data,
	}
}

func (t *HashSetStr) Size() int {
	return len(t.data)
}

func (t *HashSetStr) Add(v string) {
	t.data[v] = struct{}{}
}

func (t *HashSetStr) Contains(v string) bool {
	// if _, ok := t.data[v]; ok {
	// 	return true
	// }
	_, ok := t.data[v]
	return ok
}
