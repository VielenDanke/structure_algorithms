package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{1, 3, 5, 55, 61, 32, 41, 33, 21, 14, 15, 16, 47, 1, 2, 3, 55, 4}
	fs := time.Now().Nanosecond()
	fmt.Println(sumWithoutOne(arr))
	fmt.Println("First approach time ", time.Now().Nanosecond()-fs)
	ss := time.Now().Nanosecond()
	fmt.Println(sumWithoutOneBetter(arr))
	fmt.Println("Second approach time ", time.Now().Nanosecond()-ss)
}
