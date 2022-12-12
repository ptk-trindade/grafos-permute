package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func main() {
	fmt.Println(" --- START ---")

	run_tests := "0"
	for run_tests != "1" && run_tests != "2" {
		fmt.Println("\nRun ALL tests? (Hint: Don't)")
		fmt.Println("1) Yes")
		fmt.Println("2) No")
		fmt.Scan(&run_tests)
	}
	if run_tests == "1" {
		tests()
		return
	}

	fmt.Print("Insert file name (with extension) \nfilename: ")
	var filename string
	fmt.Scan(&filename)

	start_time := time.Now()
	lenVertex, degrees, edges := readInitFile(filename) // reads the file and returns the vertexes and edges
	fmt.Println("Time reading files:", time.Since(start_time))

	adjacency := make([][]*Edge, lenVertex)

	// allocate memory for the slices (not crucial but decreases memory allocation)
	for i := uint32(0); i < lenVertex; i++ {
		adjacency[i] = make([]*Edge, 0, degrees[i])
	}

	var directed int
	for directed != 1 { // && directed != 2
		fmt.Println("\nIs the graph directed?")
		fmt.Println("1) Yes")
		fmt.Println("2) No")
		fmt.Scan(&directed)
	}

	// var negative_edge = false
	// create the adjacency list
	if directed == 1 {
		// var edgeMap map[string]*Edge = make(map[string]*Edge)
		for _, raw_edge := range edges {
			new_edge := Edge{raw_edge.vertex1, raw_edge.vertex2, raw_edge.weight, nil} // create the edge
			// edgeMap[fmt.Sprintf("%d_%d", raw_edge.vertex1, raw_edge.vertex2)] = &new_edge // add the edge to the map

			// comp_edge, exists := edgeMap[fmt.Sprintf("%d_%d", raw_edge.vertex2, raw_edge.vertex1)] // check if complementary edge exists
			// if exists {
			// new_edge.comp = comp_edge
			// new_edge.comp = &new_edge
			// }

			adjacency[raw_edge.vertex1] = append(adjacency[raw_edge.vertex1], &new_edge)

		}
	} else {
		for i := uint32(0); i < uint32(len(edges)); i++ {
			// if edges[i].weight < 0 {
			// 	negative_edge = true
			// }
			edge1 := &Edge{edges[i].vertex1, edges[i].vertex2, edges[i].weight, nil}
			edge2 := &Edge{edges[i].vertex2, edges[i].vertex1, edges[i].weight, edge1}
			edge1.comp = edge2

			adjacency[edges[i].vertex1] = append(adjacency[edges[i].vertex1], edge1)
			adjacency[edges[i].vertex2] = append(adjacency[edges[i].vertex2], edge2)
		}
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
	var output bytes.Buffer
	output.WriteString("Num of vertexes: " + strconv.Itoa(int(lenVertex)) + "\n")

	// ----- NUMBER OF EDGES -----
	output.WriteString("Num of edges: " + strconv.Itoa(int(lenEdges)) + "\n")

	// ----- DEGREES -----
	output.WriteString("\nDEGREES:\n")
	output.WriteString("max degree: " + strconv.Itoa(int(maxDegree)) + "\n")
	output.WriteString("min degree: " + strconv.Itoa(int(minDegree)) + "\n")
	output.WriteString("average degree: " + strconv.FormatFloat(float64(avgDegree), 'f', 2, 32) + "\n")
	output.WriteString("medianDegree degree: " + strconv.FormatFloat(float64(medianDegree), 'f', 2, 32) + "\n")

	// ----- WEIGHTS -----
	output.WriteString("\nWEIGHTS:\n")
	output.WriteString("max weight: " + strconv.FormatFloat(float64(maxWeight), 'f', 2, 64) + "\n")
	output.WriteString("min weight: " + strconv.FormatFloat(float64(minWeight), 'f', 2, 64) + "\n")
	output.WriteString("average weight: " + strconv.FormatFloat(avgWeight, 'f', 2, 64) + "\n")
	output.WriteString("median weight: " + strconv.FormatFloat(float64(medianWeight), 'f', 2, 64) + "\n")

	WriteFile("output.txt", output.String())

	fmt.Println("Time writing output.txt file:", time.Since(write_time))
	var source uint32
	var sink uint32

	var run int = 1
	for run == 1 || run == 2 {
		fmt.Println("\nRun Ford Fulkerson?")
		fmt.Println("1) Run Version 1")
		fmt.Println("2) Run Version 2")
		fmt.Println("3) Exit")
		fmt.Scan(&run)

		if run == 1 { // ----- FORD FULKERSON V1 -----

			fmt.Println("\nPick a source vertex")
			fmt.Scan(&source)

			fmt.Println("\nPick a sink vertex")
			fmt.Scan(&sink)

			if source == 0 || sink == 0 {
				fmt.Println("Warning: Vertex 0 does not exist")
			}

			fmt.Println("--- Finding max flow (fordFulkerson) from", source, "to", sink)
			ford_fulkerson_time := time.Now()
			max_flow, edgeFlow := fordFulkerson(adjacency, source, sink)
			fmt.Println("Time running Ford Fulkerson:", time.Since(ford_fulkerson_time))
			fmt.Println("Max flow:", max_flow)

			var flow_output bytes.Buffer
			flow_output.WriteString("Max flow: " + strconv.Itoa(int(max_flow)) + "\n")
			flow_output.WriteString("Edges\nFrom\t\tTo\t\tCapacity/Flow\n")
			for _, edgeF := range edgeFlow {
				str := fmt.Sprintf("%d\t\t%d\t\t%.2f/%.2f\n", edgeF.edge.origin, edgeF.edge.dest, edgeF.edge.weight, edgeF.flow)
				flow_output.WriteString(str)

			}
			WriteFile("ford_fulkersonV1.txt", flow_output.String())
		} else if run == 2 { // ----- FORD FULKERSON V2 -----
			fmt.Println("\nPick a source vertex")
			fmt.Scan(&source)

			fmt.Println("\nPick a sink vertex")
			fmt.Scan(&sink)

			if source == 0 || sink == 0 {
				fmt.Println("Warning: Vertex 0 does not exist")
			}

			fmt.Println("--- Finding max flow (fordFulkersonV2) from", source, "to", sink)
			ford_fulkerson_time := time.Now()
			max_flow, edgeFlow := fordFulkersonV2(adjacency, source, sink)
			fmt.Println("Time running Ford Fulkerson:", time.Since(ford_fulkerson_time))
			fmt.Println("Max flow:", max_flow)

			var flow_output bytes.Buffer
			flow_output.WriteString("Max flow: " + strconv.Itoa(int(max_flow)) + "\n")
			flow_output.WriteString("Edges\nFrom\t\tTo\t\tCapacity/Flow\n")
			for _, edgeF := range edgeFlow {
				str := fmt.Sprintf("%d\t\t%d\t\t%.2f/%.2f\n", edgeF.edge.origin, edgeF.edge.dest, edgeF.edge.weight, edgeF.flow)
				flow_output.WriteString(str)

			}
			WriteFile("ford_fulkersonV2.txt", flow_output.String())
		}
	}

	fmt.Println("Time in execution:", time.Since(start_time))
	fmt.Println(" --- END ---")
}
