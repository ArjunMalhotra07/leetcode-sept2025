package practice_graphs

import "fmt"

// if graph is from 0 -> 4; n = 5
func CheckForHeap(graph [][]int, n int) {
	heap := &MinHeap{}
	// Weight, Node, Parent
	heap.add([]int{0, 0, -1})
	visited := make([]bool, n)
	ans := [][]int{}
	dist := 0

	for len(heap.Nums) != 0 {
		root := heap.extractRoot()
		wt := root[0]
		node := root[1]
		parent := root[2]
		if !visited[node] && parent != -1 {
			ans = append(ans, root)
			dist += wt
		}
		if !visited[node] {
			visited[node] = true
			for i, weight := range graph[node] {
				if weight > 0 && !visited[i] {
					heap.add([]int{weight, i, node})
				}
			}
		}
	}
	for _, e := range ans {
		fmt.Printf("%d -- %d (w=%d)\n", e[2], e[1], e[0])
	}
	fmt.Println("Total weight:", dist)

}

type MinHeap struct {
	Nums [][]int
}

func (h *MinHeap) add(n []int) {
	h.Nums = append(h.Nums, n)
	h.heapifyUp(len(h.Nums) - 1)
}
func (h *MinHeap) heapifyUp(i int) {
	for i > 0 && h.Nums[i][0] < h.Nums[parent(i)][0] {
		h.swap(i, parent(i))
		i = parent(i)
	}
}
func (h *MinHeap) extractRoot() []int {
	n := len(h.Nums)
	if n == 0 {
		return []int{-1, -1, -1}
	}
	root := h.Nums[0]
	h.Nums[0] = h.Nums[n-1]
	h.Nums = h.Nums[:n-1]
	h.heapifyDown(0)
	return root
}
func (h *MinHeap) heapifyDown(i int) {
	n := len(h.Nums)
	if n > 0 {
		left := leftChild(i)
		right := rightChild(i)
		for left < n {
			cc := left
			if right < n && h.Nums[right][0] < h.Nums[left][0] {
				cc = right
			}
			if h.Nums[i][0] > h.Nums[cc][0] {
				h.swap(i, cc)
				i = cc
				left = leftChild(i)
				right = rightChild(i)
			} else {
				return
			}
		}
	}
}
func parent(i int) int {
	return (i - 1) / 2
}
func leftChild(i int) int {
	return (i * 2) + 1
}
func rightChild(i int) int {
	return (i * 2) + 2
}
func (h *MinHeap) swap(i1, i2 int) {
	h.Nums[i1], h.Nums[i2] = h.Nums[i2], h.Nums[i1]
}
