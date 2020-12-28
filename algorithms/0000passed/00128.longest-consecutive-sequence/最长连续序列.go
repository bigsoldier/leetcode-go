package code

func longestConsecutive(nums []int) int {
	var m = map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	var longestStreak int
	for _, num := range nums {
		if !m[num-1] {
			cur := num
			curStreak := 1
			for m[cur+1] {
				cur++
				curStreak++
			}
			if longestStreak < curStreak {
				longestStreak = curStreak
			}
		}
	}
	return longestStreak
}
