package main

import "sort"

/*
--- in:
adjacency: adjacency matrix
start: start vertex
--- out:
tree {{node, father, level}, ...}
*/
func bfsMatrix(adjacency [][]uint8, start uint32) [][3]uint32 {
	// ----- BFS -----
	visited := make([]bool, len(adjacency))
	visited[start] = true
	queue := [][]uint32{{start, 0}} // {node, level}

	tree := [][3]uint32{{start, 0, 0}}

	for len(queue) > 0 {
		// pop()
		current := queue[0]
		queue = queue[1:]

		// add neighbors to queue
		for neighbor := uint32(1); neighbor < uint32(len(adjacency[current[0]])); neighbor++ { // skip null vertex (0)
			if adjacency[current[0]][neighbor] == 1 && !visited[neighbor] { // if neighbor and not visited
				visited[neighbor] = true
				queue = append(queue, []uint32{neighbor, current[1] + 1})
				tree = append(tree, [3]uint32{neighbor, current[0], current[1] + 1})
			}
		}
	}

	return tree
}

/*
--- in:
adjacency: adjacency matrix
start: start vertex
--- out:
tree {{node, father, level}, ...}
*/
func dfsMatrix(adjacency [][]uint8, start uint32) [][3]uint32 {
	// ----- DFS -----
	visited := make([]bool, len(adjacency))

	stack := [][3]uint32{{start, 0, 0}} // {node, father, level}

	tree := make([][3]uint32, 0, len(adjacency))

	for len(stack) > 0 {
		// pop(-1)
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[current[0]] {
			visited[current[0]] = true

			tree = append(tree, current)

			// add neighbors to queue
			for neighbor := uint32(1); neighbor < uint32(len(adjacency[current[0]])); neighbor++ { // skip null vertex (0)
				if adjacency[current[0]][neighbor] == 1 && !visited[neighbor] { // if neighbor and not visited
					stack = append(stack, [3]uint32{neighbor, current[0], current[2] + 1})
				}
			}
		}
	}

	return tree
}

/*
Finds the shortest path between two vertices in a graph. (BFS)
--- in:
adjacency: adjacency Matrix
root: start vertex
end: destination vertex
--- out:
path {end, ..., root}
*/
func findPathMatrix(adjacency [][]uint8, root uint32, end uint32) []uint32 {
	// ----- BFS -----
	father := make([]uint32, len(adjacency))
	queue := [][2]uint32{{root, 0}} // {node, level}
	father[root] = root

	found := false
	// create tree
	for !found && len(queue) > 0 {
		// pop()
		current := queue[0]
		queue = queue[1:]

		// add neighbors to queue
		for neighbor := uint32(1); neighbor < uint32(len(adjacency[current[0]])); neighbor++ { // skip null vertex (0)
			if adjacency[current[0]][neighbor] == 1 {
				if father[neighbor] == 0 { // if not visited (father not set)
					father[neighbor] = current[0]
					queue = append(queue, [2]uint32{neighbor, current[1] + 1})
				}
				if neighbor == end {
					found = true
					break
				}
			}
		}
	}

	// if no path found return empty path
	if father[end] == 0 {
		return []uint32{}
	}

	path := []uint32{end}
	for path[len(path)-1] != root { // runs through the tree until it reaches the root
		path = append(path, father[path[len(path)-1]])
	}

	return path
}

/*
Calls the BFS for each vertex and return the depth of the highest
--- in:
adjacency: adjacency matrix
--- out:
diameter: the max min distance between any two vertices
vextex1, vertex2: the two vertices with the max min distance
*/
func findDiameterMatrix(adjacency [][]uint8) (uint32, uint32, uint32) {
	var diameter, vertex1, vertex2 uint32
	for i := uint32(1); i < uint32(len(adjacency)); i++ {
		tree := bfsMatrix(adjacency, i)
		lastVertex := tree[len(tree)-1]
		if lastVertex[2] > diameter { // tree_depth > current_diameter
			vertex1 = uint32(i)
			vertex2 = lastVertex[0]
			diameter = lastVertex[2]
		}
	}

	return diameter, vertex1, vertex2
}

/*
Calls the BFS for one vertex and consider the most distant vertex to be one vertice of the diameter
--- in:
adjacency: adjacency matrix
component_vertex: one vertex of each component (if empty, the function finds them)
--- out:
diameter: the max min distance between any two vertices
vextex1, vertex2: the two vertices with the max min distance
*/
func findDiameterQuickMatrix(adjacency [][]uint8, component_vertex []uint32) (uint32, uint32, uint32) {

	if len(component_vertex) == 0 { // if no components, find them
		components := findComponentsMatrix(adjacency)

		for _, component := range components { // get one vertex of each component
			component_vertex = append(component_vertex, component[0])
		}
	}

	var diameter, vertex1, vertex2 uint32
	for _, vertex := range component_vertex { // for each component
		tree := bfsMatrix(adjacency, vertex) // run bfs from any vertex
		v1 := tree[len(tree)-1][0]           // get the most distant vertex

		tree = bfsMatrix(adjacency, v1) // run bfs from the most distant vertex
		lastVertex := tree[len(tree)-1]
		v2 := lastVertex[0]
		distance := lastVertex[2]

		if distance > diameter { // found a new diameter
			diameter = distance
			vertex1 = v1
			vertex2 = v2
		}
	}

	return diameter, vertex1, vertex2
}

/*
Goes through the vertexes and runs the BFS for each unvisited vertex
--- in:
adjacency: adjacency matrix
--- out:
components: {{vertex1, vertex2, ...}, component2, ...}
*/
func findComponentsMatrix(adjacency [][]uint8) [][]uint32 {
	visited := make([]bool, len(adjacency))
	components := make([][]uint32, 0)
	for i := 1; i < len(adjacency); i++ { // for each vertex
		if !visited[i] { // if not visited, belongs to a new component
			tree := bfsMatrix(adjacency, uint32(i))
			components[len(components)] = make([]uint32, 0, len(tree))

			// add all vertices of the component to the visited vector
			for _, vertex := range tree {
				visited[vertex[0]] = true
				components[len(components)-1] = append(components[len(components)-1], vertex[0])
			}
		}
	}

	sort.Slice(components, func(i, j int) bool { return len(components[i]) < len(components[j]) }) // ascending order

	return components
}
