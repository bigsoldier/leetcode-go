/**
 *Created by Xie Jian on 2020/1/16 10:50
 */
package _011_Container_With_Most_Water

func maxArea(height []int) int {
	var left, right = 0, len(height) - 1
	var maxArea int
	for left < right {
		h := height[left] //  较小的一端
		if h > height[right] {
			h = height[right]
		}
		area := h * (right - left)
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
