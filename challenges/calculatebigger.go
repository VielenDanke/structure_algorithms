package main

import "sort"

func calculateAllBigger(arr []int) []int {
	cp := append([]int{}, arr...)
	cpm := make(map[int]int)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	for k, v := range arr {
		cpm[v] = len(arr) - 1 - k
	}
	for k, v := range cp {
		cp[k] = cpm[v]
	}
	return cp
}
