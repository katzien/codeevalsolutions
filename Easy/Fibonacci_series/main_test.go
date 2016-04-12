package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeInputReader struct {
	values []int64
}

func (fi FakeInputReader) ReadInput() []int64 {
	return fi.values
}

func TestReadInputFromFile(t *testing.T) {
	inputValues := []int64{5, 12, 35}

	fi := FakeInputReader{inputValues}
	result := fi.ReadInput()

	assert.Equal(t, result, inputValues)
}
