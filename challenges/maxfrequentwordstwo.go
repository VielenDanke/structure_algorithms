package main

import (
	"bufio"
	"io"
	"log"
	"sort"
	"unicode"
)

type node struct {
	word   string
	amount int
}

func maxFrequentWordsTwo(reader io.Reader, amount int) []string {
	bReader := bufio.NewReader(reader)
	countWords := make(map[string]int)
	word := make([]rune, 0)
	inWord := false
	for {
		r, err := readByte(bReader)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Could not read from the wile")
		}
		word = append(word, r)
		if unicode.IsSpace(r) && inWord {
			v, ok := countWords[string(word)]
			if ok {
				countWords[string(word)] = v + 1
			} else {
				countWords[string(word)] = 1
			}
			inWord = false
			word = make([]rune, 0)
		}
		inWord = unicode.IsLetter(r)
	}
	nodes := make([]*node, 0)
	for k, v := range countWords {
		nodes = append(nodes, &node{word: k, amount: v})
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].amount > nodes[j].amount
	})
	nodes = nodes[:amount]
	words := make([]string, 0)
	for _, v := range nodes {
		words = append(words, v.word)
	}
	return words
}

func readByte(reader io.Reader) (rune, error) {
	var buf [1]byte
	_, err := reader.Read(buf[:])
	return rune(buf[0]), err
}
