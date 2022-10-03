package tree

import (
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (t *Tree) GetRoot() *Node {
	return t.root
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = &Node{data: data}
	} else {
		t.root.Insert(data)
	}
}

func (n *Node) Insert(data int) {
	if data < n.data {
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data}
		} else {
			n.right.Insert(data)
		}
	}
}

func (t *Tree) InOrderParse(node *Node) {
	if node == nil {
		return
	}
	t.InOrderParse(node.left)
	fmt.Println(node.data)
	t.InOrderParse(node.right)
}

func (t *Tree) PreOrderParse(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	t.InOrderParse(node.left)
	t.InOrderParse(node.right)
}

func (t *Tree) PostOrderParse(node *Node) {
	if node == nil {
		return
	}
	t.InOrderParse(node.left)
	t.InOrderParse(node.right)
	fmt.Println(node.data)
}
