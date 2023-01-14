package Golang_Generics_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func IsSameType[T comparable](a, b T) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSameType[string]("hai", "hai"))
}
