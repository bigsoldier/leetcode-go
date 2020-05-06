#### 题目
<p>给定一组<strong>不含重复元素</strong>的整数数组&nbsp;<em>nums</em>，返回该数组所有可能的子集（幂集）。</p>

<p><strong>说明：</strong>解集不能包含重复的子集。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> nums = [1,2,3]
<strong>输出:</strong>
[
  [3],
&nbsp; [1],
&nbsp; [2],
&nbsp; [1,2,3],
&nbsp; [1,3],
&nbsp; [2,3],
&nbsp; [1,2],
&nbsp; []
]</pre>


 #### 题解
 是dfs深度搜索加剪枝的变形题
 ```go
func subsets(nums []int) [][]int {
	var used []int
	var result [][]int
	add(&result,used,nums)
	return result
}

func add(result *[][]int,used, nums []int) {
	var c = make([]int,len(used))
	copy(c,used)
	*result = append(*result, c)
	for i, num := range nums {
		add(result, append(used, num),nums[i+1:])
	}
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-29/007801.png)