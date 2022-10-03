package graph

import (
	"fmt"
	"sync"
)

type NodeQueue struct {
	nodes []*Node
	lock  sync.RWMutex
}

func (q *NodeQueue) New() *NodeQueue {
	q.lock.Lock()
	q.nodes = []*Node{}
	q.lock.Unlock()
	return q
}

func (q *NodeQueue) EnQueue(n *Node) {
	q.lock.Lock()
	q.nodes = append(q.nodes, n)
	q.lock.Unlock()
}

func (q *NodeQueue) DeQueue() *Node {
	q.lock.Lock()
	defer q.lock.Unlock()
	result := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]
	return result
}

func (q *NodeQueue) Front() *Node {
	q.lock.RLock()
	defer q.lock.RUnlock()
	result := q.nodes[0]
	return result

}

func (q *NodeQueue) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if len(q.nodes) <= 0 {
		return true
	} else {
		return false
	}
}

func (q *NodeQueue) Size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.nodes)
}

func (g *Graph) BFS() {
	g.lock.RLock()

	que := NodeQueue{}
	que.New()
	n := g.nodes[0]
	que.EnQueue(n)
	fmt.Println(n.data)
	visited := make(map[*Node]bool)
	for {
		if que.IsEmpty() {
			break
		}
		newNode := que.DeQueue()
		visited[newNode] = true

		for _, value := range g.edges[*newNode] {
			if !visited[value] {
				visited[value] = true
				que.EnQueue(value)
				fmt.Println(value.data)
			}
		}
	}

	g.lock.RUnlock()
}
