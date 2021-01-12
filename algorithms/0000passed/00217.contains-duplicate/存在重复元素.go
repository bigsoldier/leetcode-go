package code

func containsDuplicate(nums []int) bool {
	quickSort(nums, 0, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func quickSort(nums []int, left, right int) {
	if left < right {
		mid := partition(nums, left, right)
		quickSort(nums, left, mid)
		quickSort(nums, mid+1, right)
	}
}

func partition(nums []int, left, right int) int {
	x, i := nums[left], left
	for j := left + 1; j < right; j++ {
		if x > nums[j] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}
