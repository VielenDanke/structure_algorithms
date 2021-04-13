package main

import (
	"bytes"
	"io"
	"io/ioutil"
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
					_, wErr := filteredBuff.Write(word)
					if wErr != nil {
						return nil, wErr
					}
					nodes = append(nodes, &Node{key: word, value: bytes.Count(data, word)})
				}
				if wbErr := filteredBuff.WriteByte(32); wbErr != nil {
					return nil, wbErr
				}
				word = make([]byte, 0)
				isRepeated = true
			}
		}
	}
	nodes = mergeSort(nodes)
	return nodes[0:amount], nil
}

func mergeSort(arr []*Node) []*Node {
	if len(arr) <= 1 {
		return arr
	}
	left := mergeSort(arr[:len(arr)/2])
	right := mergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(first, second []*Node) []*Node {
	var results []*Node
	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		if first[i].value > second[j].value {
			results = append(results, first[i])
			i++
		} else {
			results = append(results, second[j])
			j++
		}
	}
	for i < len(first) {
		results = append(results, first[i])
		i++
	}
	for j < len(second) {
		results = append(results, second[j])
		j++
	}
	return results
}