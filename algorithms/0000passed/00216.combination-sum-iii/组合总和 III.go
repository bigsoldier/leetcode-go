package code

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	dfs(k, n, 1, nil, &result)
	return result
}

func dfs(k, n int, idx int, ret []int, result *[][]int) {
	if n == 0 && k == 0 {
		a := make([]int, len(ret))
		copy(a, ret)
		*result = append(*result, a)
		return
	}
	for i := idx; i <= 9; i++ {
		dfs(k-1, n-i, i+1, append(ret, i), result)
	}
}
