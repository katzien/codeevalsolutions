package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibZero(t *testing.T) {

	actual := Fibonacci(0)

	assert.Equal(t, actual, int64(0))
}

func TestFibOne(t *testing.T) {

	actual := Fibonacci(1)

	assert.Equal(t, actual, int64(1))
}

func TestFib(t *testing.T) {

	cases := map[int64]int64{
		5:  5,
		12: 144,
		35: 9227465,
		// 50: 12586269025, // slow - takes 100.695s
		// 74: 1304969544928657, // slow - doesn't finish anytime soon
	}

	for n, expected := range cases {
		actual := Fibonacci(n)
		assert.Equal(t, actual, expected)
	}
}
