#### 题目
给定 *n* 个非负整数，在坐标内画 *n* 条垂直线，找出其中的两条线，使得它们与 *x* 轴共同构成的容器可以容纳最多的水。
**说明：** 你不能倾斜容器，且 *n* 的值至少为2。
 ![题目](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-06/001103.png)
**实例：**
```$xslt
输入：[1,8,6,2,5,4,8,3,7]
输出：49
```

#### 题解
1、暴力法
```go
func maxArea(height []int) int {
	var maxArea int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			length := j-i
			var minHeight = height[i]
			if height[j] < minHeight {
				minHeight = height[j]
			}
			var area = length * minHeight
			if maxArea < area {
				maxArea = area
			}
		}
	}
	return maxArea
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-06/001101.png)
时间复杂度：O(n^2^)，空间复杂度O(1)

2、双指针法
由于面积取决于较短的那一端，所以要想获取到较大面积，必须舍弃较短的那一端
```go
func maxArea(height []int) int {
	var left,right = 0,len(height)-1
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
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-06/001102.png)
时间复杂度：O(n),空间复杂度O(1)