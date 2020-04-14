#### 题目
<p>给定一个非负整数数组，你最初位于数组的第一个位置。</p>

<p>数组中的每个元素代表你在该位置可以跳跃的最大长度。</p>

<p>判断你是否能够到达最后一个位置。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> [2,3,1,1,4]
<strong>输出:</strong> true
<strong>解释:</strong> 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [3,2,1,0,4]
<strong>输出:</strong> false
<strong>解释:</strong> 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
</pre>


 #### 题解
 这题与前面一道题很像。这里用类似的方法，查找最大跳跃步数，如果当前是0值返回错误
 ```go
func canJump(nums []int) bool {
	var step int
	for step < len(nums)-1 { // 这里 step<len(nums)&&step!=len(nums)-1,表示step要比数组长度小并且不在最后一位上
		var maxStep,index int
		if nums[step] <= 0 { // 如果当前位置的元素为0,则肯定跳不出去了
			return false
		}
		for i := 1; i <= nums[step]; i++ {
			if step+i >= len(nums) {
				return true
			}
			temp := i + nums[step+i]
			if temp > maxStep {
				maxStep = temp
				index = i
			}
		}
		step += index
	}
	return true
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-09/005501.png)
时间复杂度O(n),空间复杂度O(1)

