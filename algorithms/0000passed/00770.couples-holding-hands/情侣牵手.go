package code

func minSwapsCouples(row []int) (ans int) {
	n := len(row)
	graph := make([][]int, n/2)
	for i := 0; i < n; i += 2 {
		l, r := row[i]/2, row[i+1]/2
		if l != r {
			graph[l] = append(graph[l], r)
			graph[r] = append(graph[r], l)
		}
	}
	visited := make([]bool, n/2)
	for i, v := range visited {
		if !v {
			visited[i] = true
			cnt := 0
			q := []int{i}
			for len(q) > 0 {
				cnt++
				a := q[0]
				q = q[1:]
				for _, w := range graph[a] {
					if !visited[w] {
						visited[w] = true
						q = append(q, w)
					}
				}
			}
			ans += cnt - 1
		}
	}
	return
}
