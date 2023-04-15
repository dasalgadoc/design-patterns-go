package main

func GetFibonacci(n int) (interface{}, error) {
	return HeavyFibonacci(n), nil
}

func HeavyFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return HeavyFibonacci(n-1) + HeavyFibonacci(n-2)
}
