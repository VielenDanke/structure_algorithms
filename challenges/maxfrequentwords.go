package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

type Node struct {
	key   []byte
	value int
}

func (n Node) String() string {
	return fmt.Sprintf("Word: %s, amount: %d", n.key, n.value)
}

func readAllBytes(reader io.Reader) ([]byte, error) {
	return ioutil.ReadAll(reader)
}

func maxFrequentWords(data []byte, amount int) ([]*Node, error) {
	if amount == 0 {
		return []*Node{}, nil
	}
	nodes := make([]*Node, 0)
	filteredBuff := bytes.NewBuffer(make([]byte, 0))
	dataBuff := bytes.NewBuffer(data)
	isWrongLetterRepeated := false
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
			isWrongLetterRepeated = false
		} else {
			if !isWrongLetterRepeated {
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
				isWrongLetterRepeated = true
			}
		}
	}
	nodes = mergeSort(nodes)
	if len(nodes) < amount {
		amount = len(nodes)
	}
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