package main

// 单调栈

// 下一个更大的元素
func nextGreaterElement(nums []int) []int {
	var result = make([]int, len(nums))
	var stack []int
	// 倒着放入nums2
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = -1
		} else {
			result[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return result
}

// 下一个更暖和的日子
func dailyTemperatures(temperatures []int) []int {
	var result = make([]int, len(temperatures))
	var stack []int
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return result
}

// 环形数组
func nextGreaterElements(nums []int) []int {
	// 数组翻倍
	n := len(nums)
	result := make([]int, n)
	var stack []int
	for i := 2*n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%n] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i%n] = -1
		} else {
			result[i%n] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%n])
	}
	return result
}
