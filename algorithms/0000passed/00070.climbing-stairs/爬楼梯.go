package code

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var climb = make([]int, n)
	climb[0] = 1
	climb[1] = 2
	for i := 2; i < n; i++ {
		climb[i] = climb[i-2] + climb[i-1]
	}
	return climb[n-1]
}
