// Package topn is an experimental pkg for computing the top N items
// from a list.
package topn

import (
	"container/heap"
	"sort"
)

// MaxInts returns the max n ints from x.
func MaxInts(x []int, n int) []int {
	if len(x) <= n {
		sort.Ints(x) // changes x
		return x
	}

	min := make(minIntHeap, n)
	for i := 0; i < n; i++ {
		min[i] = x[i]
	}

	heap.Init(&min)
	for _, z := range x[n:] {
		if min[0] < z {
			min[0] = z
			heap.Fix(&min, 0)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(min)))
	return min
}

// An minIntHeap is a min-heap of ints.
type minIntHeap []int

func (h minIntHeap) Len() int           { return len(h) }
func (h minIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minIntHeap) Push(x interface{}) {
	panic("should not use")
}

func (h *minIntHeap) Pop() interface{} {
	panic("should hot use")
}
