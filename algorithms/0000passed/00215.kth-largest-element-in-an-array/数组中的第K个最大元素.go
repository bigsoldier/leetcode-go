package code

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	buildHeap(nums, n)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		n--
		maxHeapify(nums, 0, n)
	}
	return nums[0]
}

func buildHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	curr, left, right := i, i*2+1, i*2+2
	if left < heapSize && a[left] > a[curr] {
		curr = left
	}
	if right < heapSize && a[right] > a[curr] {
		curr = right
	}
	if curr != i {
		a[i], a[curr] = a[curr], a[i]
		maxHeapify(a, curr, heapSize)
	}
}
