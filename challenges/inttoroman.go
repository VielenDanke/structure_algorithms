package main

import "fmt"

func intToRoman(num int) string {
	incr := 1
	for num != 0 {
		temp := num % 10
		num /= 10
		fmt.Println(temp)
		incr *= 10
	}
	return ""
}
