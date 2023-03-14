package demo

type T struct {
	Bar int
}

func New() *T {
	return &T{
		Bar: 42,
	}
}

func (t *T) Get() int {
	return t.Bar
}

func (t *T) Set(v int) {
	t.Bar = v
}
