package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

func main() {
	fmt.Println(" --- START ---")

	run_tests := "0"
	for run_tests != "1" && run_tests != "2" {
		fmt.Println("\nRun tests?")
		fmt.Println("1 - Yes")
		fmt.Println("2 - No")
		fmt.Scan(&run_tests)
	}

	if run_tests == "1" {
		tests()
		return
	}

	var filename string
	fmt.Print("Insert file name: ")
	fmt.Scan(&filename)

	start_time := time.Now()
	lenVertex, degrees, edges := readInitFile(filename) // reads the file and returns the vertexes and edges
	fmt.Println("Time reading files:", time.Since(start_time))

	representation_id := "0"
	for representation_id != "1" && representation_id != "2" {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Adjacency List")
		fmt.Println("2 - Adjacency Matrix")
		fmt.Scan(&representation_id)
	}

	// ----- PROCESSING -----
	var components [][]uint32
	var treeText, pathText, diameterText string
	if representation_id == "1" {
		components, treeText, pathText, diameterText = adjacencyList(lenVertex, degrees, edges)
	} else if representation_id == "2" {
		components, treeText, pathText, diameterText = adjacencyMatrix(lenVertex, edges)
	} else {
		log.Fatal("Invalid option")
	}

	// ----- EDGES.DESCRIBE() -----
	degrees = degrees[1:] // removes the first element (0)
	// Sum degrees
	sort.Slice(degrees, func(i, j int) bool { return degrees[i] < degrees[j] })

	maxDegree := uint32(degrees[len(degrees)-1])            // last element
	minDegree := uint32(degrees[0])                         // first element
	avgDegree := float32(len(edges)*2) / float32(lenVertex) // average degree
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
	WriteFile("path.txt", pathText)
	WriteFile("diameter.txt", diameterText)

	fmt.Println("Time writing files:", time.Since(write_time))
	fmt.Println("Time in execution:", time.Since(start_time))
	fmt.Println(" --- END ---")
	fmt.Println("Enter anything to exit")
	fmt.Scan(&representation_id)
}

/*
lenVertex: number of vertexes
neighCount: number of edges for each vertex
edges: list of edges {{id1, id2}, ...}
---
components: list of components {size, {vertexes}}
treeText: string with the tree (node, father, level)
pathText: string with paths
diameterText: string with the diameter of graph
*/
func adjacencyList(lenVertex uint32, neighCount []uint32, edges [][2]uint32) ([][]uint32, string, string, string) {
	// ----- CREATE LIST -----
	adjacency := make([][]uint32, lenVertex)

	log.Println("Allocating slices")
	// allocate memory for the slices (not crucial but decreases memory allocation)
	for i := uint32(0); i < lenVertex; i++ {
		adjacency[i] = make([]uint32, 0, neighCount[i])
	}

	log.Println("len flling edges: ", len(edges))
	// fill the slices
	for i := uint32(0); i < uint32(len(edges)); i++ {
		adjacency[edges[i][0]] = append(adjacency[edges[i][0]], edges[i][1])
		adjacency[edges[i][1]] = append(adjacency[edges[i][1]], edges[i][0])
	}

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
		fmt.Println("Insert the id of the start node:")
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
		tree = bfsList(adjacency, start_id)
	} else if search_id == "2" {
		tree = dfsList(adjacency, start_id)
	} else {
		log.Fatal("Invalid option")
	}

	treeText := "node\tfather\tlevel"
	for _, node := range tree {
		treeText += fmt.Sprintf("\n%d\t\t%d\t\t%d", node[0], node[1], node[2])
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

	var components [][]uint32
	if find_compnents == "1" {
		components_time := time.Now()
		components = findComponentsList(adjacency)
		fmt.Println("components: ", components)

		fmt.Println("Time finding components:", time.Since(components_time))
	}

	// ----- FIND PATH -----
	fmt.Println("\nFind path:")
	var pathText string
	var path_start_str string
	var path_end_str string
	var end_id64 uint64
	for {
		valid = false
		for !valid {
			fmt.Println("\nInser the id of the start node (or 0 to exit):")
			fmt.Scan(&path_start_str)

			start_id64, err = strconv.ParseUint(path_start_str, 10, 32)
			if err == nil && uint32(start_id64) < lenVertex {
				valid = true
			}
		}

		if path_start_str == "0" {
			break
		}

		valid = false
		for !valid {
			fmt.Println("\nInsert the id of the end node:")
			fmt.Scan(&path_end_str)

			end_id64, err = strconv.ParseUint(path_end_str, 10, 32)
			if err == nil && uint32(end_id64) < lenVertex {
				valid = true
			}
		}

		path_time := time.Now()

		path := findPathList(adjacency, uint32(start_id64), uint32(end_id64))

		pathText += "Path from " + path_start_str + " to " + path_end_str + "(distance: " + strconv.Itoa(len(path)-1) + ")\n"
		for i := len(path) - 1; i >= 0; i-- {
			pathText += fmt.Sprintf("%d -> ", path[i])
		}
		pathText += "\n\n"

		fmt.Println("Time finding path:", time.Since(path_time))
	}

	// ----- FIND DIAMETER -----
	fmt.Println("\nFind diameter:")
	quick_run := "0"
	for quick_run != "1" && quick_run != "2" {
		fmt.Println("\nRun in quick mode (diameter found might be lower than the real):")
		fmt.Println("1 - Yes")
		fmt.Println("2 - No")
		fmt.Scan(&quick_run)
	}

	var diameter, v1, v2 uint32
	diameter_time := time.Now()
	if quick_run == "1" {
		var component_vertex []uint32
		for _, component := range components { // get one vertex of each component
			component_vertex = append(component_vertex, component[0])
		}
		diameter, v1, v2 = findDiameterQuickList(adjacency, component_vertex)
	} else {
		diameter, v1, v2 = findDiameterList(adjacency)
	}
	diameterText := "Diameter: " + strconv.Itoa(int(diameter)) + "\n"
	diameterText += "From vertex " + strconv.Itoa(int(v1)) + " to " + strconv.Itoa(int(v2)) + "\n"
	fmt.Println("Time finding diameter:", time.Since(diameter_time))

	return components, treeText, pathText, diameterText
}

