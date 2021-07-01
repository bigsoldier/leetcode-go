package main

// 图遍历，有向无环图的所有可能路径
func allPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	traverse3(graph, 0, &result, nil)
	return result
}

func traverse3(graph [][]int, key int, ret *[][]int, tmp []int) {
	tmp = append(tmp, key)
	n := len(graph)
	if key == n-1 {
		l := make([]int, len(tmp))
		copy(l, tmp)
		*ret = append(*ret, l)
		tmp = tmp[:len(tmp)-1]
		return
	}
	for _, v := range graph[key] {
		traverse3(graph, v, ret, tmp)
	}
	tmp = tmp[:len(tmp)-1]
}
