#### 题目
<p>找到给定字符串（由小写字符组成）中的最长子串 <strong><em>T</em></strong> ，&nbsp;要求&nbsp;<strong><em>T</em></strong>&nbsp;中的每一字符出现次数都不少于 <em>k</em> 。输出 <strong><em>T&nbsp;</em></strong>的长度。</p>

<p><strong>示例 1:</strong></p>

<pre>
输入:
s = &quot;aaabb&quot;, k = 3

输出:
3

最长子串为 &quot;aaa&quot; ，其中 &#39;a&#39; 重复了 3 次。
</pre>

<p><strong>示例 2:</strong></p>

<pre>
输入:
s = &quot;ababbc&quot;, k = 2

输出:
5

最长子串为 &quot;ababb&quot; ，其中 &#39;a&#39; 重复了 2 次， &#39;b&#39; 重复了 3 次。
</pre>


 #### 题解
 1、分治
 
 对于一个字符串来说，如果存在某个字符ch，它出现的次数在(0,k)之间，那么包含ch的子串肯定不满足要求。
 所以，字符串按ch切分成若干段，则满足要求的最长子串一定在某个切分子串中。
 ```go
func longestSubstring(s string, k int) (ans int) {
    if len(s) == 0 {
        return 0
    }
    cnt := [26]int{}
    for _,ch := range s {
        cnt[ch-'a']++
    }
    var ch byte
    for i,c := range cnt {
        if c > 0 && c < k {
            ch = 'a' + byte(i)
            break
        }
    }
    if ch == 0 {
        return len(s)
    }
    for _,subStr := range strings.Split(s,string(ch)) {
        ans = max(ans,longestSubstring(subStr,k))
    }
    return
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
```
 时间复杂度O(n*E),E为字符集，空间复杂度O(E^2^)
 
 2、滑动窗口
 