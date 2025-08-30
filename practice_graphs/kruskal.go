package practice_graphs

import "sort"

func KruskalAlgorithm(graph [][]int, n int) int {
	uf := NewunionBySize(n)
	edges := [][]int{}
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			edgeWeight := graph[i][j]
			if edgeWeight != 0 {
				edges = append(edges, []int{edgeWeight, i, j})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})
	sum := 0
	for _, edge := range edges {
		if uf.AddEdge(edge[1], edge[2]) {
			sum += edge[0]
		}
	}
	return sum
}

type unionBySize struct {
	Size   []int
	Parent []int
}

func NewunionBySize(n int) *unionBySize {
	ubs := &unionBySize{
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
func (ubs *unionBySize) GetParent(i int) int {
	if i == ubs.Parent[i] {
		return i
	}
	ubs.Parent[i] = ubs.GetParent(ubs.Parent[i])
	return ubs.Parent[i]
}
func (ubs *unionBySize) AddEdge(u, v int) bool {
	pu := ubs.GetParent(u)
	pv := ubs.GetParent(v)
	if pu == pv {
		return false
	}
	if ubs.Size[pu] > ubs.Size[pv] {
		ubs.Parent[pv] = pu
		ubs.Size[pu] += ubs.Size[pv]
	} else {
		ubs.Parent[pu] = pv
		ubs.Size[pv] += ubs.Size[pu]
	}
	return true
}
func (ubs *unionBySize) IsConnected(u, v int) bool {
	return ubs.GetParent(u) == ubs.GetParent(v)
}
