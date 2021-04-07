package main

import "fmt"

func main() {
	arr := []int{67, 50, 32, 45, 929, 3214, 4, 323}
	fmt.Println(quickSort(arr, 0, len(arr)))
}
