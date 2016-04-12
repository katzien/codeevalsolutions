package fibonacci

// Fibonacci is a function which calculates the Fibonacci value for the given number
func Fibonacci(n int64) int64 {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return (Fibonacci(n-1) + Fibonacci(n-2))
}
