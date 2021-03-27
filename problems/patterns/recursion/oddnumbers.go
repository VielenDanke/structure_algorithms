package main

func collectOddNumbers(arr []int) []int {
	var result []int

	if len(arr) == 0 {
		return result
	}
	if arr[0]%2 != 0 {
		result = append(result, arr[0])
	}
	return append(result, collectOddNumbers(arr[1:])...)
}
