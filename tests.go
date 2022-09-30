package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {

	filename := "grafo_1.txt"
	// testsList(filename)
	testsMatrix(filename)
}

func testsList(filename string) {
	fmt.Println(" --- START LIST---")

	// fmt.Print("Insert file name: ")
	// fmt.Scan(&filename)

	start_time := time.Now()
	lenVertex, neighCount, edges := readInitFile(filename) // reads the file and returns the vertexes and edges
	fmt.Println("Time reading files:", time.Since(start_time))

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

	fmt.Print("Check memory usage (adjacency list): ")
	fmt.Scan(&filename)

	// ------ START TESTS ------
	// test 1: see allocated space
	fmt.Println("\n --- TEST 1 (SPACE) ---")

	// test 2: run 1000 BFS
	fmt.Println("\n --- TEST 2 (BFS) ---")
	// create a slice of 1000 random vertexes
	var randomVertexes []uint32 = make([]uint32, 1000)
	for i := 0; i < 1000; i++ {
		randomVertexes[i] = uint32(rand.Intn(int(lenVertex)))
	}

	// run 1000 BFS
	start_time = time.Now()
	for i := 0; i < 1000; i++ {
		bfsList(adjacency, randomVertexes[i])
	}
	fmt.Println("Time 1000 BFS:", time.Since(start_time))

	// test 3: run 1000 DFS
	fmt.Println("\n --- TEST 3 (DFS) ---")
	start_time = time.Now()
	for i := 0; i < 1000; i++ {
		dfsList(adjacency, randomVertexes[i])
	}
	fmt.Println("Time 1000 DFS:", time.Since(start_time))

	// test 4: Get father of vertices 10, 20 and 30, starting on vertex 1, 2 and 3
	fmt.Println("\n --- TEST 4 (Fathers) ---")

	fmt.Println("BFS")
	tree := bfsList(adjacency, 1)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("1. father of", v[0], "is", v[1])
		}
	}

	tree = bfsList(adjacency, 2)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("2. father of", v[0], "is", v[1])
		}
	}

	tree = bfsList(adjacency, 3)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("3. father of", v[0], "is", v[1])
		}
	}

	fmt.Println("DFS")
	tree = dfsList(adjacency, 1)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("1. father of", v[0], "is", v[1])
		}
	}

	tree = dfsList(adjacency, 2)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("2. father of", v[0], "is", v[1])
		}
	}

	tree = dfsList(adjacency, 3)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("3. father of", v[0], "is", v[1])
		}
	}

	// test 5: distance between vertex (10, 20) (10, 30), (20, 30)
	fmt.Println("\n --- TEST 5 (Distances) ---")

	path := findPathList(adjacency, 10, 20)
	fmt.Println("Distance between 10 and 20: ", len(path)-1)

	path = findPathList(adjacency, 10, 30)
	fmt.Println("Distance between 10 and 30: ", len(path)-1)

	path = findPathList(adjacency, 20, 30)
	fmt.Println("Distance between 20 and 30: ", len(path)-1)

	// test 6: Connected components
	fmt.Println("\n --- TEST 6 (Connected components) ---")

	components_time := time.Now()
	components := findComponentsList(adjacency)

	fmt.Println("Time finding components:", time.Since(components_time))
	fmt.Println("Number of components:", len(components))

	// test 7: Diameter
	fmt.Println("\n --- TEST 7 (Diameter) ---")

	diameter_time := time.Now()
	diameter, _, _ := findDiameterList(adjacency)

	fmt.Println("Time finding diameter (slow):", time.Since(diameter_time))
	fmt.Println("Diameter:", diameter)

	diameter_time = time.Now()
	diameter, _, _ = findDiameterQuickList(adjacency)

	fmt.Println("Time finding diameter (fast):", time.Since(diameter_time))
	fmt.Println("Diameter:", diameter)

	fmt.Println("\n --- END TESTS ---")
}

