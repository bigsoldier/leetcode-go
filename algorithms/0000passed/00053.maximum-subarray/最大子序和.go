package code

func maxSubArray(nums []int) int {
	var sum = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if sum < nums[i] {
			sum = nums[i]
		}
	}
	return sum
}
