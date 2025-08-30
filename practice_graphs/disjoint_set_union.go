package practice_graphs

type UnionBySize struct {
	Size   []int
	Parent []int
}

func NewUnionBySize(n int) *UnionBySize {
	ubs := &UnionBySize{
		Size:   make([]int, n+1),
		Parent: make([]int, n+1),
	}
	for i := 0; i < n+1; i++ {
		ubs.Size[i] = 1
	}
	for i := 0; i < n+1; i++ {
		ubs.Parent[i] = i
	}
	return ubs
}
func (ubs *UnionBySize) GetParent(i int) int {
	if i == ubs.Parent[i] {
		return i
	}
	ubs.Parent[i] = ubs.GetParent(ubs.Parent[i])
	return ubs.Parent[i]
}
func (ubs *UnionBySize) AddEdge(u, v int) {
	pu := ubs.GetParent(u)
	pv := ubs.GetParent(v)
	if pu == pv {
		return
	}
	if ubs.Size[pu] > ubs.Size[pv] {
		ubs.Parent[pv] = pu
		ubs.Size[pu] += ubs.Size[pv]
	} else {
		ubs.Parent[pu] = pv
		ubs.Size[pv] += ubs.Size[pu]
	}
}
func (ubs *UnionBySize) IsConnected(u, v int) bool {
	return ubs.GetParent(u) == ubs.GetParent(v)
}
