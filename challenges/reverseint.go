package main

func reverse(x int) int {
	maxInt32 := 2147483647
	minInt32 := -2147483648
	res := 0
	for x != 0 {
		temp := x % 10
		x /= 10
		if res > maxInt32/10 || (res == maxInt32/10 && temp > 7) {
			return 0
		}
		if res < minInt32/10 || (res == minInt32/10 && temp < -8) {
			return 0
		}
		res = res * 10 + temp
	}
	return res
}