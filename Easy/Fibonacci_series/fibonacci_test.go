package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibZero(t *testing.Testing) {

	actual := Fibonacci(0)

	assert.Equal(t, actual, 0)
}

func TestFibOne(t *testing.Testing) {

	actual := Fibonacci(1)

	assert.Equal(t, actual, 1)
}

func TestFib(t *testing.Testing) {

	cases := map[int64]int64{
		5:   5,
		12:  144,
		145: 898923707008479989274290850145,
	}

	for n, expected := range cases {
		actual := Fibonacci(n)
		assert.Equal(t, actual, expected)
	}
}
