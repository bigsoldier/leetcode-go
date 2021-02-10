#### 题目
<p>给定两个字符串&nbsp;<strong>s1</strong>&nbsp;和&nbsp;<strong>s2</strong>，写一个函数来判断 <strong>s2</strong> 是否包含 <strong>s1&nbsp;</strong>的排列。</p>

<p>换句话说，第一个字符串的排列之一是第二个字符串的子串。</p>

<p><strong>示例1:</strong></p>

<pre>
<strong>输入: </strong>s1 = &quot;ab&quot; s2 = &quot;eidbaooo&quot;
<strong>输出: </strong>True
<strong>解释:</strong> s2 包含 s1 的排列之一 (&quot;ba&quot;).
</pre>

<p>&nbsp;</p>

<p><strong>示例2:</strong></p>

<pre>
<strong>输入: </strong>s1= &quot;ab&quot; s2 = &quot;eidboaoo&quot;
<strong>输出:</strong> False
</pre>

<p>&nbsp;</p>

<p><strong>注意：</strong></p>

<ol>
	<li>输入的字符串只包含小写字母</li>
	<li>两个字符串的长度都在 [1, 10,000] 之间</li>
</ol>


 #### 题解
 1、动态规划
 由于排列不会改变字符串中每个字符串的个数，所以只有当两个字符串每个字符的个数均相等。
 ```go
func checkInclusion(s1 string, s2 string) bool {
    m,n := len(s1),len(s2)
    if m > n {
        return false
    }
    var cnt1,cnt2 [26]int 
    for i,ch := range s1 {
        cnt1[ch-'a']++
        cnt2[s2[i]-'a']++
    }
    if cnt1 == cnt2 {
        return true
    }
    for i := m; i < n; i++ {
        cnt2[s2[i]-'a']++
        cnt2[s2[i-m]-'a']--
        if cnt1 == cnt2 {
            return true
        }
    }
    return false
}
```
 时间复杂度O(m+n+26),空间复杂度O(26)
 
 2、双指针
 ```go
func checkInclusion(s1 string, s2 string) bool {
    m,n := len(s1),len(s2)
    if m > n {
        return false
    }
    cnt := [26]int{}
    for _,ch := range s1 {
        cnt[ch-'a']--
    }
    left := 0
    for right,ch := range s2 {
        x := ch - 'a'
        cnt[x]++
        for cnt[x] > 0 {
            cnt[s2[left]-'a']--
            left++
        }
        if right - left + 1 == m {
            return true
        }
    }
    return false
}
```
 时间复杂度O(m+n+26),空间复杂度O(26)