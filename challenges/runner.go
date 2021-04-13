package main

import (
	"fmt"
	"os"
)

func main() {
	open, _ := os.Open("text.txt")
	words, _ := maxFrequentWords(open, 15)
	for _, v := range words {
		fmt.Printf("%s\n", v.key)
	}
}
