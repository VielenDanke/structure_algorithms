package main

func isSame(first []int, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		correctIndex := findIndex(first[i], second)
		if correctIndex == -1 {
			return false
		}
	}
	return true
}

func findIndex(num int, arr []int) int {
	for k, v := range arr {
		if v == num*num {
			return k
		}
	}
	return -1
}

func isSameBetter(first []int, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	holder := make(map[int]interface{})
	for _, v := range second {
		holder[v] = nil
	}
	for _, v := range first {
		_, ok := holder[v*v]
		if !ok {
			return false
		}
	}
	return true
}
