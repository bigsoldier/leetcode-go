package code

func minCostClimbingStairs(cost []int) int {
	var pre, cur int
	for i := 2; i <= len(cost); i++ {
		pre, cur = cur, min(pre+cost[i-2], cur+cost[i-1])
	}
	return cur
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
