package practice_graphs

func LargestIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	uf := NewUnionFind(m * n)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	//! make conn components
	for i, row := range grid {
		for j, cell := range row {
			id := GetCellId(i, j, n)
			if cell == 1 {
				for d := 0; d < len(dirs); d++ {
					newR := i + dirs[d][0]
					newC := j + dirs[d][1]
					if newR >= 0 && newR < m && newC >= 0 && newC < n && grid[newR][newC] == 1 {
						newId := GetCellId(newR, newC, n)
						uf.AddEdge(id, newId)
					}
				}
			}
		}
	}

	//! convert 0 to 1
	maxSize := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 {
				neighbours := make(map[int]struct{})
				for d := 0; d < len(dirs); d++ {
					newR := i + dirs[d][0]
					newC := j + dirs[d][1]
					if newR >= 0 && newR < m && newC >= 0 && newC < n && grid[newR][newC] == 1 {
						newCellId := GetCellId(newR, newC, n)
						//! add parents to map
						p := uf.GetParent(newCellId)
						neighbours[p] = struct{}{}
					}
				}
				//! range map to get sizes
				islandSize := 1
				for neighbour, _ := range neighbours {
					islandSize += uf.Size[neighbour]
				}
				if islandSize > maxSize {
					maxSize = islandSize
				}
			}
		}
	}
	if maxSize == 0 {
		return m * n
	}
	return maxSize
}

type UnionFind struct {
	Size   []int
	Parent []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		Size:   make([]int, n),
		Parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.Size[i] = 1
		uf.Parent[i] = i
	}
	return uf
}
func (uf *UnionFind) GetParent(i int) int {
	if uf.Parent[i] == i {
		return i
	}
	uf.Parent[i] = uf.GetParent(uf.Parent[i])
	return uf.Parent[i]
}
func (uf *UnionFind) AddEdge(u, v int) {
	pu := uf.GetParent(u)
	pv := uf.GetParent(v)
	if pu == pv {
		return
	}
	if uf.Size[pu] >= uf.Size[pv] {
		uf.Parent[pv] = pu
		uf.Size[pu] += uf.Size[pv]
	} else {
		uf.Parent[pu] = pv
		uf.Size[pv] += uf.Size[pu]
	}
}
