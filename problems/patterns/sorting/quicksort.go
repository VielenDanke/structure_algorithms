package main

func quickSort(arr []int, left, right int) {
	if left < right {
		partitionIdx := partition(arr, left, right)
		quickSort(arr, left, partitionIdx - 1)
		quickSort(arr, partitionIdx + 1, right)
	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
