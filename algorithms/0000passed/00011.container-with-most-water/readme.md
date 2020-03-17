#### 题目
<p>给你 <em>n</em> 个非负整数 <em>a</em><sub>1</sub>，<em>a</em><sub>2，</sub>...，<em>a</em><sub>n，</sub>每个数代表坐标中的一个点&nbsp;(<em>i</em>,&nbsp;<em>a<sub>i</sub></em>) 。在坐标内画 <em>n</em> 条垂直线，垂直线 <em>i</em>&nbsp;的两个端点分别为&nbsp;(<em>i</em>,&nbsp;<em>a<sub>i</sub></em>) 和 (<em>i</em>, 0)。找出其中的两条线，使得它们与&nbsp;<em>x</em>&nbsp;轴共同构成的容器可以容纳最多的水。</p>

<p><strong>说明：</strong>你不能倾斜容器，且&nbsp;<em>n</em>&nbsp;的值至少为 2。</p>

<p>&nbsp;</p>

<p><img alt="" src="https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg" style="height: 287px; width: 600px;"></p>

<p><small>图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为&nbsp;49。</small></p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre><strong>输入：</strong>[1,8,6,2,5,4,8,3,7]
<strong>输出：</strong>49</pre>


 #### 题解
 1、暴力法
 两重遍历数组，计算出最大面积
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