package practice_graphs

func countIslands(mtx [][]int) int{
	visited := make([]bool, len(mtx))

	var bfs = func(src int, mtx [][]int) {
		queue := []int{}
		queue = append(queue, src)
		visited[src] = true
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			for neigh, edge := range mtx[node] {
				if edge == 1 && !visited[neigh]{
					queue = append(queue, neigh)
					visited[neigh] = true
				}
			}
		}
	}
    islands:= 0
	for i := 0; i < len(visited); i++ {
		if !visited[i]{
		    islands ++
		    bfs(i, mtx)
		}
	}
	return islands
}