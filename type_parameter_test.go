package Golang_Generics_test

import (
	"fmt"
	"testing"
)

type a = interface{}

func length[T a](param T) T {
	fmt.Println(param)
	return param
}

func TestTypeParameter(t *testing.T) {
	length("hello")
	length(1)
}
