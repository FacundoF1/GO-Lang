package examples

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicExample(t *testing.T) {
	result := Add(1, 2)
	assert.Equal(t, 3, result)
}

type mock struct {
	called bool
}

func (m *mock) Write(p []byte) (n int, err error) {
	fmt.Println(p)
	m.called = true
	return 0, nil
}

func TestWithDependency(t *testing.T) {
	mockedDependency := mock{}
	DependencyUser(&mockedDependency)
}
