package main

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1
	for left < right {
		min := 0
		if height[left] > height[right] {
			min = height[right]
		} else {
			min = height[left]
		}
		area := min * (right - left)
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
