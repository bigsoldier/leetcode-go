#### 题目
<p>给定一个整数数组&nbsp;<code>nums</code>，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。</p>

<p><strong>示例 :</strong></p>

<pre><strong>输入:</strong> <code>[1,2,1,3,2,5]</code>
<strong>输出:</strong> <code>[3,5]</code></pre>

<p><strong>注意：</strong></p>

<ol>
	<li>结果输出的顺序并不重要，对于上面的例子，&nbsp;<code>[5, 3]</code>&nbsp;也是正确答案。</li>
	<li>你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？</li>
</ol>


 #### 题解
 1、哈希表
 ```go
func singleNumber(nums []int) []int {
	var m = map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	output := []int{}
	for num, val := range m {
		if val == 1 {
			output = append(output, num)
		}
	}
	return output
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 2、位运算
 ```go
func singleNumber(nums []int) []int {
	var xor int
	// 求出 x^y
	for _, num := range nums {
		xor ^= num
	}
	// 求出最右边的1
	lowest := xor & -xor
	var a int
	for _, num := range nums {
		if num&lowest == 0 {
			a^=num
		}
	}
	return []int{a,xor^a}
}
```
 时间复杂度O(n),空间复杂度O(1)