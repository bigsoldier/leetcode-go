package code

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	var i int
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
