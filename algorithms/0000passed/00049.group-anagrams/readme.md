#### 题目
<p>给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> <code>[&quot;eat&quot;, &quot;tea&quot;, &quot;tan&quot;, &quot;ate&quot;, &quot;nat&quot;, &quot;bat&quot;]</code>,
<strong>输出:</strong>
[
  [&quot;ate&quot;,&quot;eat&quot;,&quot;tea&quot;],
  [&quot;nat&quot;,&quot;tan&quot;],
  [&quot;bat&quot;]
]</pre>

<p><strong>说明：</strong></p>

<ul>
	<li>所有输入均为小写字母。</li>
	<li>不考虑答案输出的顺序。</li>
</ul>


 #### 题解
 ```go
func groupAnagrams(strs []string) [][]string {
	var group = make([][]uint8,len(strs))
	var resultM = make(map[string][]string)
	for i, str := range strs {
		group[i] = make([]uint8,26)
		for _, c := range str {
			group[i][c-'a'] += 1
		}
		resultM[string(group[i])] = append(resultM[string(group[i])], str)
	}

	var ret [][]string
	for _, v := range resultM {
		if len(v) > 0 {
			ret = append(ret, v)
		}
	}
	return ret
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-25/004901.png)
 时间复杂度O(KN),空间复杂度O(KN),K表示字符串最大的长度，N表示字符串组的长度