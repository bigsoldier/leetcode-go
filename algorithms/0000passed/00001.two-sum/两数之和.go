package code

func twoSum(nums []int, target int) []int {
	// 把数据放到map中
	var m = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		res := target - nums[i]
		if _, ok := m[res]; ok {
			return []int{m[res], i}
		}
		m[nums[i]] = i
	}
	return nil
}
