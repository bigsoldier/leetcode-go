package code

func subsets(nums []int) [][]int {
	var used []int
	var result [][]int
	add(&result, used, nums)
	return result
}

func add(result *[][]int, used, nums []int) {
	var c = make([]int, len(used))
	copy(c, used)
	*result = append(*result, c)
	for i, num := range nums {
		add(result, append(used, num), nums[i+1:])
	}
}
