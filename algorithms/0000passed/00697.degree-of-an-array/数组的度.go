package code

func findShortestSubArray(nums []int) (ans int) {
	var set = make(map[int]entry)
	// 记录元素
	for i, v := range nums {
		if e, has := set[v]; has {
			e.cnt++
			e.right = i
			set[v] = e
		} else {
			set[v] = entry{i, i, 1}
		}
	}
	var maxCnt int
	for _, e := range set {
		if e.cnt > maxCnt {
			maxCnt, ans = e.cnt, e.right-e.left+1
		} else if e.cnt == maxCnt {
			ans = min(ans, e.right-e.left+1)
		}
	}
	return
}

type entry struct {
	left, right, cnt int // 记录元素第一次和最后一次出现的位置及次数
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
