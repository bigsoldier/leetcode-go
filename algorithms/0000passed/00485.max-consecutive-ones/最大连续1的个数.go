package code

func findMaxConsecutiveOnes(nums []int) (ans int) {
	var cnt int
	for _, num := range nums {
		if num == 1 {
			cnt++
		} else {
			ans = max(ans, cnt)
			cnt = 0
		}
	}
	ans = max(ans, cnt)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
