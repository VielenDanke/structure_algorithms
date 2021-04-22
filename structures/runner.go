package main

import (
	"fmt"
	hashtable "github.com/vielendanke/structure_algorithms/structures/hash_table"
	"log"
)

type obj struct {
	id       int
	username string
}

func (o *obj) Equal(val interface{}) bool {
	incObj, ok := val.(*obj)
	if !ok {
		return false
	}
	return o.id == incObj.id
}

func (o *obj) Hash() int {
	return o.id
}

func main() {
	hm, err := hashtable.NewHashMap(16, func(leftKey interface{}, rightKey interface{}) bool {
		f := leftKey.(*obj)
		s := rightKey.(*obj)
		return f.id > s.id
	})
	if err != nil {
		log.Fatalln(err)
	}
	counter := 0

	for counter < 5000 {
		hm.Put(&obj{id: counter, username: "bla"}, "New World!")
		hm.Put(&obj{id: counter + 1, username: "vo"}, "Old World!")
		hm.Put(&obj{id: counter + 2, username: "ded"}, "Between World!")
		hm.Put(&obj{id: counter + 3, username: "keep"}, "Big World!")
		counter += 5
	}

	fmt.Println(hm.Contains(&obj{id: 301, username: "vo"}))
}
