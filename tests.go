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
	flows := [][]float64{
		{0, 27, 0, 21, 25, 25, 30, 56, 15, 20},
		{0, 12178, 0, 13646, 14698, 19796, 1944, 12846, 11225, 18582},
		{0, 46, 0, 32, 19, 31, 39, 10, 27, 35},
		{0, 9623, 0, 14415, 27256, 28358, 28071, 17701, 15298, 23563},
		{0, 12, 0, 39, 28, 44, 34, 42, 40, 54},
		{0, 28960, 0, 52777, 39931, 30261, 60146, 47395, 66609, 39032},
		{0, 64, 0, 80, 23, 60, 37, 24, 30, 25},
		{0, 462474, 0, 494372, 390674, 187092, 254688, 198591, 566312, 233418},
	}
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
		ford_fulkerson_time := time.Now()
		fmt.Print("Time per ff: ")
		for i := 0; i < 10; i++ {
			ff_time := time.Now()
			maxFlow, _ := fordFulkersonV2(adjacency, uint32(i), uint32(i*10))
			if maxFlow != flows[case_i][i] {
				fmt.Println("ERROR: flow", maxFlow, "!= expected", flows[case_i][i])
			}
			fmt.Print("\t", maxFlow, "(", time.Since(ff_time), ")")
		}
		fmt.Println("\nTotal time V1:", time.Since(ford_fulkerson_time))

		// VERSION 2: (2 bfs and no parallelism)
		ford_fulkerson_time = time.Now()
		fmt.Print("Time per ff: ")
		for i := 0; i < 10; i++ {
			ff_time := time.Now()
			maxFlow, _ := fordFulkersonV2(adjacency, uint32(i), uint32(i*10))
			if maxFlow != flows[case_i][i] {
				fmt.Println("ERROR: flow", maxFlow, "!= expected", flows[case_i][i])
			}
			fmt.Print("\t", maxFlow, "(", time.Since(ff_time), ")")
		}
		fmt.Println("\nTotal time V2:", time.Since(ford_fulkerson_time))

		// VERSION 3: (2 bfs and parallelism)
		ford_fulkerson_time = time.Now()
		fmt.Print("Time per ff: ")
		for i := 0; i < 10; i++ {
			ff_time := time.Now()
			maxFlow, _ := fordFulkersonV3(adjacency, uint32(i), uint32(i*10))
			if maxFlow != flows[case_i][i] {
				fmt.Println("ERROR V3: flow", maxFlow, "!= expected", flows[case_i][i])
			}
			fmt.Print("\t", maxFlow, "(", time.Since(ff_time), ")")
		}
		fmt.Println("\nTotal time V3:", time.Since(ford_fulkerson_time))
	}
	fmt.Println(" --- END --- ")
}
