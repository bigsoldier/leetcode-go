package code

func majorityElement(nums []int) int {
	count, result := 0, -1
	for _, num := range nums {
		if count == 0 {
			result = num
		}
		if num == result {
			count++
		} else {
			count--
		}
	}
	return result
}
