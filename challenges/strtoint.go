package main

func myAtoi(s string) int {
	if len(s) < 0 || len(s) > 200 {
		return 0
	}
	maxInt32 := 2147483647
	minInt32 := -2147483648
	temp := 1
	res := 0
	counter := 0
	isNegative := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] > 47 && s[i] < 58 {
			elem := int(s[i]) - 48
			elem *= temp
			res += elem
			temp *= 10
			counter++
			continue
		}
		if s[i] == 45 && counter != 0 {
			isNegative = !isNegative
			continue
		}
		counter = 0
	}
	if isNegative {
		res = -res * 2 + res
	}
	if minInt32 > res {
		return minInt32
	}
	if maxInt32 < res {
		return maxInt32
	}
	return res
}

func myAtoiCorrect(s string) int {
	if len(s) < 0 || len(s) > 200 {
		return 0
	}
	maxInt32 := 2147483647
	minInt32 := -2147483648
	res := 0
	isRun := false
	isNegative := false
	arr := make([]int, 0)
	for _, v := range s {
		if v == 32 {
			if isRun {
				break
			}
			continue
		}
		if v == 43 {
			if isRun {
				break
			}
			isRun = true
			isNegative = false
			continue
		}
		if v == 45 {
			if isRun {
				break
			}
			isRun = true
			isNegative = true
			continue
		}
		if v > 47 && v < 58 {
			elem := int(v) - 48
			if len(arr) == 0 && v == 48 {
				isRun = true
				continue
			}
			arr = append(arr, elem)
			continue
		} else {
			break
		}
	}
	if len(arr) == 0 {
		return 0
	}
	incr := 1
	for i := 1; i < len(arr); i++ {
		incr *= 10
	}
	for _, v := range arr {
		res += v * incr
		incr /= 10
	}
	if isNegative {
		res = -res * 2 + res
	}
	if !isNegative && res < 0 {
		return maxInt32
	}
	if res < minInt32 {
		return minInt32
	}
	if res > maxInt32 {
		return maxInt32
	}
	return res
}