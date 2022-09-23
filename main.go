package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

func main() {
	filename := "MST.txt"

	fmt.Println(" --- START ---")

	representation_id := "0"
	for representation_id != "1" && representation_id != "2" {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Adjacency List")
		fmt.Println("2 - Adjacency Matrix")
		fmt.Scan(&representation_id)
	}

	start_time := time.Now()
	lenVertex, degrees, edges := readInitFile(filename) // reads the file and returns the vertexes and edges
	fmt.Println("Time reading files:", time.Since(start_time))

	// ----- PROCESSING -----
	var components [][]uint32
	var treeText string

	if representation_id == "1" {
		components, treeText = adjacencyList(lenVertex, degrees, edges)
	} else if representation_id == "2" {
		// degrees, components := adjacencyMatrix(lenVertex, neighCount, edges)
	} else {
		log.Fatal("Invalid option")
	}

	// CALCULATE MEDIANS
	degrees = degrees[1:] // removes the first element (0)
	// Sum degrees
	sort.Slice(degrees, func(i, j int) bool { return degrees[i] < degrees[j] })
	sum := uint32(0)
	for _, degree := range degrees { // for each vertex
		sum += degree
	}

	maxDegree := uint32(degrees[len(degrees)-1])      // last element
	minDegree := uint32(degrees[0])                   // first element
	avgDegree := float32(sum) / float32(len(degrees)) // average degree
	var medianDegree float32
	if len(degrees)%2 == 0 {
		medianDegree = float32(degrees[len(degrees)/2-1]+degrees[len(degrees)/2]) / 2
	} else {
		medianDegree = float32(degrees[len(degrees)/2])
	}

	// ----- WRITE OUTPUT -----
	write_time := time.Now()

	//lenVertex-1 : doesn't count the null vertex (0)
	writeOutput(lenVertex-1, uint32(len(edges)), minDegree, maxDegree, avgDegree, medianDegree, components)
	WriteFile("tree.txt", treeText)

	fmt.Println("Time writing files:", time.Since(write_time))
	fmt.Println("Time in execution:", time.Since(start_time))
	fmt.Println(" --- END ---")
}

/*
lenVertex: number of vertexes
neighCount: number of edges for each vertex
edges: list of edges {{id1, id2}, ...}
---
components: list of components {size, {vertexes}}
treeText: string with the tree (node, father, level)
*/
func adjacencyList(lenVertex uint32, neighCount []uint32, edges [][2]uint32) ([][]uint32, string) {
	// ----- CREATE LIST -----
	adjacency := make([][]uint32, lenVertex)

	// allocate memory for the slices
	for i := uint32(0); i < lenVertex; i++ {
		log.Println("node: ", i, " neigh_qnt: ", neighCount[i])
		adjacency[i] = make([]uint32, 0, neighCount[i])
	}

	log.Println("len edges: ", len(edges))
	// fill the slices
	for i := uint32(0); i < uint32(len(edges)); i++ {
		log.Println("neighbors: ", edges[i][0], edges[i][1])
		adjacency[edges[i][0]] = append(adjacency[edges[i][0]], edges[i][1])
		adjacency[edges[i][1]] = append(adjacency[edges[i][1]], edges[i][0])
		fmt.Println(adjacency)
	}
	fmt.Println("Lista logo apos criada: ", adjacency)

	// ----- GET USER INPUT -----
	// pick BFS or DFS
	search_id := "0"
	for search_id != "1" && search_id != "2" {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Breadth First Search (BFS)")
		fmt.Println("2 - Depth First Search (DFS)")
		fmt.Scan(&search_id)
	}

	// pick start vertex
	var start_id uint32
	var start_id64 uint64
	var start_id_str string
	var err error
	valid := false
	for !valid {
		fmt.Println("\nInser the id of the start node:")
		fmt.Scan(&start_id_str)

		start_id64, err = strconv.ParseUint(start_id_str, 10, 32)
		if err == nil && uint32(start_id64) < lenVertex {
			valid = true
		}
	}
	start_id = uint32(start_id64)

	// ----- FIND TREE -----
	start_time := time.Now()
	log.Println("start_id: ", start_id)
	var tree [][3]uint32
	if search_id == "1" {
		fmt.Print("Lista que foi passada: ", adjacency)
		tree = bfsList(adjacency, start_id)
	} else if search_id == "2" {
		tree = dfsList(adjacency, start_id)
	} else {
		log.Fatal("Invalid option")
	}

	treeText := "node\tfather\tlevel"
	components := [][]uint32{{}} // {{vertex1, vextex2, ...}, ...}
	for _, node := range tree {
		treeText += fmt.Sprintf("\n%d\t\t%d\t\t%d", node[0], node[1], node[2])
		components[0] = append(components[0], node[0])
	}

	fmt.Println("Time processing search tree:", time.Since(start_time))

	// ----- FIND COMPONENTS -----
	find_compnents := "0"
	for find_compnents != "1" && find_compnents != "2" {
		fmt.Println("\nFind all the components:")
		fmt.Println("1 - Yes")
		fmt.Println("2 - No")
		fmt.Scan(&find_compnents)
	}

	if find_compnents == "1" {
		components_time := time.Now()
		visited_vec := make([]*Node, lenVertex)
		visited_lis := List{nil, nil}

		// fill visited list
		for i := uint32(1); i < lenVertex; i++ {
			node := visited_lis.Push(i)
			visited_vec[i] = node
		}

		// while there are unvisited nodes
		for {
			visited_lis.Display() // log

			//clear visited nodes from list
			for _, tuple := range tree {
				visited_lis.Remove(visited_vec[tuple[0]])
			}

			if visited_lis.head == nil {
				break
			}

			// find new component
			start_id = visited_lis.head.value

			fmt.Println("start_id: ", start_id)
			tree = bfsList(adjacency, start_id)
			fmt.Println("tree: ", tree)

			// create components list
			components = append(components, []uint32{})
			for _, node := range tree {
				components[len(components)-1] = append(components[len(components)-1], node[0])
			}
			fmt.Println("components: ", components)
		}

		fmt.Println("Time finding components:", time.Since(components_time))
	}

	return components, treeText
}
