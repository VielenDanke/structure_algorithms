package main

// time O(n log n)
// space O(n)
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := mergeSort(arr[:len(arr)/2])
	right := mergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(first, second []int) []int {
	var results []int
	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			results = append(results, first[i])
			i++
		} else {
			results = append(results, second[j])
			j++
		}
	}
	for i < len(first) {
		results = append(results, first[i])
		i++
	}
	for j < len(second) {
		results = append(results, second[j])
		j++
	}
	return results
}
