package code

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var solution []int
	var result [][]int
	sort.Ints(candidates)
	findCombiationSum(candidates, solution, target, &result)
	return result
}

// solution 存放当前的可能解
func findCombiationSum(candidates, solution []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, solution)
	}
	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	solution = solution[:len(solution):len(solution)]
	// 等价于
	// var ret = make([]int,len(solution))
	// copy(ret,solution)
	findCombiationSum(candidates, append(solution, candidates[0]), target-candidates[0], result)
	findCombiationSum(candidates[1:], solution, target, result)
}
