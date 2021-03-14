package wrapper

import "github.com/stretchr/testify/mock"

type Wrapper struct {
	mock.Mock
}

func (w *Wrapper) Atoi(s string) (int, error) {
	ret := w.Called(s)

	return ret.Int(0), ret.Error(1)
}
