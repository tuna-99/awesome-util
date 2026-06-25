package heap

type Heap[T any] struct {
	data []T

	priority func(a, b T) bool
}

// [1, 2, 3, 4, 5, 6, 7]
//
//			1
//		2       3
//	  4   5   6   7
//
// root of the tree is at index 0
// children of index i are at 2i + 1 and 2i + 2
// parent of index i is at (i - 1) / 2
func NewHeap[T any](priority func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0),

		priority: priority,
	}
}

func (h *Heap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.shiftUp(len(h.data) - 1) // bubble up
}

func (h *Heap[T]) Pop() T {
	data := h.data[0]
	lastIdx := len(h.data) - 1
	h.data[0] = h.data[lastIdx] // copy last to root
	h.data = h.data[:lastIdx]   // remove last
	h.shiftDown(0)              // sink down root

	return data
}

func (h *Heap[T]) shiftUp(idx int) {
	if idx <= 0 {
		return
	}

	parentIdx := (idx - 1) / 2
	if h.priority(h.data[parentIdx], h.data[idx]) {
		return // already at correct position
	}

	h.data[idx], h.data[parentIdx] = h.data[parentIdx], h.data[idx]
	h.shiftUp(parentIdx)
}

func (h *Heap[T]) shiftDown(idx int) {
	leftIdx := 2*idx + 1
	rightIdx := 2*idx + 2

	priorityIdx := idx
	if leftIdx < len(h.data) && h.priority(h.data[leftIdx], h.data[priorityIdx]) {
		priorityIdx = leftIdx
	}
	if rightIdx < len(h.data) && h.priority(h.data[rightIdx], h.data[priorityIdx]) {
		priorityIdx = rightIdx
	}

	if priorityIdx == idx {
		return // already at correct position
	}

	h.data[idx], h.data[priorityIdx] = h.data[priorityIdx], h.data[idx]
	h.shiftDown(priorityIdx)
}
