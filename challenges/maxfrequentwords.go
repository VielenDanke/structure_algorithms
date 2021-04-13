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

func maxFrequentWords(reader io.Reader, amount int) ([]*Node, error) {
	nodes := make([]*Node, 0)
	data, readAllErr := ioutil.ReadAll(reader)
	if readAllErr != nil {
		return nil, readAllErr
	}
	filteredBuff := bytes.NewBuffer(make([]byte, 0))
	dataBuff := bytes.NewBuffer(data)
	isRepeated := false
	temp := make([]byte, 1)
	word := make([]byte, 0)
	for {
		n, err := dataBuff.Read(temp)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		temp = bytes.ToLower(temp[0:n])
		if temp[0] > 96 && temp[0] < 123 {
			word = append(word, temp[0])
			isRepeated = false
		} else {
			if !isRepeated {
				if !bytes.Contains(filteredBuff.Bytes(), word) {
					filteredBuff.Write(word)
					nodes = append(nodes, &Node{key: word, value: bytes.Count(data, word)})
				}
				filteredBuff.WriteByte(32)
				word = make([]byte, 0)
				isRepeated = true
			}
		}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].value > nodes[j].value
	})
	return nodes[0:amount], nil
}
