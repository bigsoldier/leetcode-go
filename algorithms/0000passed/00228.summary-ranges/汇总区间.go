package code

import "strconv"

func summaryRanges(nums []int) (ans []string) {
	for i := 0; i < len(nums); {
		var j = i + 1
		var pre = nums[i]
		prefix := strconv.Itoa(nums[i])
		var suffix string
		for j < len(nums) {
			if nums[j]-pre == 1 {
				pre = nums[j]
				suffix = "->" + strconv.Itoa(nums[j])
			} else {
				break
			}
			j++
		}
		ans = append(ans, prefix+suffix)
		i = j
	}
	return
}
