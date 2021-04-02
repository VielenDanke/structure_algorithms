package main

func bubbleSortArr(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			i = 0
		}
	}
	return arr
}

// better approach
func bubbleSortArrSecond(arr []int) []int {
	isSorted := false
	for !isSorted {
		isSorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
	}
	return arr
}
