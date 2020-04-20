package code

func insert(intervals [][]int, newInterval []int) [][]int {
	var ret [][]int

	var curIndex int
	for curIndex < len(intervals) && intervals[curIndex][1] < newInterval[0] {
		ret = append(ret, intervals[curIndex])
		curIndex++
	}
	for curIndex < len(intervals) && intervals[curIndex][0] <= newInterval[1] {
		if intervals[curIndex][0] < newInterval[0] {
			newInterval[0] = intervals[curIndex][0]
		}
		if intervals[curIndex][1] > newInterval[1] {
			newInterval[1] = intervals[curIndex][1]
		}
		curIndex++
	}
	ret = append(ret, newInterval)
	for curIndex < len(intervals) {
		ret = append(ret, intervals[curIndex])
		curIndex++
	}

	return ret
}
