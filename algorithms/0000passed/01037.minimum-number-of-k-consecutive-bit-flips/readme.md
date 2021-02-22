#### 题目
<p>在仅包含 <code>0</code> 和 <code>1</code> 的数组 <code>A</code> 中，一次 <em><code>K</code> 位翻转</em>包括选择一个长度为 <code>K</code> 的（连续）子数组，同时将子数组中的每个 <code>0</code> 更改为 <code>1</code>，而每个 <code>1</code> 更改为 <code>0</code>。</p>

<p>返回所需的 <code>K</code> 位翻转的次数，以便数组没有值为 <code>0</code> 的元素。如果不可能，返回 <code>-1</code>。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>A = [0,1,0], K = 1
<strong>输出：</strong>2
<strong>解释：</strong>先翻转 A[0]，然后翻转 A[2]。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>A = [1,1,0], K = 2
<strong>输出：</strong>-1
<strong>解释：</strong>无论我们怎样翻转大小为 2 的子数组，我们都不能使数组变为 [1,1,1]。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>A = [0,0,0,1,0,1,1,0], K = 3
<strong>输出：</strong>3
<strong>解释：</strong>
翻转 A[0],A[1],A[2]:&nbsp;A变成 [1,1,1,1,0,1,1,0]
翻转 A[4],A[5],A[6]:&nbsp;A变成 [1,1,1,1,1,0,0,0]
翻转 A[5],A[6],A[7]:&nbsp;A变成 [1,1,1,1,1,1,1,1]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ol>
	<li><code>1 &lt;= A.length &lt;=&nbsp;30000</code></li>
	<li><code>1 &lt;= K &lt;= A.length</code></li>
</ol>


 #### 题解
 1、差分数组
 对一个子数组执行两次翻转不会改变数组，所以对每个长度为K的子数组，应至少执行一次翻转操作。
 ```go
func minKBitFlips(A []int, K int) (ans int) {
    n := len(A)
    diff := make([]int,n+1)
    var cnt int
    for i,v := range A {
        cnt += diff[i]
        if (v+cnt)%2 == 0 {
            if i+K>n {
                return -1
            }
            ans++
            cnt++
            diff[i+K]--
        }
    }
    return 
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 2、滑动窗口
 ```go
 func minKBitFlips(A []int, K int) (ans int) {
     n := len(A)
     var cnt int
     for i,v := range A {
         if i >= K&&A[i-K]>1 {
             cnt ^= 1
             A[i-K] -= 2
         }
         if v==cnt {
             if i+K>n {
                 return -1
             }
             ans++
             cnt ^= 1
             A[i] += 2
         }
     }
     return 
 }
```
 时间复杂度O(n),空间复杂度O(1)