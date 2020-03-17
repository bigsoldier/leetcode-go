#### 题目
<p>给定一个<strong>无重复元素</strong>的数组&nbsp;<code>candidates</code>&nbsp;和一个目标数&nbsp;<code>target</code>&nbsp;，找出&nbsp;<code>candidates</code>&nbsp;中所有可以使数字和为&nbsp;<code>target</code>&nbsp;的组合。</p>

<p><code>candidates</code>&nbsp;中的数字可以无限制重复被选取。</p>

<p><strong>说明：</strong></p>

<ul>
	<li>所有数字（包括&nbsp;<code>target</code>）都是正整数。</li>
	<li>解集不能包含重复的组合。&nbsp;</li>
</ul>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> candidates = <code>[2,3,6,7], </code>target = <code>7</code>,
<strong>所求解集为:</strong>
[
  [7],
  [2,2,3]
]
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> candidates = [2,3,5]<code>, </code>target = 8,
<strong>所求解集为:</strong>
[
&nbsp; [2,2,2,2],
&nbsp; [2,3,3],
&nbsp; [3,5]
]</pre>


 #### 题解
 1、考虑使用递归，使用target减去当前数字，在之后的递归中在减去当前之后的数字，如果等于0，则加入解集中
 ```go
 func combinationSum(candidates []int, target int) [][]int {
 	if len(candidates) == 0 {
 		return nil
 	}
 
 	var solution []int
 	var result [][]int
 	findCombiationSum(candidates,solution,target,0,&result)
 	return result
 }
 
 // solution 存放当前的可能解
 func findCombiationSum(candidates, solution []int, target, index int, result *[][]int) {
 	if target <= 0 {
 		if target == 0 {
 			// 这里需要将solutio新建一个空间存储，否则会影响
 			ret := make([]int,len(solution))
 			copy(ret,solution)
 			*result = append(*result, ret)
 		}
 		return
 	}
 
 	for i := index; i < len(candidates); i++ {
 		findCombiationSum(candidates, append(solution, candidates[i]), target-candidates[i], i,result)
 	}
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-12/003902.png)
 
 2、优化
 排序之后，向后遍历
 ```go
 func combinationSum(candidates []int, target int) [][]int {
 	var solution []int
 	var result [][]int
 	sort.Ints(candidates)
 	findCombiationSum(candidates,solution,target,&result)
 	return result
 }
 
 // solution 存放当前的可能解
 func findCombiationSum(candidates, solution []int, target int, result *[][]int) {
 	if target == 0 {
 		*result = append(*result, solution)
 	}
 	if len(candidates) == 0 || target < candidates[0] {
 		return
 	}
 	
 	solution = solution[:len(solution):len(solution)]
 	// 等价于
 	// var ret = make([]int,len(solution))
 	// copy(ret,solution)
 	findCombiationSum(candidates, append(solution, candidates[0]), target-candidates[0],result)
 	findCombiationSum(candidates[1:],solution,target,result)
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-12/003903.png)