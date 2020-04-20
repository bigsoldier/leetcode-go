#### 题目
<p>给出集合&nbsp;<code>[1,2,3,&hellip;,<em>n</em>]</code>，其所有元素共有&nbsp;<em>n</em>! 种排列。</p>

<p>按大小顺序列出所有排列情况，并一一标记，当&nbsp;<em>n </em>= 3 时, 所有排列如下：</p>

<ol>
	<li><code>&quot;123&quot;</code></li>
	<li><code>&quot;132&quot;</code></li>
	<li><code>&quot;213&quot;</code></li>
	<li><code>&quot;231&quot;</code></li>
	<li><code>&quot;312&quot;</code></li>
	<li><code>&quot;321&quot;</code></li>
</ol>

<p>给定&nbsp;<em>n</em> 和&nbsp;<em>k</em>，返回第&nbsp;<em>k</em>&nbsp;个排列。</p>

<p><strong>说明：</strong></p>

<ul>
	<li>给定<em> n</em>&nbsp;的范围是 [1, 9]。</li>
	<li>给定 <em>k&nbsp;</em>的范围是[1, &nbsp;<em>n</em>!]。</li>
</ul>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> n = 3, k = 3
<strong>输出:</strong> &quot;213&quot;
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> n = 4, k = 9
<strong>输出:</strong> &quot;2314&quot;
</pre>


 #### 题解
 1、暴力法
 用回溯的思想(递归)写出全排列，生成n!个全排列需要时间复杂度O(n*n!)。
 
 2、直接找第k个排列
 举个栗子:n = 5,k = 5
 这种情况下,我们知道第一位为1的时候会有24种情况(4!),此时k<24,那么第二位为2的时候有6种情况(3!),此时k<6,那么第三位为3的时候有2种情况,那么k>2,同理，第三位为4和5的时候也有2种情况,可以取第三位等于5的情况，k/2=2,k%2=1,结果就是12534
 ```go
func getPermutation(n int, k int) string {
	k-- // 从0开始
	if n == 0 {
		return ""
	}
	ret := make([]byte,n)

	// 存放 从1到n 的数字
	nums := make([]byte,n)
	for i := 0; i < n; i++ {
		nums[i] = byte(i) + '1'
	}

	// 阶乘
	fib := 1
	for i := 2; i < n; i++ {
		fib *= i
	}

	for i := 0; i < n-1; i++ {
		idx := k / fib
		ret[i] = nums[idx]
		// 去除已使用的数
		nums = append(nums[:idx], nums[idx+1:]...)
		k %= fib
		fib /= (n-i-1)
	}
	ret[n-1] = nums[0]
	return string(ret)
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-17/006001.png)
时间复杂度O(n),空间复杂度O(n)