package heap_test

import (
	"awesome-util/go/heap"
	"testing"
)

type Data struct {
	value int
}

func MinHeap(a, b Data) bool { return a.value < b.value } // min-heap
func MaxHeap(a, b Data) bool { return a.value > b.value } // max-heap

func TestHeap(t *testing.T) {
	h := heap.NewHeap(MaxHeap)
	h.Push(Data{value: 2})
	h.Push(Data{value: 1})
	h.Push(Data{value: 3})
	h.Push(Data{value: 4})
	h.Push(Data{value: 5})
	h.Push(Data{value: 7})
	h.Push(Data{value: 6})

	for range 7 {
		println(h.Pop().value)
	}
}
