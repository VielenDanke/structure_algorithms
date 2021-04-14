package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arr := make([]int, 0)
	for x != 0 {
		i := x % 10
		x /= 10
		arr = append(arr, i)
	}
	left := 0
	right := len(arr) - 1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindromeSecond(x int) bool {
	if x < 0 {
		return false
	}
	arr := strconv.Itoa(x)
	left := 0
	right := len(arr) - 1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}