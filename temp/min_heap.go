package main

import "fmt"

type MinHeap struct {
	nodes   []int
	size    int
	maxSize int
}

func (h *MinHeap) IsLeaf(index int) bool {
	if index >= (h.size/2) && index < h.size {
		return true
	}
	return false
}

func (h *MinHeap) ParentIndex(index int) int {
	return (index - 1) / 2
}

func (h *MinHeap) LeftChildIndex(index int) int {
	return (index * 2) + 1
}

func (h *MinHeap) RightChildIndex(index int) int {
	return (index * 2) + 2
}

func (h *MinHeap) Swap(leftIndex int, rightIndex int) {
	h.nodes[leftIndex], h.nodes[rightIndex] = h.nodes[rightIndex], h.nodes[leftIndex]
}

func (h *MinHeap) InsertNode(data int) error {
	if h.size >= h.maxSize {
		return fmt.Errorf("heap is full")
	} else {
		h.nodes = append(h.nodes, data)
		h.size++
		h.UpHeapify(h.size - 1)
		return nil
	}
}

func (h *MinHeap) UpHeapify(index int) {
	for h.nodes[index] < h.nodes[h.ParentIndex(index)] {
		h.Swap(index, h.ParentIndex(index))
	}
}

func (h *MinHeap) DownHeapify(current int) {
	if h.IsLeaf(current) {
		return
	}
	smallest := current
	leftIndex := h.LeftChildIndex(current)
	rightIndex := h.RightChildIndex(current)

	if leftIndex < h.size && h.nodes[leftIndex] < h.nodes[smallest] {
		smallest = leftIndex
	}
	if rightIndex < h.size && h.nodes[rightIndex] < h.nodes[smallest] {
		smallest = rightIndex
	}

	if smallest != current {
		h.Swap(current, smallest)
		h.DownHeapify(smallest)
	}
	return
}

func (h *MinHeap) BuildHeap() {
	for i := (h.size / 2) - 1; i >= 0; i-- {
		h.DownHeapify(i)
	}
}

func (h *MinHeap) Remove() int {
	result := h.nodes[0]
	h.nodes[0] = h.nodes[h.size-1]
	h.nodes = h.nodes[:h.size-1]
	h.size--
	h.DownHeapify(0)
	return result
}

func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	minHeap := &MinHeap{nodes: []int{}, size: 0, maxSize: 50}
	for i := 0; i < len(inputArray); i++ {
		minHeap.InsertNode(inputArray[i])
	}
	minHeap.BuildHeap()
	for i := 0; i < len(inputArray); i++ {
		fmt.Println(minHeap.Remove())
	}
}
