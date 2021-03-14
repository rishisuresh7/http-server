package wrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	w := NewWrapper()
	res, err := w.Atoi("9000")

	assert.Equal(t, 9000, res)
	assert.Equal(t, nil, err)
}
