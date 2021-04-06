package main

import "fmt"

func main() {
	arr := []int{1, 5, 11, 47, 8, 15, 2}
	fmt.Println(quickSort(arr, 0, len(arr)))
}
