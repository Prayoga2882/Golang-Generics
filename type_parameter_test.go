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
	result := length[string]("Hello")
	fmt.Println(result)
}

func MultipleTypeParams(T a, U a) {
	fmt.Println(T, U)
}

func TestMultipleTypeParams(t *testing.T) {
	MultipleTypeParams("Hello", 1)
	MultipleTypeParams(1, "Hello")
}
