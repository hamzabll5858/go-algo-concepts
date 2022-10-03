package main

import "fmt"

type NodeTree struct {
	data  int
	left  *NodeTree
	right *NodeTree
}

type Tree struct {
	root *NodeTree
}

func (t *Tree) InsertTree(data int) {
	if t.root == nil {
		t.root = &NodeTree{data: data}
	} else {
		t.Insert(t.root, data)
	}
}

func (t *Tree) Insert(root *NodeTree, data int) *NodeTree {
	if root == nil {
		root = &NodeTree{data: data}
		return root
	} else {
		if data > root.data {
			root.right = t.Insert(root.right, data)
		} else {
			root.left = t.Insert(root.left, data)
		}
	}
	return root
}

func (t *Tree) PostOrderTraversal(node *NodeTree) {
	if node == nil {
		return
	} else {
		t.PostOrderTraversal(node.left)
		t.PostOrderTraversal(node.right)
		fmt.Println(node.data)
	}

}

func main() {
	t := new(Tree)

	t.InsertTree(9)
	t.InsertTree(4)
	t.InsertTree(1)
	t.InsertTree(3)
	t.InsertTree(11)

	t.PostOrderTraversal(t.root)

}
