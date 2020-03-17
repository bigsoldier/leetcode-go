package code

func searchRange(nums []int, target int) []int {
	var index = []int{-1, -1}
	var leftIndex = insertIndex(nums, target, true)
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return index
	}
	index[0] = leftIndex
	index[1] = insertIndex(nums, target, false) - 1
	return index
}

// 遍历列表，haw为true时，找左边界，haw为false时，找右边界
func insertIndex(nums []int, target int, haw bool) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if target < nums[mid] || (haw && target == nums[mid]) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
