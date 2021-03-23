package main

func findLongestSubstring(str string) int {
	m := make(map[rune]int)
	idx := 0
	start := 0
	longest := 0

	r := []rune(str)

	for idx < len(str) {
		ch := r[idx]
		_, ok := m[ch]
		if !ok {
			m[ch] = idx
			idx++
		} else {
			temp := idx - start
			if longest < temp {
				longest = temp
			}
			delete(m, r[start])
			start++
		}
	}
	return longest
}
