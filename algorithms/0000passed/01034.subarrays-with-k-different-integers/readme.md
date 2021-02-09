#### 题目
<p>给定一个正整数数组 <code>A</code>，如果 <code>A</code>&nbsp;的某个子数组中不同整数的个数恰好为 <code>K</code>，则称 <code>A</code> 的这个连续、不一定独立的子数组为<em>好子数组</em>。</p>

<p>（例如，<code>[1,2,3,1,2]</code> 中有&nbsp;<code>3</code>&nbsp;个不同的整数：<code>1</code>，<code>2</code>，以及&nbsp;<code>3</code>。）</p>

<p>返回&nbsp;<code>A</code>&nbsp;中<em>好子数组</em>的数目。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>A = [1,2,1,2,3], K = 2
<strong>输出：</strong>7
<strong>解释：</strong>恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>A = [1,2,1,3,4], K = 3
<strong>输出：</strong>3
<strong>解释：</strong>恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ol>
	<li><code>1 &lt;= A.length &lt;= 20000</code></li>
	<li><code>1 &lt;= A[i] &lt;= A.length</code></li>
	<li><code>1 &lt;= K &lt;= A.length</code></li>
</ol>


 #### 题解
 1、双指针
 ```go
func subarraysWithKDistinct(A []int, K int) int {
    return subarryAtMost(A,K)-subarryAtMost(A,K-1)
}

func subarryAtMost(A []int,K int) (res int) {
    var left,right,count int
    n := len(A)
    freq := make(map[int]int,n+1)
    for right < n {
        if freq[A[right]] == 0 {
            count++
        }
        freq[A[right]]++
        right++
        for count > K {
            freq[A[left]]--
            if freq[A[left]] == 0 {
                count--
            }
            left++
        }
        res += right-left
    }
    return
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 2、优化
 ```go
func subarraysWithKDistinct(A []int, K int) (ans int) {
    n := len(A)
    num1,num2 := make(map[int]int,n+1),make(map[int]int,n+1)
    var left1,left2,to1,to2 int
    for _,a := range A {
        if num1[a] == 0 {
            to1++
        }
        num1[a]++
        if num2[a] == 0 {
            to2++
        }
        num2[a]++
        for to1 > K {
            num1[A[left1]]--
            if num1[A[left1]] == 0 {
                to1--
            }
            left1++
        }
        for to2 > K-1 {
            num2[A[left2]]--
            if num2[A[left2]] == 0 {
                to2--
            }
            left2++
        }
        ans += left2 - left1
    }
    return
}
```
 时间复杂度O(n),空间复杂度O(n)