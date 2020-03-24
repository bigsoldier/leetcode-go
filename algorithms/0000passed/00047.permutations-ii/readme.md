#### 题目
<p>给定一个可包含重复数字的序列，返回所有不重复的全排列。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> [1,1,2]
<strong>输出:</strong>
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]</pre>


 #### 题解
 和47题类似，添加了去重
 ```go
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var output [][]int
	used := make([]bool,len(nums))
	var p []int
	backtrack(nums,0, p,&output, &used)
	return output
}

func backtrack(nums []int, index int, p []int, output *[][]int, used *[]bool) {
	if len(nums) == index {
		var tmp = make([]int,len(nums))
		copy(tmp,p)
		*output = append(*output, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
				continue
			}
			(*used)[i] = true
			backtrack(nums,index+1,append(p,nums[i]),output,used)
			(*used)[i] = false
		}
	}
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-24/004701.png)
