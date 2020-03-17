package code

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// 这里使用for循环，是因为交换完成后该位置的新值也需要进行处理
		for nums[i]-1 >= 0 && nums[i]-1 < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
		// 这里的时间复杂度是O(n),如果元素在一次遍历里已经被交换到他们该在的位置上，那么在执行的时候会跳过。
		// 最坏的情况是第一次里循环的时候把所有元素都交换了一遍，那么之后的循环则不会再进入里循环
	}
	for j := range nums {
		if j+1 != nums[j] {
			return j + 1
		}
	}
	return len(nums) + 1
}
