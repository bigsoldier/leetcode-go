package code

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses) // 0表示未搜索，1表示进行中，2表示已完成
		valid   = true                    // 是否无环
		dfs     func(i int)
		result  []int
	)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
	}
	dfs = func(i int) {
		visited[i] = 1
		for _, v := range edges[i] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[i] = 2
		result = append(result, i)
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return nil
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return result
}
