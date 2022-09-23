package main

// import "fmt"

/*
--- in:
adjaency: adjaency list
start: start vertex
--- out:
tree {{node, father, level}, ...}
*/
func bfsList(adjacency [][]uint32, start uint32) [][3]uint32 {
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
		for _, neighbor := range adjacency[current[0]] {
			if !visited[neighbor] {
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
adjaency: adjaency list
start: start vertex
--- out:
tree {{node, father, level}, ...}
*/
func dfsList(adjacency [][]uint32, start uint32) [][3]uint32 {
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
			for _, neighbor := range adjacency[current[0]] {
				if !visited[neighbor] {
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
adjaency: adjaency list
start: start vertex
goal: destination vertex
--- out:
path {end, ..., start}
*/
func findPathList(adjacency [][]uint32, root uint32, end uint32) []uint32 {
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
		for _, neighbor := range adjacency[current[0]] {
			if father[neighbor] == 0 {
				father[neighbor] = current[0]
				queue = append(queue, [2]uint32{neighbor, current[1] + 1})
			}
			if neighbor == end {
				found = true
				break
			}
		}
	}

	// if no path found return empty path
	if father[end] == 0 {
		return []uint32{}
	}

	path := []uint32{end}
	for path[len(path)-1] != root {
		path = append(path, father[path[len(path)-1]])
	}

	return path
}

/*
Calls the BFS for each vertex and return the depth of the highest
--- in:
adjacency: adjaency list
--- out:
diameter: the max min distance between any two vertices
vextex1, vertex2: the two vertices with the max min distance
*/
func findDiameterList(adjacency [][]uint32) (uint32, uint32, uint32) {
	var diameter, vertex1, vertex2 uint32
	for i := range adjacency {
		tree := bfsList(adjacency, uint32(i))
		lastVertex := tree[len(tree)-1]
		if lastVertex[2] > diameter { // level > max
			vertex1 = uint32(i)
			vertex2 = lastVertex[0]
			diameter = lastVertex[2]
		}
	}

	return diameter, vertex1, vertex2
}
