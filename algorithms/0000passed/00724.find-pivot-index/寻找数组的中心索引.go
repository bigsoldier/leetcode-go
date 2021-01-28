package code

func pivotIndex(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	var left int
	for i := 0; i < len(nums); i++ {
		if sum-left-nums[i] == left {
			return i
		}
		left += nums[i]
	}
	return -1
}
