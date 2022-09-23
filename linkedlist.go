package main

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	value uint32
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Push(value uint32) *Node {
	node := &Node{
		next:  L.head,
		value: value,
	}
	if L.head != nil {
		L.head.prev = node
	}
	L.head = node

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l

	return node
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.value)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.value)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.value)
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

func (L *List) Remove(n *Node) {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		L.head = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		L.tail = n.prev
	}
}
