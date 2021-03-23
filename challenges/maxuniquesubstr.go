package main

func findLongestSubstring(str string) int {
	longest := 0
	m := make(map[rune]int)
	for k, v := range str {
		_, ok := m[v]
		if !ok {
			m[v] = k
		} else {
			longest = k
		}
	}
	return longest
}
