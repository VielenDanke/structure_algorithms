package main

func maxSubarray(arr []int, sublen int) int {
	// if sublen > len(arr)
	if sublen > len(arr) {
		return -1
	}
	// if arr == nil
	if arr == nil {
		return -1
	}
	// if sublen == 0
	if sublen == 0 {
		return -1
	}
	// calculate sum in len == sublen
	maxSum := 0
	tempSum := 0

	for i := 0; i < sublen; i++ {
		maxSum += arr[i]
	}
	tempSum = maxSum

	for i := sublen; i < len(arr); i++ {
		tempSum = tempSum - arr[i-sublen] + arr[i]
		if tempSum > maxSum {
			maxSum = tempSum
		}
	}
	// return
	return maxSum
}
