#### 题目
<p>给定一个包含 <code>0, 1, 2, ..., n</code>&nbsp;中&nbsp;<em>n</em>&nbsp;个数的序列，找出 0 .. <em>n</em>&nbsp;中没有出现在序列中的那个数。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [3,0,1]
<strong>输出:</strong> 2
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [9,6,4,2,3,5,7,0,1]
<strong>输出:</strong> 8
</pre>

<p><strong>说明:</strong><br>
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?</p>


 #### 题解
 1、排序
 ```go
func missingNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if nums[n-1] != n {
		return n
	} else if nums[0] != 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		if i != nums[i] {
			return i
		}
	}
	return -1
}
```
 时间复杂度O(nlogn),空间复杂度O(1)
 
 2、哈希表
 ```go
func missingNumber(nums []int) int {
    var m = map[int]bool{}
	for _,num := range nums {
        m[num] = true
    }
    for i := 0;i < len(nums)+1; i++ {
        if !m[i] {
            return i
        }
    }
    return -1
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 3、位运算
 ```go
func missingNumber(nums []int) int {
	var missing = len(nums)
	for i, num := range nums {
		missing ^= i ^ num
	}
	return missing
}
```
 时间复杂度O(n),空间复杂度O(1)