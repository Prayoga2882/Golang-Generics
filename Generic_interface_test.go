package Golang_Generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](gs GetterSetter[T], value T) T {
	gs.SetValue(value)
	return gs.GetValue()
}

type MyData[T any] struct {
	value T
}

func (d *MyData[T]) GetValue() T {
	return d.value
}

func (d *MyData[T]) SetValue(value T) {
	d.value = value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}
	result := ChangeValue[string](&myData, "Hello")

	assert.Equal(t, "Hello", result)
}
