package main

// 数组中滑动窗口k的最大值

type queue struct {
	q []int
}

func (q *queue) push(x int) {
	for len(q.q) > 0 && q.q[len(q.q)-1] < x {
		q.q = q.q[:len(q.q)-1]
	}
	q.q = append(q.q, x)
}

func (q *queue) max() int {
	return q.q[0]
}

func (q *queue) pop(x int) {
	if x == q.q[0] {
		q.q = q.q[:len(q.q)-1]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	var windows = &queue{}
	var result []int
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			windows.push(nums[i])
		} else {
			windows.push(nums[i])
			result = append(result, windows.max())
			windows.pop(nums[i-k+1])
		}
	}
	return result
}
