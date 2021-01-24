package code

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var cnt, max int
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cnt++
		} else {
			cnt = 0
		}
		if cnt > max {
			max = cnt
		}
	}
	return max + 1
}
