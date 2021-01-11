#### 题目
<p>给你一个字符串&nbsp;<code>s</code>，以及该字符串中的一些「索引对」数组&nbsp;<code>pairs</code>，其中&nbsp;<code>pairs[i] =&nbsp;[a, b]</code>&nbsp;表示字符串中的两个索引（编号从 0 开始）。</p>

<p>你可以 <strong>任意多次交换</strong> 在&nbsp;<code>pairs</code>&nbsp;中任意一对索引处的字符。</p>

<p>返回在经过若干次交换后，<code>s</code>&nbsp;可以变成的按字典序最小的字符串。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入：</strong>s = &quot;dcab&quot;, pairs = [[0,3],[1,2]]
<strong>输出：</strong>&quot;bacd&quot;
<strong>解释：</strong> 
交换 s[0] 和 s[3], s = &quot;bcad&quot;
交换 s[1] 和 s[2], s = &quot;bacd&quot;
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>s = &quot;dcab&quot;, pairs = [[0,3],[1,2],[0,2]]
<strong>输出：</strong>&quot;abcd&quot;
<strong>解释：</strong>
交换 s[0] 和 s[3], s = &quot;bcad&quot;
交换 s[0] 和 s[2], s = &quot;acbd&quot;
交换 s[1] 和 s[2], s = &quot;abcd&quot;</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>s = &quot;cba&quot;, pairs = [[0,1],[1,2]]
<strong>输出：</strong>&quot;abc&quot;
<strong>解释：</strong>
交换 s[0] 和 s[1], s = &quot;bca&quot;
交换 s[1] 和 s[2], s = &quot;bac&quot;
交换 s[0] 和 s[1], s = &quot;abc&quot;
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10^5</code></li>
	<li><code>0 &lt;= pairs.length &lt;= 10^5</code></li>
	<li><code>0 &lt;= pairs[i][0], pairs[i][1] &lt;&nbsp;s.length</code></li>
	<li><code>s</code>&nbsp;中只含有小写英文字母</li>
</ul>


 #### 题解
 根据题意,[0,2],[0,3]有公共索引0，所以所以0、2、3可以任意交换。
 所以这题使用并查集找到不同的集合来交换集合内部的数据
 
 ```go
func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	fa := make([]int,n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}
	merge := func(x,y int) {
		fa[find(x)] = find(y)
	}

	for _, p := range pairs {
		merge(p[0],p[1])
	}

	groups := make(map[int][]int,n)
	for i, f := range fa {
		f = find(f)
		// 相连的索引值，放在一起
		// 例: [[0,2,4],[1,3,5]]
		groups[f] = append(groups[f], i)
	}

	bytes := []byte(s)
	res := make([]byte,n)

	for _, g := range groups {
		size := len(g)
		a := make([]int,size)
		copy(a,g)
		// a中的索引值，按照其在bytes中排序
		sort.Slice(a, func(i, j int) bool {
			return bytes[a[i]] < bytes[a[j]]
		})
		// g中索引值排序
		sort.Ints(g)
		for i := 0; i < size; i++ {
			res[g[i]] = bytes[a[i]]
		}
	}
	return string(res)
}
```
 时间复杂度O((M+N)a(N)+NlogN),空间复杂度O(n),M是pairs的长度，N是字符串的长度