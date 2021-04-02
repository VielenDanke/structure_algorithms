package main

import "sort"

type IntArr []int

func (ir IntArr) Len() int {
	return len(ir)
}

func (ir IntArr) Swap(i, j int) {
	ir[i], ir[j] = ir[j], ir[i]
}

func (ir IntArr) Less(i, j int) bool {
	return ir[i] < ir[j]
}

func sortBuiltIn(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

func sortBuiltInSecond(arr []int) []int {
	sort.Sort(IntArr(arr))
	return arr
}
