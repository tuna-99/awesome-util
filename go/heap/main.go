package tutorial

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) First() int {
	return (*h)[0]
}

func totalCost(costs []int, k int, candidates int) int64 {
	first := &IntHeap{}
	last := &IntHeap{}
	heap.Init(first)
	heap.Init(last)
	n := len(costs)
	for i := 0; i < candidates; i++ {
		heap.Push(first, costs[i])
		heap.Push(last, costs[n-i-1])
	}

	var total int64
	left := candidates
	right := n - 1 - candidates
	for i := 0; i < k; i++ {
		if first.Len() == 0 && last.Len() == 0 {
			break
		}

		if n < 2*candidates {
			cost, _ := (heap.Pop(first)).(int64)
			total += cost
			continue
		}

		if first.Len() == 0 {
			cost, _ := (heap.Pop(last)).(int64)
			total += cost
			continue
		}

		if last.Len() == 0 {
			cost, _ := (heap.Pop(first)).(int64)
			total += cost
			continue
		}

		if first.First() <= last.First() {
			cost, _ := (heap.Pop(first)).(int64)
			total += cost
			if left < right {
				heap.Push(first, left)
				left++
			}
		} else {
			cost, _ := (heap.Pop(last)).(int64)
			total += cost
			if left < right {
				heap.Push(last, right)
				right--
			}
		}
	}

	return total
}
