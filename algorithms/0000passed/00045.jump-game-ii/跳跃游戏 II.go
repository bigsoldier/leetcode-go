package code

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var index, count int
	for index < len(nums) {
		if nums[index]+index+1 >= len(nums) {
			return count + 1
		}
		var maxStep int  // 记录当前跳跃范围内可统计的最大步数
		var maxIndex int // 最大步数的index
		// nums[index]表示当前位置的最大跳跃步数
		for i := 1; i <= nums[index]; i++ { // 遍历可跳的最大步数
			if nums[index+i]+i > maxStep {
				maxStep = nums[index+i] + i
				maxIndex = i
			}
		}
		index += maxIndex
		count++
	}
	return count
}