func testsMatrix(filename string) {
	fmt.Println(" --- START MATRIX---")

	// fmt.Print("Insert file name: ")
	// fmt.Scan(&filename)

	start_time := time.Now()
	lenVertex, _, edges := readInitFile(filename) // reads the file and returns the vertexes and edges
	fmt.Println("Time reading files:", time.Since(start_time))

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

	fmt.Print("Check memory usage (adjacency list): ")
	fmt.Scan(&filename)
	// ------ START TESTS ------
	// test 1: see allocated space
	fmt.Println("\n --- TEST 1 (SPACE) ---")

	// test 2: run 1000 BFS
	fmt.Println("\n --- TEST 2 (BFS) ---")
	// create a slice of 1000 random vertexes
	var randomVertexes []uint32 = make([]uint32, 1000)
	for i := 0; i < 1000; i++ {
		randomVertexes[i] = uint32(rand.Intn(int(lenVertex)))
	}

	// run 1000 BFS
	start_time = time.Now()
	for i := 0; i < 1000; i++ {
		bfsMatrix(adjacency, randomVertexes[i])
	}
	fmt.Println("Time 1000 BFS:", time.Since(start_time))

	// test 3: run 1000 DFS
	fmt.Println("\n --- TEST 3 (DFS) ---")
	start_time = time.Now()
	for i := 0; i < 1000; i++ {
		dfsMatrix(adjacency, randomVertexes[i])
	}
	fmt.Println("Time 1000 DFS:", time.Since(start_time))

	// test 4: Get father of vertices 10, 20 and 30, starting on vertex 1, 2 and 3
	fmt.Println("\n --- TEST 4 (Fathers) ---")

	fmt.Println("BFS")
	tree := bfsMatrix(adjacency, 1)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("1. father of", v[0], "is", v[1])
		}
	}

	tree = bfsMatrix(adjacency, 2)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("2. father of", v[0], "is", v[1])
		}
	}

	tree = bfsMatrix(adjacency, 3)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("3. father of", v[0], "is", v[1])
		}
	}

	fmt.Println("DFS")
	tree = dfsMatrix(adjacency, 1)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("1. father of", v[0], "is", v[1])
		}
	}

	tree = dfsMatrix(adjacency, 2)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("2. father of", v[0], "is", v[1])
		}
	}

	tree = dfsMatrix(adjacency, 3)
	for _, v := range tree {
		if v[0] == 10 || v[0] == 20 || v[0] == 30 {
			fmt.Println("3. father of", v[0], "is", v[1])
		}
	}

	// test 5: distance between vertex (10, 20) (10, 30), (20, 30)
	fmt.Println("\n --- TEST 5 (Distances) ---")

	path := findPathMatrix(adjacency, 10, 20)
	fmt.Println("Distance between 10 and 20: ", len(path)-1)

	path = findPathMatrix(adjacency, 10, 30)
	fmt.Println("Distance between 10 and 30: ", len(path)-1)

	path = findPathMatrix(adjacency, 20, 30)
	fmt.Println("Distance between 20 and 30: ", len(path)-1)

	// test 6: Connected components
	fmt.Println("\n --- TEST 6 (Connected components) ---")

	components_time := time.Now()
	components := findComponentsMatrix(adjacency)

	fmt.Println("Time finding components:", time.Since(components_time))
	fmt.Println("Number of components:", len(components))

	// test 7: Diameter
	fmt.Println("\n --- TEST 7 (Diameter) ---")

	diameter_time := time.Now()
	diameter, _, _ := findDiameterMatrix(adjacency)

	fmt.Println("Time finding diameter (slow):", time.Since(diameter_time))
	fmt.Println("Diameter:", diameter)

	diameter_time = time.Now()
	diameter, _, _ = findDiameterQuickMatrix(adjacency)

	fmt.Println("Time finding diameter (fast):", time.Since(diameter_time))
	fmt.Println("Diameter:", diameter)

	fmt.Println("\n --- END TESTS ---")
}
