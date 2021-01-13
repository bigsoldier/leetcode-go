package code

func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}
	return merge(buildings, 0, len(buildings)-1)
}

func merge(buildings [][]int, start, end int) (ans [][]int) {
	if start == end {
		// 只有一个建筑，将结果集加入
		ans = append(ans,
			[]int{buildings[start][0], buildings[start][2]},
			[]int{buildings[start][1], 0})
		return ans
	}
	mid := start + (end-start)/2
	left := merge(buildings, start, mid)
	right := merge(buildings, mid+1, end)

	// 合并解
	var i, j, h1, h2 int
	for i < len(left) || j < len(right) {
		var x1, x2 = math.MaxInt64, math.MaxInt64
		if i < len(left) {
			x1 = left[i][0]
		}
		if j < len(right) {
			x2 = right[j][0]
		}
		// 比较两个坐标,取较小的那一边
		if x1 < x2 {
			h1 = left[i][1]
			i++
		} else if x1 > x2 {
			h2 = right[j][1]
			x1 = x2
			j++
		} else {
			h1 = left[i][1]
			h2 = right[j][1]
			i++
			j++
		}
		height := max(h1, h2)
		if len(ans) == 0 || height != ans[len(ans)-1][1] {
			ans = append(ans, []int{x1, height})
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
