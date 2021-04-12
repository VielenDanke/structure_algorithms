package main

func longestPalindrome(s string) string {
	length := len(s)
	if length < 1 || length > 1000 {
		return ""
	}
	if length == 1 {
		return s
	}
	start := 0
	end := 0
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] == arr[j] {
				temp := arr[i : j+1]
				isPalindrome := checkIfPalindrome(temp, len(temp))
				if end-start < j+1-i && isPalindrome {
					start = i
					end = j + 1
				}
			}
		}
	}
	return string(arr[start:end])
}

func checkIfPalindrome(arr []rune, length int) bool {
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}
