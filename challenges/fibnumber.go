package main

func calculateFibNumber(n int) int {
	x := 1
	y := 1
	for i := 0; i < n; i++ {
		x, y = x+y, x
	}
	return x
}
