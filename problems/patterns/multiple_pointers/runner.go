package main

import "fmt"

/*
take two values in the memory instead of cycle with inner cycle
*/
func main() {
	arr := []int{-2, -2, -1, 0, 1, 15, 15, 16, 22, 23, 27, 27, 32, 32, 32, 33}
	fmt.Println(countUnique(arr))
	fmt.Println(countUniqueSecond(arr))
	fmt.Println(countUniqueThird(arr))
}
