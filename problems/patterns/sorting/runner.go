package main

import "fmt"

func main() {
	arr := []int{67, 11, 32, 45, 929, 311, 4, 15}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
