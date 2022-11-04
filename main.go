package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

func main() {
	fmt.Println(" --- START ---")

	fmt.Print("graph number (1 to 5)\ninsert: ")
	var graph_num string
	fmt.Scan(&graph_num)

	filename := "grafo_W_" + graph_num + "_1.txt"
	// filename = "graph_1.txt"
	fmt.Println("filename:", filename)
	// fmt.Print("Insert file name: ")
	// fmt.Scan(&filename)

	start_time := time.Now()
	lenVertex, degrees, edges := readInitFile(filename) // reads the file and returns the vertexes and edges
	fmt.Println("Time reading files:", time.Since(start_time))

	adjacency := make([][]*Neighbor, lenVertex)

	log.Println("Allocating slices")
	// allocate memory for the slices (not crucial but decreases memory allocation)
	for i := uint32(0); i < lenVertex; i++ {
		adjacency[i] = make([]*Neighbor, 0, degrees[i])
	}

	log.Println("len flling edges: ", len(edges))
	// fill the slices
	negative_edge := false
	for i := uint32(0); i < uint32(len(edges)); i++ {
		if edges[i].weight < 0 {
			negative_edge = true
		}
		adjacency[edges[i].vertex1] = append(adjacency[edges[i].vertex1], &Neighbor{vertex_id: edges[i].vertex2, weight: edges[i].weight})
		adjacency[edges[i].vertex2] = append(adjacency[edges[i].vertex2], &Neighbor{vertex_id: edges[i].vertex1, weight: edges[i].weight})
	}

	run_tests := "0"
	for run_tests != "1" && run_tests != "2" {
		fmt.Println("\nRun tests?")
		fmt.Println("1 - Yes")
		fmt.Println("2 - No")
		fmt.Scan(&run_tests)
	}

	if run_tests == "1" {
		tests(adjacency, lenVertex, edges, negative_edge)
		return
	}

	// ----- EDGES.DESCRIBE() -----
	degrees = degrees[1:] // removes the first element (0)
	// Sum degrees
	sort.Slice(degrees, func(i, j int) bool { return degrees[i] < degrees[j] }) // ascending

	maxDegree := uint32(degrees[len(degrees)-1])            // last element
	minDegree := uint32(degrees[0])                         // first element
	avgDegree := float32(len(edges)*2) / float32(lenVertex) // average degree
	var medianDegree float32
	if len(degrees)%2 == 0 {
		medianDegree = float32(degrees[len(degrees)/2-1]+degrees[len(degrees)/2]) / 2
	} else {
		medianDegree = float32(degrees[len(degrees)/2])
	}

	sort.Slice(edges, func(i, j int) bool { return edges[i].weight < edges[j].weight }) // ascending

	maxWeight := edges[len(edges)-1].weight // last element
	minWeight := edges[0].weight            // first element
	var medianWeight float32
	if len(edges)%2 == 0 {
		medianWeight = float32(edges[len(edges)/2-1].weight+edges[len(edges)/2].weight) / 2
	} else {
		medianWeight = float32(edges[len(edges)/2].weight)
	}

	avgWeight := float64(0)
	for i := 0; i < len(edges); i++ {
		avgWeight += edges[i].weight
	}
	avgWeight /= float64(len(edges))

	// ----- WRITE OUTPUT -----
	write_time := time.Now()

	//lenVertex-1 : doesn't count the null vertex (0)
	lenVertex = lenVertex - 1
	lenEdges := uint32(len(edges))
	// ----- NUMBER OF VERTEX -----
	output := "Num of vertexes: " + strconv.Itoa(int(lenVertex)) + "\n"

	// ----- NUMBER OF EDGES -----
	output += "Num of edges: " + strconv.Itoa(int(lenEdges)) + "\n"

	// ----- DEGREES -----
	output += "\nDEGREES:\n"
	output += "max degree: " + strconv.Itoa(int(maxDegree)) + "\n"
	output += "min degree: " + strconv.Itoa(int(minDegree)) + "\n"
	output += "average degree: " + strconv.FormatFloat(float64(avgDegree), 'f', 2, 32) + "\n"
	output += "medianDegree degree: " + strconv.FormatFloat(float64(medianDegree), 'f', 2, 32) + "\n"

	// ----- WEIGHTS -----
	output += "\nWEIGHTS:\n"
	output += "max weight: " + strconv.FormatFloat(float64(maxWeight), 'f', 2, 64) + "\n"
	output += "min weight: " + strconv.FormatFloat(float64(minWeight), 'f', 2, 64) + "\n"
	output += "average weight: " + strconv.FormatFloat(avgWeight, 'f', 2, 64) + "\n"
	output += "median weight: " + strconv.FormatFloat(float64(medianWeight), 'f', 2, 64) + "\n"

	WriteFile("output.txt", output)

	fmt.Println("Time writing files:", time.Since(write_time))
	var start_str string
	var end_str string
	var start64 uint64
	var end64 uint64
	var error error = errors.New("no value assigned")
	for error != nil {
		fmt.Println("\nPick a starting vertex")
		fmt.Scan(&start_str)
		start64, error = strconv.ParseUint(start_str, 10, 32)
	}

	error = errors.New("no value assigned")
	for error != nil {
		fmt.Println("\nPick an ending vertex")
		fmt.Scan(&end_str)
		end64, error = strconv.ParseUint(end_str, 10, 32)
	}

	// DIJKSTRA BINARY HEAP
	dijkstraTime := time.Now()
	tree_heap := dijkstraHeap(adjacency, uint32(start64), uint32(end64))
	fmt.Println("Time dijkstra Heap:", time.Since(dijkstraTime))

	heap_string := "id\t\tfather\t\tcost\n"
	heap_string += tree2text(tree_heap)
	WriteFile("dijkstraHeap.txt", heap_string)

	// DIJKSTRA LIST
	dijkstraTime = time.Now()
	tree_list := dijkstraList(adjacency, uint32(start64), uint32(end64))
	fmt.Println("Time dijkstra List:", time.Since(dijkstraTime))

	list_string := "id\t\tfather\t\tcost\n"
	list_string += tree2text(tree_list)
	WriteFile("dijkstraList.txt", list_string)

	// MST
	start_mst := time.Now()
	mst, cost := primHeap(adjacency)
	fmt.Println("Time MST:", time.Since(start_mst))
	mst_string := "(total cost: " + fmt.Sprintf("%.2f", cost) + ")\nid\t\tfather\t\tcost\n"
	mst_string += tree2text(mst)

	WriteFile("mst.txt", mst_string)

	fmt.Println("Time writing files:", time.Since(write_time))
	fmt.Println("Time in execution:", time.Since(start_time))
	fmt.Println(" --- END ---")
	fmt.Println("Enter anything to exit")
}
