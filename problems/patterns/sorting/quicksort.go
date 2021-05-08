package main

func quickSort(arr []int, left, right int) []int {
	if left < right {
		pivotIdx := pivot(arr, left, right)
		quickSort(arr, left, pivotIdx-1)
		quickSort(arr, pivotIdx+1, right)
	}
	return arr
}

// 67, 50, 32, 45, 929, 3214, 4, 323
func pivot(arr []int, start, end int) int {
	p := arr[start]
	swapIdx := start
	for i := start + 1; i < end; i++ {
		if p > arr[i] {
			swapIdx++
			swap(arr, swapIdx, i)
		}
	}
	swap(arr, start, swapIdx)
	return swapIdx
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
