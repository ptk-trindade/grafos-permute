package main

import (
	"fmt"
	"math/rand"
	"time"
)

func tests(adjacency [][]*Neighbor, lenVertex uint32, edges []Edge, negative_edge bool) {

	// ------ START TESTS ------
	// test 1: Calculate the distance between two nodes
	fmt.Println("\n --- TEST 1 (Distances) ---")
	if !negative_edge {
		start_id := uint32(10)
		end_id := []uint32{20, 30, 40, 50, 60} // ADD OS 0 DPS

		var paths_string string
		distance_time := time.Now()
		for i := 0; i < len(end_id); i++ {
			tree := dijkstraHeap(adjacency, start_id, end_id[i])
			path, cost := findPath(tree, start_id, end_id[i])
			fmt.Println("Distance from", start_id, "to", end_id[i], ":", cost)
			paths_string += "Path from " + fmt.Sprint(start_id) + " to " + fmt.Sprint(end_id[i]) + " (cost: " + fmt.Sprintf("%.2f", cost) + ", edges: " + fmt.Sprint(len(path)-1) + "):\n" + fmt.Sprint(path) + "\n\n"
		}
		fmt.Println("Time calculating distances:", time.Since(distance_time))

		fmt.Println("writing paths to file")
		WriteFile("paths.txt", paths_string)

		// test 2: Running Dijkstra's k times
		fmt.Println("\n --- TEST 2 (Dijkstra) ---")
		k := 100

		// create a slice of k random vertexes
		var randomVertexes []uint32 = make([]uint32, k)
		for i := 0; i < k; i++ {
			randomVertexes[i] = uint32(rand.Intn(int(lenVertex)))
		}

		fmt.Println("Run dijsktra", k, "times")
		start_heap := time.Now()
		for i := 0; i < k; i++ {
			dijkstraHeap(adjacency, randomVertexes[i], 0)
		}
		fmt.Println("\nTime dijkstra heap:", time.Since(start_heap))

		start_list := time.Now()
		for i := 0; i < k; i++ {
			dijkstraList(adjacency, randomVertexes[i], 0)
		}
		fmt.Println("\nTime dijkstra list:", time.Since(start_list))
	} else {
		fmt.Println("Graph contains negative edges, skipping tests 1 and 2")
	}

	// test 3: Finding MST
	fmt.Println("\n --- TEST 3 (MST) ---")
	start_mst := time.Now()
	mst, cost := primHeap(adjacency)
	fmt.Println("Time MST:", time.Since(start_mst))

	// transform the MST in a string
	mst_string := "(total cost: " + fmt.Sprintf("%.2f", cost) + ")\nid\t\tfather\t\tcost\n"
	mst_string += tree2text(mst)

	WriteFile("mst.txt", mst_string)

	fmt.Println(" --- END --- ")

}
