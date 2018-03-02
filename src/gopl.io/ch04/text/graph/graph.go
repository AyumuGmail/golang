package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("たて", "よこ")
	internal := graph["たて"]
	fmt.Printf("%v\n", internal["よこ"])
	fmt.Printf("%v\n", graph["よこ"]["たて"])
	fmt.Printf("%v\n", graph["たて"]["よこ"])
}
func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
