package code

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var solution []int
	var result [][]int
	findCombinationSum(candidates, solution, target, &result)
	return result
}

func findCombinationSum(candidates, solution []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, solution)
	}

	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	findCombinationSum(candidates[1:], append(solution, candidates[0]), target-candidates[0], result)
	var i int
	for i < len(candidates)-1 {
		if candidates[i] != candidates[i+1] {
			findCombinationSum(candidates[i+1:], solution, target, result)
			break
		} else {
			i++
		}
	}
}
