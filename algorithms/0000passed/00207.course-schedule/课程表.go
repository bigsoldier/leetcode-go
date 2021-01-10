package code

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses) // 记录某课程完成后可以进行的课程
		indeg  = make([]int, numCourses)   // 记录课程的先修量
		result []int
	)

	for _, info := range prerequisites {
		// info第一位表示后修课程，第二位表示先修课程
		edges[info[1]] = append(edges[info[1]], info[0]) // 这里添加先修课程完之后对应的后修课程
		indeg[info[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		course := q[0]
		q = q[1:]
		result = append(result, course)
		for _, v := range edges[course] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}
