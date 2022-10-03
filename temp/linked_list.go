package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	root *Node
}

func (l *LinkedList) Insert(data int) {
	if l.root == nil {
		l.root = &Node{data: data, next: nil, prev: nil}
	} else {
		newNode := &Node{data: data, next: nil, prev: nil}
		curr := l.root
		for curr.next != nil {
			curr = curr.next
		}
		newNode.prev = curr
		curr.next = newNode
	}
}

func (l *LinkedList) Remove(data int) {
	if l.root != nil {
		if l.root.data == data {
			l.root = l.root.next
		} else {
			curr := l.root
			var prev *Node
			for curr != nil {
				if curr.data == data {
					prev.next = curr.next
					if curr.next != nil {
						curr.next.prev = prev
					}
				}
				prev = curr
				curr = curr.next
			}
		}
	}
}

func (l *LinkedList) ReverseList() {
	curr := l.root
	var next *Node
	var temp *Node
	for curr != nil {
		temp = curr
		next = curr.next
		curr.next, curr.prev = curr.prev, curr.next
		curr = next
	}
	l.root = temp

}

func (l *LinkedList) PrintList() {
	curr := l.root
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
}

func main() {

	var list LinkedList

	for i := 0; i < 4; i++ {
		list.Insert(i)
	}
	list.PrintList()
	list.ReverseList()
	list.PrintList()

}
