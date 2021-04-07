package main

import (
	"math"
)

func radixSort(arr []int) []int {
	most := mostDigist(arr)
	for k := 0; k < most; k++ {
		digitBuckets := make([][]int, 10)
		for i := 0; i < len(arr); i++ {
			digit := getDigit(arr[i], k)
			digitBuckets[digit] = append(digitBuckets[digit], arr[i])
		}
		var temp []int
		for _, v := range digitBuckets {
			temp = append(temp, v...)
		}
		arr = temp
	}
	return arr
}

func getDigit(digit, idx int) int {
	return int(math.Floor(math.Abs(float64(digit))/math.Pow(10, float64(idx)))) % 10
}

func digitCount(num int) int {
	if num == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(math.Abs(float64(num))))) + 1
}

func mostDigist(arr []int) int {
	most := 0
	for _, v := range arr {
		most = int(math.Max(float64(most), float64(digitCount(v))))
	}
	return most
}
