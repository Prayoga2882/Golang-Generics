package Golang_Generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	first  T
	second T
}

func (d Data[_]) SayHello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func (d Data[T]) ChangeFirst(first T) T {
	d.first = first
	return d.first
}

func TestData(t *testing.T) {
	data := Data[int]{
		first:  1,
		second: 2,
	}
	fmt.Println(data)
}

func TestDataMethod(t *testing.T) {
	data := Data[string]{
		first:  "first",
		second: "second",
	}
	assert.Equal(t, "Hello, world", data.SayHello("world"))
}
