package code

func missingNumber(nums []int) int {
	var missing = len(nums)
	for i, num := range nums {
		missing ^= i ^ num
	}
	return missing
}
