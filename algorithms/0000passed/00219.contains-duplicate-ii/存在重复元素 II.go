package code

func containsNearbyDuplicate(nums []int, k int) bool {
	var set = map[int]int{}
	for i, num := range nums {
		if set[num] != 0 && (i-set[num]+1) <= k {
			return true
		}
		set[num] = i + 1
	}
	return false
}
