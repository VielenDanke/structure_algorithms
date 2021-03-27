package main

func recursionPalidrome(str string) bool {
	temp := []rune(str)
	if len(str) == 1 {
		return true
	}
	if len(str) == 2 {
		return temp[0] == temp[1]
	}
	if temp[0] == temp[len(temp)-1] {
		return recursionPalidrome(string(temp[1 : len(temp)-1]))
	}
	return false
}
