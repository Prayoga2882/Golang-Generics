package Golang_Generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Path: Employee interface
type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

// path: Manager interface
type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

// path: VicePresident interface
type VicePresident interface {
	GetName() string
	GetVicePresident() string
	GetAge() int
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresident() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "John", GetName(&MyManager{Name: "John"}))
	assert.Equal(t, "John", GetName(&MyVicePresident{Name: "John"}))
}
