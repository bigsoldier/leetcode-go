package code

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	size := len(nums)
	if k == 0 || t < 0 || size <= 1 {
		return false
	}
	t++
	nMap := map[int]int{}
	for i, num := range nums {
		m := num / t
		if num < 0 {
			m--
		}
		if _, has := nMap[m]; has {
			return true
		} else if n, ok := nMap[m-1]; ok && abs(n, num) < t {
			return true
		} else if n, ok = nMap[m+1]; ok && abs(n, num) < t {
			return true
		}
		nMap[m] = num
		if i >= k {
			delete(nMap, nums[i-k]/t)
		}
	}

	return false
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
