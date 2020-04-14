package code

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0, len(intervals))

	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= temp[1] {
			if temp[1] < intervals[i][1] {
				temp[1] = intervals[i][1]
			}
		} else {
			result = append(result, temp)
			temp = intervals[i]
		}
	}
	result = append(result, temp)
	return result
}
