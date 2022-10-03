package graph

import (
	"fmt"
	"sync"
)

type Node struct {
	data int
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()

	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()

	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	g.lock.Unlock()
}

func (g *Graph) PrintAdjList() {
	for key, arr := range g.edges {
		fmt.Print(key.data, ": ")
		for _, value := range arr {
			fmt.Print(value.data, " -> ")
		}
		fmt.Println("NULL")
	}
}

func FillGraph() Graph {
	var graph Graph
	n1 := Node{1}
	n2 := Node{2}
	n3 := Node{3}
	n4 := Node{4}

	graph.AddNode(&n1)
	graph.AddNode(&n2)
	graph.AddNode(&n3)
	graph.AddNode(&n4)

	graph.AddEdge(&n1, &n2)
	graph.AddEdge(&n2, &n4)
	graph.AddEdge(&n3, &n4)
	graph.AddEdge(&n3, &n1)

	return graph
}
