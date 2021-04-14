package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := readAllBytes(file)
	if err != nil {
		log.Fatal(err)
	}
	words, err := maxFrequentWords(data, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(words)
}
