package main

func divideConquer(arr []int, toFind int) int {
	min := 0
	max := len(arr) - 1
	for min <= max {
		middle := (min + max) / 2
		current := arr[middle]
		if current < toFind {
			min = middle + 1
		} else if current > toFind {
			max = middle - 1
		} else {
			return middle
		}
	}
	return -1
}
