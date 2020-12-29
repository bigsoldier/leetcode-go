package code

func singleNumber(nums []int) int {
	var once, twice int
	for _, num := range nums {
		once = ^twice & (once ^ num)
		twice = ^once & (twice ^ num)
	}
	return once
}
