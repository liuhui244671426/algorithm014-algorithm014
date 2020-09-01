package Week_04

var m map[int]int

func Fib(N int) int {
	return fib(N)
}

func fib(N int) int {
	if N <= 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}
