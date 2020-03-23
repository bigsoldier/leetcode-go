#### 题目
<p>给定一个非负整数数组，你最初位于数组的第一个位置。</p>

<p>数组中的每个元素代表你在该位置可以跳跃的最大长度。</p>

<p>你的目标是使用最少的跳跃次数到达数组的最后一个位置。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> [2,3,1,1,4]
<strong>输出:</strong> 2
<strong>解释:</strong> 跳到最后一个位置的最小跳跃数是 <code>2</code>。
&nbsp;    从下标为 0 跳到下标为 1 的位置，跳&nbsp;<code>1</code>&nbsp;步，然后跳&nbsp;<code>3</code>&nbsp;步到达数组的最后一个位置。
</pre>

<p><strong>说明:</strong></p>

<p>假设你总是可以到达数组的最后一个位置。</p>


 #### 题解
 需要确保下一次跳跃是最远的位置。那么选取可跳跃最大位置内的元素，计算最大的步数。
 例子：[2,3,1,1,4]
 步骤1：当前索引为0,对应的元素值为2,可跳跃的目标的索引为[1,2],对应的元素值加上索引值为[4,3],所以最大步数为4,则下一步在索引1上,count++
 步骤2：当前索引为1,对应的元素值为3,可跳跃的目标的索引为[2,3,4],对应的元素值加上索引值为[3,4,8],所以最大步数为8,而此时当前元素值加上最大步数已经超过数组的长度，所以长度为2
 
 ```go
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var index,count int
	for index < len(nums) {
		if nums[index]+index+1 >= len(nums) {
			return count+1
		}
		var maxStep int	// 记录当前跳跃范围内可统计的最大步数
		var maxIndex int	// 最大步数的index
		// nums[index]表示当前位置的最大跳跃步数
		for i := 1; i <= nums[index]; i++ {	// 遍历可跳的最大步数
			if nums[index+i]+i > maxStep {
				maxStep = nums[index+i]+i
				maxIndex = i
			}
		}
		index += maxIndex
		count++
	}
	return count
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-23/004501.png)
时间复杂度O(n^2^),空间复杂度O(1)