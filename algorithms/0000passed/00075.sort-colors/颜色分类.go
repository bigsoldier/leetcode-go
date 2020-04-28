package code

func sortColors(nums []int) {
	p0, curr, p2 := 0, 0, len(nums)-1
	for curr <= p2 {
		if nums[curr] == 0 {
			nums[p0], nums[curr] = nums[curr], nums[p0]
			p0++
			curr++
		} else if nums[curr] == 2 {
			nums[curr], nums[p2] = nums[p2], nums[curr]
			p2--
		} else {
			curr++
		}
	}
}
