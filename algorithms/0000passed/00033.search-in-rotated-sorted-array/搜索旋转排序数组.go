package code

func search(nums []int, target int) int {
	// 二分法
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}

		// 前半部分有序:[left:mid]递增
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && target > nums[mid] { // 在后半部分有序中
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
