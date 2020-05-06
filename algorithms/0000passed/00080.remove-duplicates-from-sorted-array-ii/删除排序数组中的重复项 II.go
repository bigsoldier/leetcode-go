package code

func removeDuplicates(nums []int) int {
	var i, count = 1, 1
	for i < len(nums) {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count > 2 {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}
