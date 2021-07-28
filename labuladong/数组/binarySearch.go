package main

// 二分查找
func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	return []int{searchLeft(nums, target), searchRight(nums, target)}
}

func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

// 吃香蕉的珂珂
func minEatingSpeed(piles []int, h int) int {
	// 二分查找吃香蕉的速度
	left, right := 1, 1000000000
	for left <= right {
		mid := left + (right-left)/2
		x := eating(piles, mid)
		if x <= h {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func eating(piles []int, x int) int {
	var hours int
	for i := 0; i < len(piles); i++ {
		hours += piles[i] / x
		if piles[i]%x > 0 {
			hours++
		}
	}
	return hours
}

// 在 D 天内送达包裹的能力
func shipWithinDays(weights []int, days int) int {
	var left, right int
	for _, w := range weights {
		left = max(left, w)
		right += w
	}

	for left <= right {
		mid := left + (right-left)/2
		if ship(weights, mid) > days {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func ship(weights []int, weight int) int {
	var days int
	for i := 0; i < len(weights); {
		caps := weight
		for i < len(weights) {
			if caps < weights[i] {
				break
			} else {
				caps -= weights[i]
			}
			i++
		}
		days++
	}
	return days
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
