package main

import (
	"fmt"
	"os"
)

func main() {
	open, _ := os.Open("text.txt")
	words, _ := maxFrequentWords(open, 1)
	for _, v := range words {
		fmt.Printf("Word: %s, amount: %d\n", v.key, v.value)
	}
}
