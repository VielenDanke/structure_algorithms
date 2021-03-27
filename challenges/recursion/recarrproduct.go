package main

func arrayProduct(arr []int) int {
	if len(arr) == 0 {
		return 1
	}
	num := arr[0]
	return num * arrayProduct(arr[1:])
}
