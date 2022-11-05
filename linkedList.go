package main

import "fmt"

type Node struct {
	id     uint32
	father uint32
	cost   float64
	prev   *Node
	next   *Node
}

type List struct {
	vertexes []*Node // pointer to the vertex node in list
	head     *Node
	tail     *Node
}

func (L *List) insertAfter(node *Node, new_node *Node) {
	if node == nil {
		if L.head == nil { // if the list is empty
			L.head = new_node
			L.tail = new_node
			return
		}

		new_node.next = L.head
		L.head.prev = new_node
		L.head = new_node
		return
	}

	if node.next != nil {
		node.next.prev = new_node
		new_node.next = node.next
	} else {
		L.tail = new_node
	}

	node.next = new_node
	new_node.prev = node

}

func (L *List) remove(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		L.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		L.tail = node.prev
	}

	node.prev = nil
	node.next = nil

}

func (L *List) reorder(node *Node) {
	if node.prev == nil || node.prev.cost <= node.cost {
		return
	}

	prev_node := node.prev
	L.remove(node)

	for prev_node != nil && prev_node.cost > node.cost {
		// finds the first node with a lower (or equal) cost and inserts the node after it
		prev_node = prev_node.prev
	}

	L.insertAfter(prev_node, node)

}

/*
Updates the cost to a vertex in the list (inserts if not present)
--- in:
vertex_id: the id of the vertex to update
father: the id of the father of the vertex
cost: cost from start to node
*/
func (L *List) Update(vertex_id uint32, father uint32, cost float64) {
	var updated_node *Node

	if L.vertexes[vertex_id] == nil {
		updated_node = &Node{
			id:     vertex_id,
			father: father,
			cost:   cost,
		}

		L.vertexes[vertex_id] = updated_node

		L.insertAfter(L.tail, updated_node)

	} else { // if the node is already in the list

		if L.vertexes[vertex_id].cost < cost { // if the new cost is higher than the old one
			return
		}

		updated_node = L.vertexes[vertex_id]
		updated_node.father = father
		updated_node.cost = cost
	}

	L.reorder(updated_node)
}

func (L *List) Pop() *Node {
	if L.head == nil {
		return nil
	}

	node := L.head
	L.remove(node)
	// L.head = L.head.next
	// node.next = nil
	return node
}

// --------
func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.id)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.id)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.id)
		list = list.prev
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}
