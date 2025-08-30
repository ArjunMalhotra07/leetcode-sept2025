package practiceheaps

type MinHeap struct {
	Nums []int
}

func (h *MinHeap) add(n int) {
	h.Nums = append(h.Nums, n)
	h.heapifyUp(len(h.Nums) - 1)
}
func (h *MinHeap) heapifyUp(i int) {
	for h.Nums[i] < h.Nums[parent(i)] {
		h.swap(i, parent(i))
		i = parent(i)
	}
}
func (h *MinHeap) extractRoot() int {
	n := len(h.Nums)
	root := -1
	if n > 0 {
		root = h.Nums[0]
		h.Nums[0] = h.Nums[n-1]
		h.Nums = h.Nums[:n-1]
		h.heapifyDown(0)
	}
	return root
}
func (h *MinHeap) heapifyDown(i int) {
	n := len(h.Nums)
	if n > 0 {
		left := leftChild(i)
		right := rightChild(i)
		for left < n {
			cc := left
			if right < n && h.Nums[right] < h.Nums[left] {
				cc = right
			}
			if h.Nums[i] > h.Nums[cc] {
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
