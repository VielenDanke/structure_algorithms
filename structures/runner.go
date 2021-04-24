package main

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/common"
	hashtable "github.com/vielendanke/structure_algorithms/structures/hash_table"
	"log"
)

type node struct {
	id  int
	val string
}

func (n *node) Equal(v common.EqualHashRule) bool {
	return n.id == v.(*node).id
}

func (n *node) Hash() int {
	return n.id * len(n.val)
}

func main() {
	ht, err := hashtable.NewHashMap(16, func(leftKey interface{}, rightKey interface{}) bool {
		n := leftKey.(*node)
		r := rightKey.(*node)
		return n.id > r.id
	})
	if err != nil {
		log.Fatalln(err)
	}
	ht.Put(&node{
		id:  1,
		val: "avb",
	}, "bla")
	ht.Put(&node{
		id:  2,
		val: "bbb",
	}, "bla")
	ht.Put(&node{
		id:  3,
		val: "eee",
	}, "bla")

	fmt.Println(ht.Get(&node{id: 1, val: "bbb"}))
}
