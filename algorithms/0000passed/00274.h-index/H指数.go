package code

func hIndex(citations []int) int {
	quickSort(citations, 0, len(citations)-1)
	var i int
	for ; i < len(citations) && citations[len(citations)-i-1] > i; i++ {
	}
	return i
}

func quickSort(nums []int, left, right int) {
	if left < right {
		mid := partition(nums, left, right)
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}
}

func partition(nums []int, left, right int) int {
	var i, j = left, left + 1
	for idx := j; idx <= right; idx++ {
		if nums[idx] < nums[i] {
			nums[j], nums[idx] = nums[idx], nums[j]
			j++
		}
	}
	nums[i], nums[j-1] = nums[j-1], nums[i]
	return j - 1
}
