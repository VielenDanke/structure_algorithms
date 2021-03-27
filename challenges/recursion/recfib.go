package main

func recursionFibNumber(n int) int {
	if n <= 2 {
		return 1
	}
	return recursionFibNumber(n-1) + recursionFibNumber(n-2)
}
