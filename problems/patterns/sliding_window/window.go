package main

func maxNumbersSum(num int, arr []int) int {
	tempSum := 0
	maxSum := 0
	if len(arr) == 0 || num > len(arr) {
		return -1
	}
	for i := 0; i < num; i++ {
		maxSum += arr[i]
	}
	tempSum = maxSum
	for i := num; i < len(arr); i++ {
		tempSum = tempSum - arr[i-num] + arr[i]
		if tempSum > maxSum {
			maxSum = tempSum
		}
	}
	return maxSum
}
