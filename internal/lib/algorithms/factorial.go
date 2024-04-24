package algorithms

func FactorialRecursive(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * FactorialRecursive(n-1)
	}
}
