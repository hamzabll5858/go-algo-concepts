package graph

import "fmt"

type Stack struct {
	nodes []*Node
}

func (n *Node) getData() *Node {
	return n
}

func (s *Stack) Push(data *Node) {
	s.nodes = append(s.nodes, data)
}

func (s *Stack) Pop() *Node {
	if len(s.nodes) > 1 {
		temp := s.nodes[len(s.nodes)-1]
		s.nodes = s.nodes[:len(s.nodes)-1]
		return temp
	} else if len(s.nodes) == 1 {
		temp := s.nodes[len(s.nodes)-1]
		s.nodes = []*Node{}
		return temp
	} else {
		return nil
	}
}

func (s *Stack) Top() *Node {
	if len(s.nodes) > 0 {
		return s.nodes[len(s.nodes)-1].getData()
	} else {
		return nil
	}
}

func (g *Graph) DFS() {
	g.lock.RLock()

	bucket := Stack{}
	n := g.nodes[0]
	bucket.Push(n)
	visited := make(map[*Node]bool)
	visited[n] = true
	for {
		if bucket.Top() == nil {
			break
		}
		newNode := bucket.Pop()
		fmt.Println(newNode.data)

		for _, value := range g.edges[*newNode] {
			if !visited[value] {
				bucket.Push(value)
				visited[value] = true
			}
		}
	}

	g.lock.RUnlock()
}
