package main

func recursionRange(num int) int {
	if num == 0 {
		return 0
	}
	return num + recursionRange(num-1)
}
