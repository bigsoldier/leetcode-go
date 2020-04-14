package code

func canJump(nums []int) bool {
	var step int
	for step < len(nums)-1 { // 这里 step<len(nums)&&step!=len(nums)-1,表示step要比数组长度小并且不在最后一位上
		var maxStep, index int
		if nums[step] <= 0 { // 如果当前位置的元素为0,则肯定跳不出去了
			return false
		}
		for i := 1; i <= nums[step]; i++ {
			if step+i >= len(nums) {
				return true
			}
			temp := i + nums[step+i]
			if temp > maxStep {
				maxStep = temp
				index = i
			}
		}
		step += index
	}
	return true
}
