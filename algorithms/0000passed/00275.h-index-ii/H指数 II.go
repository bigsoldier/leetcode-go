package code

func hIndex(citations []int) int {
	n := len(citations)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if citations[n-mid-1] > mid {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}
