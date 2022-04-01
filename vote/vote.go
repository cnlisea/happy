package vote

type Vote struct {
	m []interface{}
}

func New(num int) *Vote {
	return &Vote{
		m: make([]interface{}, 0, num),
	}
}

func (v *Vote) Add(key interface{}) {
	if v.Exist(key) {
		return
	}
	v.m = append(v.m, key)
}

func (v *Vote) Num() int {
	return len(v.m)
}

func (v *Vote) Exist(key interface{}) bool {
	var (
		exist bool
	)
	v.Range(func(k interface{}) bool {
		if k == key {
			exist = true
		}
		return !exist
	})
	return exist
}

func (v *Vote) Range(f func(key interface{}) bool) {
	for i := range v.m {
		if !f(v.m[i]) {
			break
		}
	}
}

func (v *Vote) Reset() {
	v.m = v.m[:0]
}
