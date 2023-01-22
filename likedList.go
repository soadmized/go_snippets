package main

import "fmt"

type Node struct {
	next *Node
	val  string
}

type List struct {
	head *Node
	tail *Node
	len  int
}

// Push adds node to the tail of list
func (l *List) Push(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.len++
}

// Pop removes last node from the list and returns it
func (l *List) Pop() *Node {
	if l.head == nil {
		return nil
	} else {
		iter := l.head.next
		for iter != nil {
			curr := iter
			iter = curr.next

			if iter.next == nil {
				l.tail = curr
				l.tail.next = nil
				l.len--
				return iter
			}
		}
		return nil
	}
}

// Print the values of nodes
func (l *List) Print() {
	iter := l.head
	fmt.Print("\n")
	for iter != nil {
		fmt.Printf("%s", iter.val)
		iter = iter.next
		if iter != nil {
			fmt.Print(" -> ")
		}
	}
}

func main() {
	var list List
	first := Node{val: "first node"}
	second := Node{val: "second node"}
	third := Node{val: "third node"}
	fourth := Node{val: "fourth node"}
	list.Push(&first)
	list.Push(&second)
	list.Push(&third)
	list.Push(&fourth)
	list.Print()
	list.Pop()
	list.Print()
}
