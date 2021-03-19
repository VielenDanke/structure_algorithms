package main

func anagram(first string, second string) bool {
	fHolder := make(map[rune]int)
	sHolder := make(map[rune]int)
	for _, v := range first {
		r, ok := fHolder[v]
		if ok {
			fHolder[v] = r + 1
		} else {
			fHolder[v] = 1
		}
	}
	for _, v := range second {
		r, ok := sHolder[v]
		if ok {
			sHolder[v] = r + 1
		} else {
			sHolder[v] = 1
		}
	}
	for k, v := range fHolder {
		sk, ok := sHolder[k]
		if !ok || !(sk == v) {
			return false
		}
	}
	return true
}
