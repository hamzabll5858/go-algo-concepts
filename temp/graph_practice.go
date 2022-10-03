package main

type NodeGraph struct {
	data int
	adj  []*NodeGraph
}

type Graph2 struct {
	root []*NodeGraph
}

func (g *Graph2) AddNode(data int) *NodeGraph {
	node := &NodeGraph{data: data}
	g.root = append(g.root, node)
	return node
}

func (g *Graph2) AddEdge(n1 *NodeGraph, n2 *NodeGraph) (*NodeGraph, *NodeGraph) {
	n1Exists, n2Exists := false, false
	for i := 0; i < len(g.root); i++ {
		if g.root[i] == n1 {
			g.root[i].adj = append(g.root[i].adj, n1)
			n1Exists = true
		} else if g.root[i] == n2 {
			g.root[i].adj = append(g.root[i].adj, n2)
			n2Exists = true
		}
	}

	if !n1Exists {
		g.root = append(g.root, n1)
		g.root = append(g.root[len(g.root)-1].adj, n2)
	}
	if !n2Exists {
		g.root = append(g.root, n2)
		g.root = append(g.root[len(g.root)-1].adj, n1)
	}
	return n1, n2
}

func main() {

	g := new(Graph2)

	n4 := g.AddNode(4)
	n3 := g.AddNode(3)
	n2 := g.AddNode(2)
	n1 := g.AddNode(1)

	g.AddEdge(n4, n3)
	g.AddEdge(n3, n2)
	g.AddEdge(n2, n1)
	g.AddEdge(n1, n4)

}
