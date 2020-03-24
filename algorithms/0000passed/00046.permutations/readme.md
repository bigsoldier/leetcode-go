#### 题目
<p>给定一个<strong>没有重复</strong>数字的序列，返回其所有可能的全排列。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> [1,2,3]
<strong>输出:</strong>
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]</pre>


 #### 题解
 递归(回溯法)
 ```go
func permute(nums []int) [][]int {
	var output [][]int
	n := len(nums)
	backtrack(n,0,nums,&output)
	return output
}

// 使用指针获取值
func backtrack(n, index int, nums []int, output *[][]int) {
	if n == index {
        // 获取数据的拷贝
		var tmp = make([]int,len(nums))
		copy(tmp,nums)
		*output = append(*output, tmp)
	}
	for i := index; i < n; i++ {
		nums[i],nums[index] = nums[index],nums[i]
		backtrack(n,index+1,nums,output)
		nums[i],nums[index] = nums[index],nums[i]
	}
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-23/004601.png)
时间复杂度O(n!),空间复杂度O(n!)