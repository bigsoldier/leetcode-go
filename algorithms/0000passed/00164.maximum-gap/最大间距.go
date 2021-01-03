package code

func maximumGap(nums []int) (ans int) {
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	nums = sort(nums, max)
	for i := 1; i < len(nums); i++ {
		n := nums[i] - nums[i-1]
		if n > ans {
			ans = n
		}
	}
	return
}

func sort(nums []int, max int) []int {
	n := len(nums)
	buf := make([]int, n)
	for e := 1; e <= max; e = e * 10 {
		cnt := [10]int{} // 0-9
		for _, v := range nums {
			digit := v / e % 10 // 取个位数
			cnt[digit]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			digit := nums[i] / e % 10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		copy(nums, buf)
	}
	return buf
}
