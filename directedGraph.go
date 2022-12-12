package main

type Edge struct {
	origin uint32
	dest   uint32
	weight float64
	comp   *Edge // the complementary edge
}

// the flow of an edge
type EdgeFlow struct {
	edge *Edge
	flow float64
}

/*
Finds the shortest path between two vertices in a graph. (BFS)
--- in:
adjacency: adjacency list
root: start vertex
end: destination vertex
--- out:
path {lastEdge, ..., firstEdge}
maxFlow
*/
func findPathResBFS(adjacency [][]*Edge, root uint32, end uint32) ([]*Edge, float64) {
	// ----- BFS -----
	father := make([]*Edge, len(adjacency))
	queue := []uint32{root}
	father[root] = nil

	found := false
	// create tree
	for !found && len(queue) > 0 {

		// pop()
		current := queue[0]
		queue = queue[1:]

		// add neighbors to queue
		for _, edge := range adjacency[current] {
			if father[edge.dest] == nil && edge.weight > 0 { // if not visited (father not set)
				father[edge.dest] = edge
				queue = append(queue, edge.dest)

				if edge.dest == end {
					found = true
					break
				}
			}
		}
	}

	// if no path found return nil
	if father[end] == nil {
		return nil, 0
	}

	path := []*Edge{father[end]}
	maxFlow := father[end].weight

	// var path = make([]*Edge) // []*Edge{father[end]}

	for path[len(path)-1].origin != root {
		path = append(path, father[path[len(path)-1].origin])
		if path[len(path)-1].weight < maxFlow {
			maxFlow = path[len(path)-1].weight
		}
	}

	return path, maxFlow
}

/*
Finds the maximum flow in a graph
--- in:
adjacency: adjacency list
source: source vertex (s)
sink: sink vertex (t)
--- out:
maxFlow
flowPerEdge {&edge, flow}: Flow per edge
*/
func fordFulkerson(adjacency [][]*Edge, source uint32, sink uint32) (float64, []EdgeFlow) {
	var maxFlow float64 = 0.0
	var path []*Edge
	var pathCapacity float64

	correlation := make([][2]*Edge, 0) // original edge, residual edge

	// -- create the residual graph [vextex_id][edges]
	// duplicate adjacency list
	residual := make([][]*Edge, len(adjacency))
	for i := range adjacency {
		// residual[i] = make([]*Edge, len(adjacency[i]))
		for j := range adjacency[i] {
			residual[i] = append(residual[i], &Edge{origin: adjacency[i][j].origin, dest: adjacency[i][j].dest, weight: adjacency[i][j].weight})
			correlation = append(correlation, [2]*Edge{adjacency[i][j], residual[i][j]})
		}
	}

	// create complementary edges
	for i := range residual {
		for j := range residual[i] {
			edge := residual[i][j]
			comp := &Edge{edge.dest, edge.origin, 0, edge}
			edge.comp = comp
			residual[edge.dest] = append(residual[edge.dest], comp)
		}
	}

	// while there is a path from source to sink
	path, pathCapacity = findPathResBFS(residual, source, sink)
	for path != nil {
		maxFlow += pathCapacity

		// update residual graph
		for i := 0; i < len(path); i++ {
			// log.Println("update edge:\t", *path[i])
			path[i].weight -= pathCapacity
			path[i].comp.weight += pathCapacity
			// log.Println("updated edge:\t", *path[i])
		}

		path, pathCapacity = findPathResBFS(residual, source, sink)
	}

	flowPerEdge := findFlowPerEdge(correlation)

	return maxFlow, flowPerEdge
}

/*
Finds the flow of each edge in the graph
By the difference between the original edge and the residual edge
--- in:
correlation: [original edge, residual edge]
--- out:
edgesFlow: flow of each edge
*/
func findFlowPerEdge(correlation [][2]*Edge) []EdgeFlow {
	edgeFlow := make([]EdgeFlow, 0)
	for _, edges := range correlation {
		edgeFlow = append(edgeFlow, EdgeFlow{edges[0], edges[0].weight - edges[1].weight})
	}

	return edgeFlow
}

