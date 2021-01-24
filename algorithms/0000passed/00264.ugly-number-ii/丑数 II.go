package code

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}
	var res = make([]int, n)
	res[0] = 1
	pos := []int{0, 0, 0}
	factors := []int{2, 3, 5}
	candidates := []int{2, 3, 5}
	for i := 1; i < n; i++ {
		res[i] = min(candidates)
		for j := 0; j < 3; j++ {
			if res[i] == candidates[j] {
				pos[j]++
				candidates[j] = res[pos[j]] * factors[j]
			}
		}
	}
	return res[n-1]
}

func min(nums []int) int {
	var m = nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}
