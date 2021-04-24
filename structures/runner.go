package main

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/graph"
)

func main() {
	wg := graph.NewWeightedGraph()
	wg.AddVertex("A")
	wg.AddVertex("B")
	wg.AddVertex("C")
	wg.AddVertex("D")
	wg.AddVertex("E")
	wg.AddVertex("F")

	wg.AddEdge("A", "B", 4)
	wg.AddEdge("A", "C", 2)
	wg.AddEdge("B", "E", 3)
	wg.AddEdge("C", "D", 2)
	wg.AddEdge("C", "F", 4)
	wg.AddEdge("D", "E", 3)
	wg.AddEdge("D", "F", 1)
	wg.AddEdge("E", "F", 1)

	fmt.Println(wg.CalculateShortestPath("A", "E"))
}
