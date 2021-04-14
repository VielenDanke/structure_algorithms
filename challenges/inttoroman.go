package main

import (
	"bytes"
)

func intToRoman(num int) string {
	strBuff := bytes.NewBufferString("")
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(nums); i++ {
		for num >= nums[i] {
			num -= nums[i]
			strBuff.WriteString(romans[i])
		}
	}
	return strBuff.String()
}

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		n := m[string(s[i])]
		if i+1 < len(s) {
			if n >= m[string(s[i+1])] {
				sum += n
			} else {
				sum -= n
			}
		} else {
			sum += n
		}
	}
	return sum
}
