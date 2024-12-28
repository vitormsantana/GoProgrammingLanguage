package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("A", "B")
	addEdge("A", "C")
	addEdge("B", "C")

	fmt.Println("Has edge A -> B:", hasEdge("A", "B"))
	fmt.Println("Has edge A -> D:", hasEdge("A", "D"))
	fmt.Println("Has edge B -> C:", hasEdge("B", "C"))
	fmt.Println("Has edge C -> B:", hasEdge("C", "B"))
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
