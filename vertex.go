package main

import "fmt"

/*
--- in:
adjaency: adjaency list
start: start vertex
--- out:
tree {{node, father, level}, ...}
*/
func bfsList(adjacency [][]uint32, start uint32) [][3]uint32 {
	fmt.Print("bfsList: ", start)
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

// func dfsListx(adjacency [][]uint32, start uint32) string {
// 	// ----- DFS -----
// 	visited := make([]bool, len(adjacency))
// 	visited[start] = true
// 	// stack := []uint32{start}
// 	level := uint32(0)
// 	stack := [][3]uint32{{start, start, level}}

// 	treeText := ""
// 	line := "node\tfather\tlevel\n"
// 	// line += strconv.Itoa(int(start)) + "\tnull\t0\n"

// 	for len(stack) > 0 {
// 		level += 1

// 		// pop()
// 		current := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		line = strconv.Itoa(int(current[0])) + "\t" + strconv.Itoa(int(current[1])) + "\t" + strconv.Itoa(int(current[2])) + "\n"

// 		// add neighbors to queue
// 		for _, neighbor := range adjacency[current[0]] {
// 			if !visited[neighbor] {
// 				visited[neighbor] = true
// 				stack = append(stack, [3]uint32{current[0], neighbor, level})

// 				treeText += line
// 			}
// 		}
// 	}

// 	return treeText
// }

// func bfsListx(adjacency [][]uint32, start uint32) string {
// 	// ----- BFS -----
// 	visited := make([]bool, len(adjacency))
// 	visited[start] = true
// 	queue := []uint32{start}

// 	treeText := ""
// 	line := "node\tfather\tlevel\n"
// 	line += strconv.Itoa(int(start)) + "\tnull\t0\n"

// 	level := 0
// 	for len(queue) > 0 {
// 		level += 1

// 		// pop()
// 		current := queue[0]
// 		queue = queue[1:]

// 		// add neighbors to queue
// 		for _, neighbor := range adjacency[current] {
// 			if !visited[neighbor] {
// 				visited[neighbor] = true
// 				queue = append(queue, neighbor)
// 				line := strconv.Itoa(int(neighbor)) + "\t" + strconv.Itoa(int(current)) + "\t" + strconv.Itoa(level) + "\n"
// 				treeText += line
// 			}
// 		}
// 	}

// 	return treeText
// }
