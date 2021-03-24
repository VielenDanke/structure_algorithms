package main

import "fmt"

func simple(num int) {
	if num <= 0 {
		fmt.Println("Done!")
		return
	}
	fmt.Println(num)
	num--
	simple(num)
}
