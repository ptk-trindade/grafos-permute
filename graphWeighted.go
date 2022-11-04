package main

import "fmt"

type Neighbor struct {
	vertex_id uint32
	weight    float64
}

type TreeNode struct {
	id     uint32
	father uint32
	cost   float64
}

/*
Find the shortest path from root to end in a given tree using Dijkstra algorithm
--- in:
adjacency: the adjacency list of the graph
root: the root of the tree
end: the end of the path (the goal vertex) / if 0 explores all the graph
--- out:
path: the path from root to end
*/
func dijkstraList(adjacency [][]*Neighbor, start uint32, end uint32) []*TreeNode {
	// initialize the list
	list := List{}
	list.vertexes = make([]*Node, len(adjacency)+1) // 0 is not used
	list.Update(start, 0, 0)

	// tree
	tree := make([]*TreeNode, 0, len(adjacency))

	node := list.Pop()
	for node != nil && node.id != end {
		// add node to tree
		tree = append(tree, &TreeNode{node.id, node.father, node.cost})

		// update neighbors
		for _, v := range adjacency[node.id] {
			list.Update(v.vertex_id, node.id, node.cost+v.weight)
		}
		node = list.Pop() // get the node with the lowest cost
	}

	if node != nil { // add the end node
		tree = append(tree, &TreeNode{node.id, node.father, node.cost})
	}

	return tree
}

func dijkstraHeap(adjacency [][]*Neighbor, start uint32, end uint32) []*TreeNode {
	// initialize the heap
	heap := Heap{}
	heap.vertexPos = make([]int, len(adjacency))
	for i := range heap.vertexPos { // set all vertexes as unexplored
		heap.vertexPos[i] = -1
	}
	heap.Update(start, 0, 0)

	// tree
	tree := make([]*TreeNode, 0, len(adjacency)+1) // 0 is not used

	node := heap.Pop()
	for node != nil && node.id != end {
		// add node to tree
		tree = append(tree, &TreeNode{node.id, node.father, node.cost})

		// update neighbors
		for _, v := range adjacency[node.id] {
			heap.Update(v.vertex_id, node.id, node.cost+v.weight)
		}

		node = heap.Pop() // get the node with the lowest cost
	}

	if node != nil { // add the end node
		tree = append(tree, &TreeNode{node.id, node.father, node.cost})
	}

	return tree
}

func primHeap(adjacency [][]*Neighbor) ([]*TreeNode, float64) {

	// initialize the heap
	heap := Heap{}
	heap.vertexPos = make([]int, len(adjacency))
	for i := range heap.vertexPos { // set all vertexes as unexplored
		heap.vertexPos[i] = -1
	}
	heap.Update(1, 0, 0) // start from vertex 1

	// tree
	tree := make([]*TreeNode, 0, len(adjacency))
	total_cost := 0.0
	node := heap.Pop()
	for node != nil {
		total_cost += node.cost

		// add node to tree
		tree = append(tree, &TreeNode{node.id, node.father, node.cost})

		// update neighbors
		for _, v := range adjacency[node.id] {
			heap.Update(v.vertex_id, node.id, v.weight)
		}

		node = heap.Pop() // get the node with the lowest cost
	}

	return tree, total_cost
}

/*
Find the path from root to end in a given tree
--- in:
tree: the tree
root: the root of the tree
end: the end of the path (the goal vertex)
--- out:
path: the path from root to end
cost: the cost of the path
*/
func findPath(tree []*TreeNode, root uint32, end uint32) ([]uint32, float64) {
	if end == 0 {
		fmt.Println("Unable to find path to 0")
		return []uint32{}, 0
	}

	i := len(tree) - 1
	for tree[i].id != end && i > 0 {
		i--
	}

	if i == 0 {
		fmt.Println("Unable to find path to ", end)
		return []uint32{}, 0
	}

	cost := tree[i].cost
	path := []uint32{end, tree[i].father}

	next_father := tree[i].father
	for ; next_father != root; i-- {
		if tree[i].id == next_father {
			path = append(path, tree[i].father)
			next_father = tree[i].father
		}
	}

	// reverse the path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, cost
}
