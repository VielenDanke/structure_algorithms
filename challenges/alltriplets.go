package main

import (
	"sort"
)

func threeSum(nums []int) (res [][]int) {
	var l, r, z int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		l = i + 1
		r = len(nums) - 1
		z = nums[i]
		for l < r {
			isRepeated := false
			xNums := nums[l]
			yNums := nums[r]
			if z+xNums+yNums == 0 {
				for _, v := range res {
					if v[0] == z && v[1] == xNums && v[2] == yNums {
						isRepeated = true
					}
				}
				if !isRepeated {
					res = append(res, []int{z, xNums, yNums})
				}
				l++
				r--
			} else if z+yNums+xNums < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return
}
