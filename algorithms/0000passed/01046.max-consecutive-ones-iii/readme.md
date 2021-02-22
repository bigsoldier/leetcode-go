#### 题目
<p>给定一个由若干 <code>0</code> 和 <code>1</code> 组成的数组&nbsp;<code>A</code>，我们最多可以将&nbsp;<code>K</code>&nbsp;个值从 0 变成 1 。</p>

<p>返回仅包含 1 的最长（连续）子数组的长度。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
<strong>输出：</strong>6
<strong>解释： </strong>
[1,1,1,0,0,<strong>1</strong>,1,1,1,1,<strong>1</strong>]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
<strong>输出：</strong>10
<strong>解释：</strong>
[0,0,1,1,<strong>1</strong>,<strong>1</strong>,1,1,1,<strong>1</strong>,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ol>
	<li><code>1 &lt;= A.length &lt;= 20000</code></li>
	<li><code>0 &lt;= K &lt;= A.length</code></li>
	<li><code>A[i]</code> 为&nbsp;<code>0</code>&nbsp;或&nbsp;<code>1</code>&nbsp;</li>
</ol>


 #### 题解
 对于数组A的区间[left,right]而言，
 只要它包含不超过K个0，就满足条件，并且长度为right-left+1
 1、二分查找
 res[right]-res[left-1]<=K
 
 ==> res[left-1]>=res[right]-K
 
 ```go
func longestOnes(A []int, K int) (ans int) {
    n := len(A)
    res := make([]int,n+1)
    for i,v := range A {
        res[i+1] = res[i]+1-v
    }
    for right,v := range res {
        left := sort.SearchInts(res,v-K)
        ans = max(ans,right-left)
    }
    return 
}

func max(a,b int) int {
    if a>b {
        return a
    }
    return b
}
```
 时间复杂度O(nlogn),空间复杂度O(n)
 
 2、滑动窗口
 实时维护left和right，在right向右移动时，同步移动left
 ```go
func longestOnes(A []int, K int) (ans int) {
    var left,lsum,rsum int
    for right,v := range A {
        rsum += 1-v
        for lsum<rsum-K {
            lsum+=1-A[left]
            left++
        }
        ans = max(ans,right-left+1)
    }
    return
}

func max(a,b int) int {
    if a>b {
        return a
    }
    return b
}
```
 时间复杂度O(n),空间复杂度O(1)