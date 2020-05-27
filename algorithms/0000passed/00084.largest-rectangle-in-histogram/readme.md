#### 题目
<p>给定 <em>n</em> 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。</p>

<p>求在该柱状图中，能够勾勒出来的矩形的最大面积。</p>

<p>&nbsp;</p>

<p><img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/histogram.png"></p>

<p><small>以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为&nbsp;<code>[2,1,5,6,2,3]</code>。</small></p>

<p>&nbsp;</p>

<p><img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/histogram_area.png"></p>

<p><small>图中阴影部分为所能勾勒出的最大矩形面积，其面积为&nbsp;<code>10</code>&nbsp;个单位。</small></p>

<p>&nbsp;</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> [2,1,5,6,2,3]
<strong>输出:</strong> 10</pre>


 #### 题解
 1、暴力法
 列举出所有两两柱子之间的面积
 ```go
func largestRectangleArea(heights []int) int {
	var maxArea int
	// 计算所有两两柱子之间的面积
	for i := 0; i < len(heights); i++ {
		for j := i; j < len(heights); j++ {
			minHeight := heights[i]
			for k := i; k <= j; k++ {
				if heights[k] < minHeight {
					minHeight = heights[k]
				}
			}
			area := (j-i+1)*minHeight
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
```
 时间复杂度O(n^3^),空间复杂度O(1)
 
 2、优化暴力法
 在运行中比较最小值
 ```go
func largestRectangleArea(heights []int) int {
	var maxArea int
	// 计算所有两两柱子之间的面积
	for i := 0; i < len(heights); i++ {
		var minHeight = heights[i]
		for j := i; j < len(heights); j++ {
			if heights[j] < minHeight {
				minHeight = heights[j]
			}
			area := (j-i+1)*minHeight
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-13/008401.png)
 时间复杂度O(n^2^),空间复杂度O(1)
 
 3、分而治之
 最大矩阵面积：
 1、确定最矮柱子，矩阵的宽尽可能向两边延伸
 2、最矮柱子左边的最大面积(子问题)
 3、最矮柱子右边的最大面积(子问题)
 ```go
func largestRectangleArea(heights []int) int {
	return calculateArea(heights, 0, len(heights)-1)
}

func calculateArea(heights []int, start, end int) int {
	if start > end {
		return 0
	}

	// 找出最低高度
	minIndex := start
	for i := start; i <= end; i++ {
		if heights[minIndex] > heights[i] {
			minIndex = i
		}
	}
	maxArea := heights[minIndex]*(end-start+1)
	left := calculateArea(heights,start,minIndex-1)
	right := calculateArea(heights,minIndex+1,end)
	if maxArea < left {
		maxArea = left
	}
	if maxArea < right {
		maxArea = right
	}
	return maxArea
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-13/008402.png)
 时间复杂度O(nlogn),最坏O(n^2^),是有序数组,空间复杂度O(n)
 
 4、栈
 ```go
func largestRectangleArea(heights []int) int {
	// 在 heights 中原本是非负的数字
	// 再首尾添加了 -2 和 -1 后，简化 for 循环中 begin 的计算
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)

	size := len(heights)

	// stack 中存的是 heights 的元素的索引号
	// stack 中索引号对应的 heights 中的值，是递增的。
	// e.g.
	//     stack = []int{1,3,5}，那么
	//     heights[1] < heights[3] < heights[5]
	stack := make([]int, 1, size)
	// 把 heights[0] 的索引号，已经放入了 stack
	// end 从 1 开始
	end := 1

	res := 0
	for end < size {
		// end 指向了新高，就把 end 放入 stack 后，指向下一个
		if heights[stack[len(stack)-1]] < heights[end] {
			stack = append(stack, end)
			end++
			continue
		}

		begin := stack[len(stack)-2]
		index := stack[len(stack)-1]
		height := heights[index]
		// area 是 heights(begin:end) 之间的最大方块的面积，因为
		// anyone of heights(begin:index) > height > anyone of heights(index:end)
		area := (end - begin - 1) * height

		if res < area {
			res = area
		}

		// h 的索引号已经没有必要留在 stack 中了
		stack = stack[:len(stack)-1]
	}

	return res
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-13/008403.png)
 时间复杂度O(n),空间复杂度O(n)