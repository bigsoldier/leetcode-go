#### 题目
<p>给定一个二进制数组， 计算其中最大连续1的个数。</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> [1,1,0,1,1,1]
<strong>输出:</strong> 3
<strong>解释:</strong> 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
</pre>

<p><strong>注意：</strong></p>

<ul>
	<li>输入的数组只包含&nbsp;<code>0</code> 和<code>1</code>。</li>
	<li>输入数组的长度是正整数，且不超过 10,000。</li>
</ul>


 #### 题解
 一次遍历
 
 为了得到连续的1的最大个数，需要统计所有的连续1的个数，当遇到不为1的时候，将计数器置0
 ```go
func findMaxConsecutiveOnes(nums []int) (ans int) {
    var cnt int 
    for _,num := range nums {
        if num==1 {
            cnt++
        } else {
            ans = max(ans,cnt)
            cnt = 0
        }
    }
    ans = max(ans,cnt)
    return
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
```
 时间复杂度O(n),空间复杂度O(1)