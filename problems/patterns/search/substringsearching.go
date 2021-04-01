package main

func searchSub(str string, sub string) int {
	if len(sub) > len(str) {
		return -1
	}
	mainArr := []rune(str)
	subArr := []rune(sub)
	matchCounter := 0

	for i := 0; i < len(mainArr); i++ {
		for j := 0; j < len(subArr); j++ {
			if mainArr[i+j] != subArr[j] {
				break
			}
			if j == len(subArr)-1 {
				matchCounter++
			}
		}
	}
	return matchCounter
}
