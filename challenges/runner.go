package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("text.txt")
	words := maxFrequentWordsTwo(f, 10)
	fmt.Println(words)
}
