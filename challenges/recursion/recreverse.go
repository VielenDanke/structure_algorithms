package main

import "fmt"

func reverseRecursion(str string) string {
	if len(str) <= 1 {
		return str
	}
	temp := []rune(str)
	fmt.Println(str)
	return reverseRecursion(string(temp[1:])) + string(temp[0])
}
