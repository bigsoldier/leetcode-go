#### 题目
<p>给定&nbsp;<em>n</em> 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。</p>

<p><img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png" style="height: 161px; width: 412px;"></p>

<p><small>上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。&nbsp;<strong>感谢 Marcos</strong> 贡献此图。</small></p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> [0,1,0,2,1,0,1,3,2,1,2,1]
<strong>输出:</strong> 6</pre>


 #### 题解
 1、暴力法
 当前点能储存多少水取决于左右高度。遍历当前点的左右，找到最大值maxLeft,maxRight.其中最小值与当前点的差值就是该点能储存水量的多少。
 ```go
func trap(height []int) int {
	var area int
	for i := 1; i < len(height)-1; i++ {
		var maxLeft,maxRight = height[i],height[i]
		for j := 0; j < i; j++ {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}

		for j := i + 1; j < len(height); j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}

		var min int
		if maxLeft > maxRight {
			min = maxRight
		} else {
			min = maxLeft
		}

		area += min - height[i]
 	}
 	return area
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-18/004201.png)
 时间复杂度O(n^2^),空间复杂度O(1)
 
 2、动态规划
 可以看出方法1中计算左右边界时重复计算了n次，如果我们先计算每个点的左右边界和当前点的比较列表，就不用重复计算了。
 ```go
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	var area int

	var maxLeft,maxRight = make([]int,len(height)),make([]int,len(height))
	maxLeft[0], maxRight[len(height)-1] = height[0],height[len(height)-1]
	for i := 1; i < len(height); i++ {
		if maxLeft[i-1] > height[i] {
			maxLeft[i] = maxLeft[i-1]
		} else {
			maxLeft[i] = height[i]
		}
	}
	for i := len(height) - 2; i >= 0; i-- {
		if maxRight[i+1] > height[i] {
			maxRight[i] = maxRight[i+1]
		} else {
			maxRight[i] = height[i]
		}
	}
	for i := 1; i < len(height)-1; i++ {
		var min int
		if maxLeft[i] > maxRight[i] {
			min = maxRight[i]
		} else {
			min = maxLeft[i]
		}

		area += min - height[i]
	}
	return area
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-18/004202.png)
时间复杂度O(n),空间复杂度O(n)

3、双指针
使用双指针方法，找到中间最高的。从两边边界向中间遍历，遇到比maxLeft和maxRight小的直接求取面积。
```go
func trap(height []int) int {
	var area, maxLeft,maxRight int
	var left,right = 1,len(height)-2
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
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-19/004203.png)
时间复杂度O(n),空间复杂度O(1)

4、栈
```go
func trap(height []int) int {
	var stack = make([]int,len(height))
	var current, sum int
	for current < len(height) {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			top := height[stack[len(stack)-1]] // 栈顶元素
			stack = stack[:len(stack)-1] // 移除栈顶元素
			if len(stack) == 0 {
				break
			}
			last := stack[len(stack)-1]
			// 两堵墙之间的距离
			distance := current - last - 1
			min := height[current]	// 获取当前值和上一个边界的最小值
			if height[last] < min {
				min = height[last]
			}
			sum += distance * (min - top)
		}
		stack = append(stack, current)
		current++
	}
	return sum
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-19/004204.png)
时间复杂度O(n),空间复杂度O(n)