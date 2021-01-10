#### 题目
<p>给定一个大小为&nbsp;<em>n&nbsp;</em>的数组，找出其中所有出现超过&nbsp;<code>&lfloor; n/3 &rfloor;</code>&nbsp;次的元素。</p>

<p><strong>说明: </strong>要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> [3,2,3]
<strong>输出:</strong> [3]</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> [1,1,1,3,3,2,2,2]
<strong>输出:</strong> [1,2]</pre>


 #### 题解
 1、哈希表
  ```go
 func majorityElement(nums []int) (ans []int) {
 	m := map[int]int{}
 	for _, num := range nums {
 		m[num]++
 	}
 
 	for k, v := range m {
 		if v > len(nums)/3 {
 			ans = append(ans, k)
 		}
 	}
 	return
 }
 ```
  时间复杂度O(n),空间复杂度O(1)
  
  2、投票法
  ```go
 func majorityElement(nums []int) (ans []int) {
 	n := len(nums)
 	// 最多两个候选人
 	cand1,cand2 := nums[0],nums[0]
 	var count1,count2 int
 	// 选出候选人，遇到不同的候选人消除选票
 	for _, num := range nums {
 		if num == cand1 {
 			count1++
 			continue
 		}
 		if num == cand2 {
 			count2++
 			continue
 		}
 		if count1 == 0 {
 			cand1 = num
 			count1++
 			continue
 		}
 		if count2 == 0 {
 			cand2 = num
 			count2++
 			continue
 		}
 		count1--
 		count2--
 	}
 
 	// 重新计数
 	count1,count2 = 0,0
 	for _, num := range nums {
 		if cand1 == num {
 			count1++
 		} else if cand2 == num {
 			count2++
 		}
 	}
 	
 	// 判断1/3
 	if count1 > n/3 {
 		ans = append(ans, cand1)
 	}
 	if count2 > n/3 {
 		ans = append(ans, cand2)
 	}
 	return
 }
 ```
  时间复杂度O(n),空间复杂度O(1)