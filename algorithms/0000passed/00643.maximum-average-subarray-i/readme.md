#### 题目
<p>给定 <code>n</code> 个整数，找出平均数最大且长度为 <code>k</code> 的连续子数组，并输出该最大平均数。</p>

<p> </p>

<p><strong>示例：</strong></p>

<pre>
<strong>输入：</strong>[1,12,-5,-6,50,3], k = 4
<strong>输出：</strong>12.75
<strong>解释：</strong>最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li>1 <= <code>k</code> <= <code>n</code> <= 30,000。</li>
	<li>所给数据范围 [-10,000，10,000]。</li>
</ul>


 #### 题解
 1、滑动窗口
 ```go
func findMaxAverage(nums []int, k int) (ans float64) {
    var sum int
    for _,num := range nums[:k] {
        sum += num
    }
    ans = float64(sum) / float64(k)
    for i := k;i < len(nums); i ++ {
        sum += nums[i]
        sum -= nums[i-k]
        ans = max(ans,float64(sum)/float64(k))
    }
    return 
}

func max(a,b float64) float64 {
	if a > b {
		return a
	}
	return b
}
```
 时间复杂度O(n),空间复杂度O(1)