/*
Finds the maximum flow in a graph (using 2 BFS)
--- in:
adjacency: adjacency list
source: source vertex (s)
sink: sink vertex (t)
--- out:
maxFlow
flowPerEdge {&edge, flow}: Flow per edge
*/
func fordFulkersonV2(adjacency [][]*Edge, source uint32, sink uint32) (float64, []EdgeFlow) {
	var maxFlow float64 = 0.0
	correlation := make([][2]*Edge, 0) // original edge, residual edge

	// -- create the residual graph [vextex_id][edges]
	// duplicate adjacency list
	residual := make([][]*Edge, len(adjacency))
	for i := range adjacency {
		// residual[i] = make([]*Edge, len(adjacency[i]))
		for j := range adjacency[i] {
			edge := &Edge{origin: adjacency[i][j].origin, dest: adjacency[i][j].dest, weight: adjacency[i][j].weight}
			residual[i] = append(residual[i], edge)
			correlation = append(correlation, [2]*Edge{adjacency[i][j], edge})

			// create complementary edges
			comp := &Edge{edge.dest, edge.origin, 0, edge}
			edge.comp = comp
			residual[edge.dest] = append(residual[edge.dest], comp)
		}
	}

	type Father struct {
		edge   *Edge
		treeID uint64
	}

	father := make([]Father, len(residual))
	sourceQueue := []uint32{source}
	sourceTreeID := uint64(1)
	sinkQueue := []uint32{sink}
	sinkTreeID := uint64(2)
	// father[source] = Father{nil, sourceTreeID}
	sourceLoop := &Edge{source, source, 1.7e+300, nil}
	sourceComp := &Edge{source, source, 1.7e+300, sourceLoop}
	sourceLoop.comp = sourceComp
	father[source] = Father{sourceLoop, sourceTreeID}

	// father[sink] = Father{nil, sinkTreeID}
	sinkLoop := &Edge{sink, sink, 1.7e+300, nil}
	sinkComp := &Edge{sink, sink, 1.7e+300, sinkLoop}
	sinkLoop.comp = sinkComp
	father[sink] = Father{sinkLoop, sinkTreeID}
	for { // while there is a path from source to sink
		//----- findPathWith2BFS -----

		var upEdge *Edge
		var downEdge *Edge
		var current uint32
		found := false
		// create tree
		for !found && len(sourceQueue) > 0 && len(sinkQueue) > 0 {

			// pop() pt1
			current = sourceQueue[0]
			sourceQueue = sourceQueue[1:]

			// add neighbors to queue
			for _, edge := range residual[current] {
				// log.Println("1 node edge > ", *edge)
				if edge.weight > 0 && father[edge.dest].treeID != sourceTreeID { // edge not saturated and not visited (by this BFS)
					if father[edge.dest].treeID == sinkTreeID { // found path
						upEdge = edge
						downEdge = father[edge.dest].edge
						found = true
						sourceQueue = append(sourceQueue, current) // not done with this node
						// log.Println("found path in 1")
						break
					} else {
						father[edge.dest] = Father{edge, sourceTreeID}
						sourceQueue = append(sourceQueue, edge.dest)
					}
				}
			}

			if !found { // sink tree
				// pop() pt1
				current = sinkQueue[0]
				sinkQueue = sinkQueue[1:]

				// add neighbors to queue
				for _, edge_comp := range residual[current] {
					edge := edge_comp.comp
					// log.Println("2 edge_comp > ", *edge_comp)
					// log.Println("2 node edge > ", *edge)
					if edge.weight > 0 && father[edge.origin].treeID != sinkTreeID { // edge not saturated and not visited (by this BFS)
						// log.Println("2 father compare", father[edge.origin].treeID, sourceTreeID)
						if father[edge.origin].treeID == sourceTreeID { // found path
							upEdge = father[edge.origin].edge
							downEdge = edge
							found = true
							// log.Println("found path in 2")
							sinkQueue = append(sinkQueue, current) // not done with this node
							break
						} else {
							father[edge.origin] = Father{edge, sinkTreeID}
							// log.Println("sinkQueue, append: ", *edge)
							sinkQueue = append(sinkQueue, edge.origin)
						}
					}
				}
			}
			// log.Println("redo? ", found, len(sourceQueue), len(sinkQueue))
		}

		if !found {
			// log.Println("no path found")
			// log.Println("found:", found, "sourceQueue:", len(sourceQueue), "sinkQueue:", len(sinkQueue))
			break // no path found
		}

		path := make([]*Edge, 0)
		// path in source direction
		path = append(path, upEdge)
		sourcePathCapacity := path[0].weight
		for path[len(path)-1].origin != source {
			path = append(path, father[path[len(path)-1].origin].edge)
			if path[len(path)-1].weight < sourcePathCapacity {
				sourcePathCapacity = path[len(path)-1].weight
			}
		}

		// path in sink direction
		path = append(path, downEdge)
		sinkPathCapacity := downEdge.weight
		for path[len(path)-1].dest != sink {
			path = append(path, father[path[len(path)-1].dest].edge)
			if path[len(path)-1].weight < sinkPathCapacity {
				sinkPathCapacity = path[len(path)-1].weight
			}
		}

		var pathCapacity float64
		if sourcePathCapacity < sinkPathCapacity {
			pathCapacity = sourcePathCapacity
			sourceTreeID += 2
			sourceQueue = []uint32{source}
			father[source] = Father{sourceLoop, sourceTreeID}

		} else if sourcePathCapacity > sinkPathCapacity {
			pathCapacity = sinkPathCapacity
			sinkTreeID += 2
			sinkQueue = []uint32{sink}
			father[sink] = Father{sinkLoop, sinkTreeID}

		} else { // sourcePathCapacity == sinkPathCapacity
			pathCapacity = sourcePathCapacity
			sourceTreeID += 2
			sinkTreeID += 2
			sourceQueue = []uint32{source}
			sinkQueue = []uint32{sink}
			father[source] = Father{sourceLoop, sourceTreeID}
			father[sink] = Father{sinkLoop, sinkTreeID}

		}

		maxFlow += pathCapacity

		// update residual graph
		for i := 0; i < len(path); i++ {
			// log.Println("update edge:\t", *path[i])
			path[i].weight -= pathCapacity
			path[i].comp.weight += pathCapacity
			// log.Println("updated edge:\t", *path[i])
		}

	}

	flowPerEdge := findFlowPerEdge(correlation)

	return maxFlow, flowPerEdge
}
