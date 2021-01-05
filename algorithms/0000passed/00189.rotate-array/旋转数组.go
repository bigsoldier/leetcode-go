package code

func rotate(nums []int, k int) {
	n := len(nums)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		a[(i+k)%n] = nums[i]
	}
	for i := 0; i < n; i++ {
		nums[i] = a[i]
	}
}
