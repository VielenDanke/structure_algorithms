package main

func countUnique(arr []int) int {
	temp := arr[0]
	counter := 1
	for _, v := range arr {
		if temp == v {
			continue
		} else {
			temp = v
			counter++
		}
	}
	return counter
}

func countUniqueSecond(arr []int) int {
	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
	}
	return i + 1
}

func countUniqueThird(arr []int) int {
	m := make(map[int]interface{})
	for _, v := range arr {
		m[v] = nil
	}
	return len(m)
}
