package practice_graphs

func NumIslands2(m int, n int, queries [][]int) []int {
	uf := newUnionFind(m * n)
	visited := make([]bool, m*n)
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	ans := make([]int, len(queries))

	islands := 0
	for q, query := range queries {
		row := query[0]
		col := query[1]
		id := GetCellId(row, col, n)
		if !visited[id] {
			visited[id] = true
			islands += 1
			for i := 0; i < 4; i++ {
				newR := row + dirs[i][0]
				newC := col + dirs[i][1]
				newId := GetCellId(newR, newC, n)
				if newR >= 0 && newR < m && newC >= 0 && newC < n && visited[newId] && uf.AddEdge(id, newId) {
					islands -= 1
				}
			}
		}
		ans[q] = islands
	}
	return ans
}

type unionFind struct{
    Size []int
    Parent []int
}
func newUnionFind(n int)*unionFind{
    uf:= &unionFind{
        Size: make([]int, n),
        Parent: make([]int, n),
    }
    for i:=0; i<n; i++{
        uf.Size[i] = 1
        uf.Parent[i] = i
    }
    return uf
}
func (uf *unionFind) GetParent(i int)int{
    if uf.Parent[i] == i{
        return i
    }
    uf.Parent[i] = uf.GetParent(uf.Parent[i])
    return uf.Parent[i]
}
func (uf *unionFind) AddEdge(u, v int)bool{
    pu:= uf.GetParent(u)
    pv:= uf.GetParent(v)
    if pu == pv{return false}
    if uf.Size[pu] >= uf.Size[pv]{
        uf.Parent[pv] = pu
        uf.Size[pu] += uf.Size[pv]
    }else{
        uf.Parent[pu] = pv
        uf.Size[pv] += uf.Size[pu]
    }
    return true
}
func GetCellId(row, col, n int)int{
    return (row*n) + col
}