func adjacencyMatrix(lenVertex uint32, edges [][2]uint32) ([][]uint32, string, string, string) {
	// ----- CREATE MATRIX -----
	adjacency := make([][]uint8, lenVertex)

	// allocate memory for the slices (not crucial decreases memory allocation)
	for i := uint32(0); i < lenVertex; i++ {
		adjacency[i] = make([]uint8, lenVertex)
	}

	log.Println("len edges: ", len(edges))
	// fill the slices
	for i := uint32(0); i < uint32(len(edges)); i++ {
		adjacency[edges[i][0]][edges[i][1]] = 1
		adjacency[edges[i][1]][edges[i][0]] = 1
	}

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
		tree = bfsMatrix(adjacency, start_id)
	} else if search_id == "2" {
		tree = dfsMatrix(adjacency, start_id)
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
		components = findComponentsMatrix(adjacency)

		fmt.Println("Time finding components:", time.Since(components_time))
	}

	// ----- FIND PATH -----
	fmt.Println("\nFind path:")
	var pathText string
	var path_start_str string
	var path_end_str string
	var end_id64 uint64
	for {
		valid = false
		for !valid {
			fmt.Println("\nInsert the id of the start node (or 0 to exit):")
			fmt.Scan(&path_start_str)

			start_id64, err = strconv.ParseUint(path_start_str, 10, 32)
			if err == nil && uint32(start_id64) < lenVertex {
				valid = true
			}
		}

		if path_start_str == "0" {
			break
		}

		valid = false
		for !valid {
			fmt.Println("\nInser the id of the end node:")
			fmt.Scan(&path_end_str)

			end_id64, err = strconv.ParseUint(path_end_str, 10, 32)
			if err == nil && uint32(end_id64) < lenVertex {
				valid = true
			}
		}

		path_time := time.Now()

		path := findPathMatrix(adjacency, uint32(start_id64), uint32(end_id64))

		pathText += "Path from " + path_start_str + " to " + path_end_str + "(distance: " + strconv.Itoa(len(path)-1) + ")\n"
		for i := len(path) - 1; i >= 0; i-- {
			pathText += fmt.Sprintf("%d -> ", path[i])
		}
		pathText += "\n\n"

		fmt.Println("Time finding path:", time.Since(path_time))
	}

	// ----- FIND DIAMETER -----
	fmt.Println("\nFind diameter:")
	quick_run := "0"
	for quick_run != "1" && quick_run != "2" {
		fmt.Println("Run in quick mode (diameter found might be lower than the real):")
		fmt.Println("1 - Yes")
		fmt.Println("2 - No")
		fmt.Scan(&quick_run)
	}

	var diameter, v1, v2 uint32
	diameter_time := time.Now()
	if quick_run == "1" {
		var component_vertex []uint32
		for _, component := range components { // get one vertex of each component
			component_vertex = append(component_vertex, component[0])
		}
		diameter, v1, v2 = findDiameterQuickMatrix(adjacency, component_vertex)
	} else {
		diameter, v1, v2 = findDiameterMatrix(adjacency)
	}
	diameterText := "Diameter: " + strconv.Itoa(int(diameter)) + "\n"
	diameterText += "From vertex " + strconv.Itoa(int(v1)) + " to " + strconv.Itoa(int(v2)) + "\n"
	fmt.Println("Time finding diameter:", time.Since(diameter_time))

	return components, treeText, pathText, diameterText
}
