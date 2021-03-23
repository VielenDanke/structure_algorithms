package main

func minSubArrayLen(arr []int, num int) int {
	minLen := len(arr)
	start := 0
	end := 0
	total := 0

	for start < len(arr) {
		if total < num && end < len(arr) {
			total += arr[end]
			end++
		} else if total >= num {
			tempLen := end - start
			if minLen > tempLen {
				minLen = tempLen
			}
			total -= arr[start]
			start++
		} else {
			break
		}
	}
	return minLen
}
