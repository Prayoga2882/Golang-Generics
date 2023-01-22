package Golang_Generics_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FindMin[T interface {
	int | int64 | float32 | float64
}](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 1, FindMin[int](100, 200))
	assert.Equal(t, 1, FindMin[int64](int64(100), int64(200)))
	assert.Equal(t, 1, FindMin[int](100.0, 200.0))
}

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{"John", "Doe"}
	first := GetFirst[[]string, string](names)
	assert.Equal(t, "John", first)
}
