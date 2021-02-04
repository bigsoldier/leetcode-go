#### 题目
<p>如果我们交换字符串&nbsp;<code>X</code> 中的两个不同位置的字母，使得它和字符串&nbsp;<code>Y</code> 相等，那么称 <code>X</code> 和 <code>Y</code> 两个字符串相似。如果这两个字符串本身是相等的，那它们也是相似的。</p>

<p>例如，<code>&quot;tars&quot;</code> 和 <code>&quot;rats&quot;</code> 是相似的 (交换 <code>0</code> 与 <code>2</code> 的位置)；&nbsp;<code>&quot;rats&quot;</code> 和 <code>&quot;arts&quot;</code> 也是相似的，但是 <code>&quot;star&quot;</code> 不与 <code>&quot;tars&quot;</code>，<code>&quot;rats&quot;</code>，或 <code>&quot;arts&quot;</code> 相似。</p>

<p>总之，它们通过相似性形成了两个关联组：<code>{&quot;tars&quot;, &quot;rats&quot;, &quot;arts&quot;}</code> 和 <code>{&quot;star&quot;}</code>。注意，<code>&quot;tars&quot;</code> 和 <code>&quot;arts&quot;</code> 是在同一组中，即使它们并不相似。形式上，对每个组而言，要确定一个单词在组中，只需要这个词和该组中至少一个单词相似。</p>

<p>我们给出了一个字符串列表 <code>A</code>。列表中的每个字符串都是 <code>A</code> 中其它所有字符串的一个字母异位词。请问 <code>A</code> 中有多少个相似字符串组？</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre><strong>输入：</strong>[&quot;tars&quot;,&quot;rats&quot;,&quot;arts&quot;,&quot;star&quot;]
<strong>输出：</strong>2</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>A.length &lt;= 2000</code></li>
	<li><code>A[i].length &lt;= 1000</code></li>
	<li><code>A.length * A[i].length &lt;= 20000</code></li>
	<li><code>A</code> 中的所有单词都只包含小写字母。</li>
	<li><code>A</code> 中的所有单词都具有相同的长度，且是彼此的字母异位词。</li>
	<li>此问题的判断限制时间已经延长。</li>
</ul>

<p>&nbsp;</p>

<p><strong>备注：</strong></p>

<p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 字母异位词[anagram]，一种把某个字符串的字母的位置（顺序）加以改换所形成的新词。</p>


 #### 题解
 1、并查集
 通过比较联通量，比较两个字符串是否相似来连通
 ```go
type unionFind struct {
    fa,rank []int
    count int
}

func newUnionFind(n int) *unionFind {
    fa,rank := make([]int,n),make([]int,n)
    for i := range fa {
        fa[i] = i
        rank[i] = 1
    }
    return &unionFind{fa,rank,n}
}

func (u *unionFind) find(x int) int {
    if x != u.fa[x] {
        u.fa[x] = u.find(u.fa[x])
    }
    return u.fa[x]
}

func (u *unionFind) union(x,y int) {
    fx,fy := u.find(x),u.find(y)
    if fx==fy {
        return
    }
    if u.rank[fx] < u.rank[fy] {
        fx,fy = fy,fx
    }
    u.rank[fx] += u.rank[fy]
    u.fa[fy] = fx
    u.count--
}

func (u *unionFind) isSame(x,y int) bool {
    return u.find(x)==u.find(y)
}

func numSimilarGroups(strs []string) int {
    n := len(strs)
    uf := newUnionFind(n)
    for i,s := range strs {
        for j := i+1;j < n; j++ {
            if !uf.isSame(i,j) && isSimilar(s,strs[j]) {
                uf.union(i,j)
            }
        }
    }
    return uf.count
}

// 判断两个字符串是否相似
func isSimilar(x,y string) bool {
    var diff int
    for i := range x {
        if x[i] != y[i] {
            diff++
            if diff > 2 {
                return false
            }
        }
    }
    return true
}
```
 时间复杂度O(n^2^m+nlogn),空间复杂度O(n),
 n是字符串的数量，需要O(n^2^)来枚举任一对字符串之间的关系，对于任一对字符串，需要O(m)来比较字符串是否相似。
 最坏情况下需要对并查集执行O(n)次合并，平摊O(logn)
 