package wrapper

import "strconv"

type Wrapper interface {
	Atoi(string) (int, error)
}

type wrapper struct{}

func NewWrapper() Wrapper {
	return &wrapper{}
}

func (w *wrapper) Atoi(s string) (int, error) {
	return strconv.Atoi(s)
}
