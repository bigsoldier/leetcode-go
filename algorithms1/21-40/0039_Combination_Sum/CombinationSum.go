package _039_Combination_Sum

//Given a set of candidate numbers (candidates) (without duplicates) and a targe
//t number (target), find all unique combinations in candidates where the candidat
//e numbers sums to target.
//
// The same repeated number may be chosen from candidates unlimited number of ti
//mes.
//
// Note:
//
//
// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
//
//
// Example 1:
//
//
//Input: candidates = [2,3,6,7], target = 7,
//A solution set is:
//[
//  [7],
//  [2,2,3]
//]
//
//
// Example 2:
//
//
//Input: candidates = [2,3,5], target = 8,
//A solution set is:
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//]
//
// Related Topics Array Backtracking

//leetcode submit region begin(Prohibit modification and deletion)
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

//leetcode submit region end(Prohibit modification and deletion)
