#### 题目
<p>给定一个非空且只包含非负数的整数数组&nbsp;<code>nums</code>, 数组的度的定义是指数组里任一元素出现频数的最大值。</p>

<p>你的任务是找到与&nbsp;<code>nums</code>&nbsp;拥有相同大小的度的最短连续子数组，返回其长度。</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> [1, 2, 2, 3, 1]
<strong>输出:</strong> 2
<strong>解释:</strong> 
输入数组的度是2，因为元素1和2的出现频数最大，均为2.
连续子数组里面拥有相同度的有如下所示:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
最短连续子数组[2, 2]的长度为2，所以返回2.
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> [1,2,2,3,1,4,2]
<strong>输出:</strong> 6
</pre>

<p><strong>注意:</strong></p>

<ul>
	<li><code>nums.length</code>&nbsp;在1到50,000区间范围内。</li>
	<li><code>nums[i]</code>&nbsp;是一个在0到49,999范围内的整数。</li>
</ul>


 #### 题解
 1、哈希表
 记录原数组出现次数最多的元素，同时记录出现第一次和最后一次的位置。
 ```go
func findShortestSubArray(nums []int) (ans int) {
    var set = make(map[int]entry)
    // 记录元素
    for i,v := range nums {
        if e,has := set[v];has {
            e.cnt++
            e.right = i
            set[v] = e
        } else {
            set[v] = entry{i,i,1}
        }
    }
    var maxCnt int
    for _,e := range set {
        if e.cnt > maxCnt {
            maxCnt,ans = e.cnt,e.right-e.left+1
        } else if e.cnt == maxCnt {
            ans = min(ans,e.right-e.left+1)
        }
    }
    return
}

type entry struct {
    left,right,cnt int // 记录元素第一次和最后一次出现的位置及次数
}

func min(a,b int) int {
    if a<b {
        return a
    }
    return b
}
```
 时间复杂度O(n),空间复杂度O(n)