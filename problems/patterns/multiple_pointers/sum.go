package main

// O(N^2)
func sumPair(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]+arr[i] == 0 {
				return []int{arr[j], arr[i]}
			}
		}
	}
	return []int{}
}

// O(N)
func sumPairBetter(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == 0 {
			return []int{arr[left], arr[right]}
		}
		if sum > 0 {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
