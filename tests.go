package main

import (
	"fmt"
	// "math/rand"
	"log"
	"time"
)

func tests() { //adjacency [][]*Edge

	// ------ START TESTS ------
	// test 1: Calculate the distance between two nodes
	log.Println("\n --- STARTING TESTS ---")
	test_files := []string{"grafo_rf_1.txt", "grafo_rf_2.txt", "grafo_rf_3.txt", "grafo_rf_4.txt", "grafo_rf_5.txt", "grafo_rf_6.txt", "grafo_rf_7.txt", "grafo_rf_8.txt"}
	run_n_times := 10
	for case_i, filename := range test_files {
		fmt.Println("--- Case", case_i+1, " ---")
		lenVertex, degrees, edges := readInitFile(filename)
		adjacency := make([][]*Edge, lenVertex)

		// allocate memory for the slices (not crucial but decreases memory allocation)
		for i := uint32(0); i < lenVertex; i++ {
			adjacency[i] = make([]*Edge, 0, degrees[i])
		}

		// create the adjacency list
		var edgeMap map[string]*Edge = make(map[string]*Edge)
		for _, raw_edge := range edges {
			new_edge := Edge{raw_edge.vertex1, raw_edge.vertex2, raw_edge.weight, nil}    // create the edge
			edgeMap[fmt.Sprintf("%d_%d", raw_edge.vertex1, raw_edge.vertex2)] = &new_edge // add the edge to the map

			comp_edge, exists := edgeMap[fmt.Sprintf("%d_%d", raw_edge.vertex2, raw_edge.vertex1)] // check if complementary edge exists
			if exists {
				new_edge.comp = comp_edge
				new_edge.comp = &new_edge
			}

			adjacency[raw_edge.vertex1] = append(adjacency[raw_edge.vertex1], &new_edge)

		}

		// VERSION 1: (1 bfs and no parallelism)
		times_list := make([]time.Duration, run_n_times)
		var maxFlow float64
		ford_fulkerson_time := time.Now()
		for i := 0; i < run_n_times; i++ {
			ff_time := time.Now()
			maxFlow, _ = fordFulkerson(adjacency, 1, 2)
			// if maxFlow != flows[case_i][i] {
			// 	fmt.Println("ERROR: flow", maxFlow, "!= expected", flows[case_i][i])
			// }
			times_list[i] = time.Since(ff_time)
		}
		total_time := time.Since(ford_fulkerson_time)
		fmt.Print("maxFlow", maxFlow, " >\t")
		for _, t := range times_list {
			fmt.Print("\t", t)
		}
		fmt.Println("\nTotal time V1:", total_time)

		// VERSION 2: (2 bfs and no parallelism)
		ford_fulkerson_time = time.Now()
		for i := 0; i < run_n_times; i++ {
			ff_time := time.Now()
			maxFlow, _ = fordFulkersonV2(adjacency, 1, 2)
			// if maxFlow != flows[case_i][i] {
			// 	fmt.Println("ERROR: flow", maxFlow, "!= expected", flows[case_i][i])
			// }
			times_list[i] = time.Since(ff_time)
		}
		total_time = time.Since(ford_fulkerson_time)
		fmt.Print("maxFlow", maxFlow, " >\t")
		for _, t := range times_list {
			fmt.Print("\t", t)
		}
		fmt.Println("\nTotal time V2:", total_time)
	}
	fmt.Println(" --- END --- ")
}
