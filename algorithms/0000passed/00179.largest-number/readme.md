#### 题目
<p>给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> <code>[10,2]</code>
<strong>输出:</strong> <code>210</code></pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> <code>[3,30,34,5,9]</code>
<strong>输出:</strong> <code>9534330</code></pre>

<p><strong>说明: </strong>输出结果可能非常大，所以你需要返回一个字符串而不是整数。</p>


 #### 题解
 根据题目的排序规则，3比30大，那么我们可以利用字符串比较的方法，将3和30组合起来进行比较
 ```go
func largestNumber(nums []int) (ans string) {
	var strs []string
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	sort.SliceStable(strs, func(i, j int) bool {
		a,b := strs[i],strs[j]
		return a+b>b+a
	})
	if strs[0] == "0" {
		return "0"
	}
	for _, str := range strs {
		ans += str
	}
	return
}
```
 时间复杂度O(nlogn),空间复杂度O(logn)