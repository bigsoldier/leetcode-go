package code

func trap(height []int) int {
	var area, maxLeft, maxRight int
	var left, right = 1, len(height) - 2
	for i := 1; i < len(height)-1; i++ {
		if height[left-1] < height[right+1] {
			if maxLeft < height[left-1] {
				maxLeft = height[left-1]
			}
			var min = maxLeft
			if min > height[left] {
				area += min - height[left]
			}
			left++
		} else {
			if maxRight < height[right+1] {
				maxRight = height[right+1]
			}
			var min = maxRight
			if min > height[right] {
				area += min - height[right]
			}
			right--
		}
	}
	return area
}
