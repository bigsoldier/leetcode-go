package code

func getPermutation(n int, k int) string {
	k-- // 从0开始
	if n == 0 {
		return ""
	}
	ret := make([]byte, n)

	// 存放 从1到n 的数字
	nums := make([]byte, n)
	for i := 0; i < n; i++ {
		nums[i] = byte(i) + '1'
	}

	// 阶乘
	fib := 1
	for i := 2; i < n; i++ {
		fib *= i
	}

	for i := 0; i < n-1; i++ {
		idx := k / fib
		ret[i] = nums[idx]
		// 去除已使用的数
		nums = append(nums[:idx], nums[idx+1:]...)
		k %= fib
		fib /= (n - i - 1)
	}
	ret[n-1] = nums[0]
	return string(ret)
}
