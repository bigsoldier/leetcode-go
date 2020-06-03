#### 题目
<p>给定一个可能包含重复元素的整数数组 <em><strong>nums</strong></em>，返回该数组所有可能的子集（幂集）。</p>

<p><strong>说明：</strong>解集不能包含重复的子集。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> [1,2,2]
<strong>输出:</strong>
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]</pre>


 #### 题解
 1、回溯法
 判断当前数字和上一个数字是否相等，相等的话则跳过。
 ```go
 func subsetsWithDup(nums []int) [][]int {
 	var result [][]int
 	sort.Ints(nums)
 	getSets(nums,0,nil,&result)
 	return result
 }
 
 func getSets(nums []int, start int, temp []int, result *[][]int) {
 	*result = append(*result, temp)
 	for i := start; i < len(nums); i++ {
 		if i > start && nums[i] == nums[i-1] {
 			continue
 		}
 		getSets(nums,i+1,append(temp, nums[i]),result)
 	}
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-03/009001.png)
 
 2、动态规划
 ```$xsl
 [1,2,3]的结果
 []的子串 []
 [1]的子串 [][1]
 [1,2]的子串 [][1][2][1,2]
 [1,2,3]的子串 [][1][2][1,2][3][1,3][2,3][1,2,3]
 ```
 ```$xsl
  [1,2,2]的结果
  []的子串 []
  [1]的子串 [][1]
  [1,2]的子串 [][1][2][1,2]
  [1,2,2]的子串 [][1][2][1,2][2,2][1,2,2]
 ```
 可以看出每增加一个数字都是在上一个数字子串的基础上添加
 ```go
 func subsetsWithDup(nums []int) [][]int {
 	var result = make([][]int,1)
 	sort.Ints(nums)
 	var start = len(result)
 	for i := 0; i < len(nums); i++ {
 		var ret [][]int
 		for j := 0; j < len(result); j++ {
 			if i > 0 && nums[i] == nums[i-1] && j < start {
 				continue
 			}
 			tmp := append(result[j], nums[i])
 			ret = append(ret, tmp)
 		}
 		start = len(result)
 		result = append(result, ret...)
 	}
 	return result
 }
```
 结果集相同，但是顺序不一致
 时间复杂度O(n^2^),空间复杂度O(n^2^)