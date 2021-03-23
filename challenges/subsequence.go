package main

func isSubsequence(sub string, str string) bool {
	i := 0
	j := 0
	if sub == "" {
		return false
	}
	// sing sting
	// 0123 01234
	for j < len(str) {
		if str[j] == sub[i] {
			i++
		}
		if i == len(sub) {
			return true
		}
		j++
	}
	return false
}
