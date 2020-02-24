package leetcode

//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// Example:
//
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].
//
// Related Topics Array Hash Table

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

// 可以暴力破解，两层循环遍历，时间复杂度O（nlogn）
// 顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。
// 如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。时间复杂度O（n）
