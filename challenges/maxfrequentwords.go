package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"sort"
)

type Node struct {
	key   []byte
	value int
}

func maxFrequentWords(reader io.Reader, amount int) ([]byte, error) {
	res := make([]byte, 0)
	nodes := make([]*Node, 0)
	data, err := ioutil.ReadAll(reader)
	data = bytes.ToLower(data)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(data)
	word := make([]byte, 0)
	for {
		b, readErr := buff.ReadByte()
		if readErr != nil {
			break
		}
		if (b > 64 && b < 91) || (b > 96 && b < 123) {
			word = append(word, b)
		} else {
			word = bytes.ToLower(word)
			count := bytes.Count(data, word)
			n := &Node{key: word, value: count}
			isFound := false
			for k, v := range nodes {
				if v == nil {
					nodes[k] = n
					break
				}
				if bytes.Contains(v.key, n.key) {
					isFound = true
					break
				}
			}
			if !isFound {
				nodes = append(nodes, n)
			}
			word = make([]byte, 0)
		}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[j].value < nodes[i].value
	})
	counter := 0
	for _, v := range nodes {
		if counter == amount {
			break
		}
		res = append(res, v.key...)
		res = append(res, 32)
		counter++
	}
	return res, nil
}
