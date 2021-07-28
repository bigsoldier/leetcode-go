package main

// 数组分成 m 个非空的连续子数组，使得这 m 个子数组各自和的最大值最小。
func splitArray(nums []int, m int) int {
	left, right := getMax(nums)
	for left < right {
		mid := left + (right-left)/2
		n := split(nums, mid)
		if n > m {
			left = mid + 1

		} else {
			right = mid
		}
	}
	return left
}

// 数组可以被分成几份
func split(nums []int, max int) int {
	count := 1
	var sum int
	for i := 0; i < len(nums); i++ {
		if nums[i]+sum > max {
			count++
			sum = nums[i]
		} else {
			sum += nums[i]
		}
	}
	return count
}

func getMax(nums []int) (int, int) {
	var max, sum = nums[0], 0
	for _, num := range nums {
		if num > max {
			max = num
		}
		sum += num
	}
	return max, sum
}
