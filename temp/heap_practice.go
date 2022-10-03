package main

type heap struct {
	arr []int
}

func (h *heap) GetLeftChildIndex(index int) int {
	return index*2 + 1
}

func (h *heap) GetRightChildIndex(index int) int {
	return index*2 + 2
}

func (h *heap) GetParentIndex(index int) int {
	return index - 1/2
}

func (h *heap) IsLeaf(index int) bool {
	if index >= len(h.arr)/2 {
		return true
	} else {
		return false
	}
}

func (h *heap) InsertNode(num int) {
	h.arr = append(h.arr, num)
	h.UpHeapify(len(h.arr) - 1)
}

func (h *heap) UpHeapify(index int) {
	if h.arr[index] > h.arr[h.GetParentIndex(index)] {
		h.arr[index], h.arr[h.GetParentIndex(index)] = h.arr[h.GetParentIndex(index)], h.arr[index]
		h.UpHeapify(h.GetParentIndex(index))
	}
}

func (h *heap) Remove(index int) int {
	result := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.DownHeapify(0)
	return result
}

func (h *heap) DownHeapify(index int) {
	if !h.IsLeaf(index) {
		smallest := index
		leftChild := h.GetLeftChildIndex(index)
		rightChild := h.GetRightChildIndex(index)

		if h.arr[smallest] > h.arr[rightChild] {
			smallest = rightChild
		}
		if h.arr[smallest] > h.arr[leftChild] {
			smallest = leftChild
		}

		if smallest != index {
			h.arr[smallest], h.arr[index] = h.arr[index], h.arr[smallest]
			h.DownHeapify(smallest)
		}
	}
}

func (h *heap) BuildHeap() {
	for i := 0; i < len(h.arr); i++ {
		h.UpHeapify(i)
	}
}

func main() {

	h := new(heap)
	for i := 0; i < 10; i++ {
		h.InsertNode(i)
	}
	h.BuildHeap()

	h.Remove(5)
}
