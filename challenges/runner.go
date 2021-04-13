package main

import (
	"fmt"
	"os"
)

func main() {
	open, _ := os.Open("text.txt")
	words, _ := maxFrequentWords(open, 1)
	fmt.Printf("%s", words)
}
