#### 题目
<p>给定一个数组&nbsp;<code>candidates</code>&nbsp;和一个目标数&nbsp;<code>target</code>&nbsp;，找出&nbsp;<code>candidates</code>&nbsp;中所有可以使数字和为&nbsp;<code>target</code>&nbsp;的组合。</p>

<p><code>candidates</code>&nbsp;中的每个数字在每个组合中只能使用一次。</p>

<p><strong>说明：</strong></p>

<ul>
	<li>所有数字（包括目标数）都是正整数。</li>
	<li>解集不能包含重复的组合。&nbsp;</li>
</ul>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> candidates =&nbsp;<code>[10,1,2,7,6,1,5]</code>, target =&nbsp;<code>8</code>,
<strong>所求解集为:</strong>
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> candidates =&nbsp;[2,5,2,1,2], target =&nbsp;5,
<strong>所求解集为:</strong>
[
&nbsp; [1,2,2],
&nbsp; [5]
]</pre>


 #### 题解
 递归，不能重复意味着要取的下一个值不能重复，所以循环获取下一个不重复的数
 ```go
 func combinationSum2(candidates []int, target int) [][]int {
 	sort.Ints(candidates)
 	var solution []int
 	var result [][]int
 	findCombinationSum(candidates,solution,target,&result)
 	return result
 }
 
 func findCombinationSum(candidates, solution []int, target int, result *[][]int) {
 	if target == 0 {
 		*result = append(*result, solution)
 	}
 
 	if len(candidates) == 0 || target < candidates[0] {
 		return
 	}
 
 	findCombinationSum(candidates[1:], append(solution, candidates[0]),target-candidates[0],result)
 	var i int
 	for i < len(candidates)-1 {
 		if candidates[i] != candidates[i+1] {
 			findCombinationSum(candidates[i+1:],solution,target,result)
 			break
 		} else {
 			i++
 		}
 	}
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-12/004002.png)