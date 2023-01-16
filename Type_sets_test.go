package Golang_Generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Number interface {
	int | float64 | float32 | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 200))
	assert.Equal(t, 100.0, Min[float64](100.0, 200.0))
	assert.Equal(t, 100.0, Min[float32](100.0, 200.0))
}
