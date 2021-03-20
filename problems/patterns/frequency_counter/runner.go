package main

import "fmt"

/*
make O(N^2) to O(N) by reducing one inner cycle with replacing it to map or struct
*/
func main() {
	fmt.Println(anagram("cabcdcce", "edbacccc"))
}
