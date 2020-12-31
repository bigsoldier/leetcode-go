package code

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans, right := 1, intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= right {
			ans++
			right = intervals[i][1]
		}
	}
	return len(intervals) - ans
}
