package main

/*
2 * (2 * 1)
*/
func pow(num int, incr int) int {
	if incr == 0 {
		return 1
	}
	return num * pow(num, incr-1)
}
