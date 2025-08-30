package practice_graphs

func CheckGraphCycleUsingStack(mtx [][]int) bool {
	visited := make([]bool, len(mtx))

	var checkForCycle = func(src int, mtx [][]int) bool {
		stack := [][]int{}
		stack = append(stack, []int{src, -1})
		visited[src] = true
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node, parent := top[0], top[1]
			for neigh, edge := range mtx[node] {
				if edge == 1 {
					if !visited[neigh] {
						stack = append(stack, []int{neigh, node})
						visited[neigh] = true
					} else if neigh != parent {
						return true
					}
				}
			}
		}
		return false
	}

	for i := 0; i < len(visited); i++ {
		if !visited[i] && checkForCycle(i, mtx) {
			return true
		}
	}
	return false
}
func CheckGraphCycleUsingQueue(mtx [][]int) bool {
	visited := make([]bool, len(mtx))

	var checkForCycle = func(src int, mtx [][]int) bool {
		queue := [][]int{}
		queue = append(queue, []int{src, -1})
		visited[src] = true
		for len(queue) > 0 {
			top := queue[0]
			queue = queue[1:]
			node, parent := top[0], top[1]
			for neigh, edge := range mtx[node] {
				if edge == 1 {
					if !visited[neigh] {
						queue = append(queue, []int{neigh, node})
						visited[neigh] = true
					} else if neigh != parent {
						return true
					}
				}
			}
		}
		return false
	}

	for i := 0; i < len(visited); i++ {
		if !visited[i] && checkForCycle(i, mtx) {
			return true
		}
	}
	return false
}

