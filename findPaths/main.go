package main

import (
	"fmt"
)

// The number of vertexes in the graph
const lenVertex int = 7 // If > 255 change uint8 in Vertex structure

func main() {
	fmt.Println(" --- START ---")

	createVertexes(lenVertex)
	populateEdges(AllVertex)

	visited := make(map[int]bool)
	visitNeighbor(AllVertex[5], AllVertex[7], 0, visited)
	fmt.Println(" --- END ---")
}

func visitNeighbor(current *Vertex, end *Vertex, cost float32, visited map[int]bool) bool {

	fmt.Println("Visiting", current.id, "with cost", cost)

	if current == end {
		fmt.Println("Concluded!")
		fmt.Println("cost:", cost)
		fmt.Println(end.id)
		return true
	}

	_, ok := visited[current.id]
	if ok {
		fmt.Println("Already visited", current.id)
		return false
	}

	visited[current.id] = true

	for _, edge := range current.edges {
		if visitNeighbor(edge.neighbor, end, cost+edge.weight, visited) {
			fmt.Println(current.id)
			return true
		}
	}

	return false
}

// func dijkstra(start *Vertex, end *Vertex) {
// 	// Initialize all vertexes
// 	for _, vertex := range AllVertex {
// 		vertex.cost = 999999
// 		vertex.parent = nil
// 	}

// }
