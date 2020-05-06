#### 题目
<p>给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。</p>

<p><strong>示例：</strong></p>

<pre><strong>输入: S</strong> = &quot;ADOBECODEBANC&quot;, <strong>T</strong> = &quot;ABC&quot;
<strong>输出:</strong> &quot;BANC&quot;</pre>

<p><strong>说明：</strong></p>

<ul>
	<li>如果 S 中不存这样的子串，则返回空字符串 <code>&quot;&quot;</code>。</li>
	<li>如果 S 中存在这样的子串，我们保证它是唯一的答案。</li>
</ul>


 #### 题解
 滑动窗口
 左右指针都在左边，右指针右移，直到包含所有的t，然后移动左指针，使得也能包含所有的t，这样就能找出最小的窗口了
 ```go
func minWindow(s string, t string) string {
	used := [128]int{}
	dict := [128]int{}
	for i := range t {
		dict[t[i]]++
	}

	size, total := len(s), len(t)

	min := size + 1
	res := ""

	// s[left:right+1] 就是 window
	// count 用于统计已有的 t 中字母的数量(在used中)。
	// count == total 表示已经收集完需要的全部字母
	for left, right, count := 0, 0, 0; right < size; right++ {
		if used[s[right]] < dict[s[right]] {
			// 出现了 window 中缺失的字母
			count++
		}
		used[s[right]]++

		// 保证 window 不丢失所需字母的前提下
		// 让 left 尽可能的大
		for left <= right && used[s[left]] > dict[s[left]] {
			used[s[left]]--
			left++
		}

		width := right - left + 1
		if count == total && min > width {
			min = width
			res = s[left : right+1]
		}

	}

	return res
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-29/007601.png)
时间复杂度O(S+T),空间复杂度O(S+T),