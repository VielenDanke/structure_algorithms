package main

func binarySearch(arr []int, search int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := (left + right) / 2
		current := arr[middle]
		if current < search {
			left = middle + 1
		} else if current > search {
			right = middle - 1
		} else {
			return middle
		}
	}
	return -1
}

func binarySearchSecond(arr []int, search int) int {
	start := 0
	end := len(arr) - 1
	middle := (start + end) / 2
	for arr[middle] != search && start <= end {
		if arr[middle] > search {
			end = middle - 1
		} else {
			start = middle + 1
		}
		middle = (start + end) / 2
	}
	if arr[middle] == search {
		return middle
	}
	return -1
}
