package main

type HeapNode struct {
	id     uint32
	father uint32
	cost   float64
}

type Heap struct {
	vertexPos []int // -1 unexplored vertex, -2 explored vertex, >= 0 position in heap
	heap      []HeapNode
}

func (H *Heap) Update(vertex_id uint32, father uint32, cost float64) {
	tree_pos := H.vertexPos[vertex_id]
	if tree_pos == -1 { // new vertex
		H.insert(&HeapNode{vertex_id, father, cost})
	} else if tree_pos == -2 { // explored vertex
		return
	} else { // vertex in heap
		if H.heap[tree_pos].cost > cost {
			H.heap[tree_pos].father = father
			H.heap[tree_pos].cost = cost
			H.bubbleUp(tree_pos)
		}
	}
}

func (H *Heap) insert(node *HeapNode) {
	H.heap = append(H.heap, *node)
	H.vertexPos[node.id] = len(H.heap) - 1
	H.bubbleUp(len(H.heap) - 1)
}

func (H *Heap) Pop() *HeapNode {
	if len(H.heap) == 0 {
		return nil
	}

	node := H.heap[0]
	H.swap(0, len(H.heap)-1)
	H.heap = H.heap[:len(H.heap)-1]
	H.vertexPos[node.id] = -2
	H.bubbleDown()

	return &node
}

func (H *Heap) swap(parent_pos int, child_pos int) {
	parent_id := H.heap[parent_pos].id
	child_id := H.heap[child_pos].id

	H.heap[parent_pos], H.heap[child_pos] = H.heap[child_pos], H.heap[parent_pos]
	H.vertexPos[parent_id], H.vertexPos[child_id] = H.vertexPos[child_id], H.vertexPos[parent_id]
}

func (H *Heap) bubbleUp(index int) {
	if index == 0 {
		return
	}

	parent_pos := (index - 1) / 2

	for index > 0 && H.heap[parent_pos].cost > H.heap[index].cost {
		H.swap(index, parent_pos)
		index = parent_pos
		parent_pos = (index - 1) / 2
	}
}

func (H *Heap) bubbleDown() {
	index := 0

	swaped := true
	for swaped {
		swaped = false
		left_pos := 2*index + 1
		right_pos := 2*index + 2
		if left_pos >= len(H.heap) { // no children
			return
		}

		if right_pos >= len(H.heap) { // only left child
			if H.heap[index].cost > H.heap[left_pos].cost {
				H.swap(index, left_pos)
				swaped = true
				index = left_pos
			}
		} else { // both children
			minor_pos := left_pos
			if H.heap[left_pos].cost > H.heap[right_pos].cost {
				minor_pos = right_pos
			}

			if H.heap[index].cost > H.heap[minor_pos].cost {
				H.swap(index, minor_pos)
				swaped = true
				index = minor_pos
			}
		}
	}
}
